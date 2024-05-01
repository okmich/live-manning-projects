package http;

import io.netty.handler.codec.mqtt.MqttQoS;
import io.vertx.core.Handler;
import io.vertx.core.buffer.Buffer;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.ext.web.client.WebClient;
import io.vertx.mqtt.MqttClient;
import io.vertx.servicediscovery.Record;
import io.vertx.servicediscovery.ServiceDiscovery;

import java.util.List;

public class DevicesHealth {

  private final Logger LOGGER = LoggerFactory.getLogger(Registration.class);

  private final ServiceDiscovery serviceDiscovery;
  private final WebClient webClient;
  private final MqttClient mqttClient;
  private final String mqttTopic;

  public DevicesHealth(ServiceDiscovery serviceDiscovery, WebClient webClient, MqttClient mqttClient, String mqttTopic) {
    this.serviceDiscovery = serviceDiscovery;
    this.webClient = webClient;
    this.mqttClient = mqttClient;
    this.mqttTopic = mqttTopic;
  }

  public void handleDevicesHealthCheck(Long aLong) {
    this.serviceDiscovery.getRecords(rec -> true).onSuccess(this.handleRecords()).onFailure(this.basicOnFailure());
  }

  private Handler<List<Record>> handleRecords() {
    return records -> {
      records.forEach(record -> {
        var location = record.getLocation();
        webClient.post(location.getInteger("port"), location.getString("host"), "/").send()
          .onSuccess(res -> {
            if (mqttClient != null) {
              if (mqttClient.isConnected()) {
                mqttClient.publish(
                  this.mqttTopic, //
                  Buffer.buffer(res.bodyAsString()),
                  MqttQoS.AT_LEAST_ONCE,
                  false, // duplicated
                  false //Retained
                );
              } else {
                //log the disconnection
                LOGGER.error("MQTT is disconnected. Cannot publish device health results");
              }
            }
            LOGGER.info("What next with " + res.bodyAsJsonObject());
          })
          .onFailure(err -> {
            LOGGER.error("Unabled to reach device. Reason: " + err.getMessage());
            this.serviceDiscovery.unpublish(record.getRegistration());
          });
      });
    };
  }

  private Handler<Throwable> basicOnFailure() {
    return throwable -> {
      LOGGER.error(throwable.getMessage(), throwable);
    };
  }
}
