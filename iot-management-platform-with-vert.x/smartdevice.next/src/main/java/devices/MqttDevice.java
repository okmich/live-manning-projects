package devices;

import communications.Mqtt;
import io.netty.handler.codec.mqtt.MqttQoS;
import io.vertx.circuitbreaker.CircuitBreaker;
import io.vertx.core.Future;
import io.vertx.core.Vertx;
import io.vertx.core.buffer.Buffer;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.mqtt.MqttClient;
import io.vertx.mqtt.MqttClientOptions;
import io.vertx.mqtt.messages.MqttConnAckMessage;
import sensors.Sensor;

import java.util.LinkedList;
import java.util.List;

public class MqttDevice implements Device, Mqtt {
  private final Logger LOGGER = LoggerFactory.getLogger(MqttDevice.class);
  private final int mqttPort;
  private final String mqttHost;
  private final String mqttTopic;
  private final LinkedList<Sensor> sensors = new LinkedList<>();
  private final String id;
  private MqttClient mqttClient;
  private CircuitBreaker circuitBreaker;
  private String position;
  private String category;

  public MqttDevice(String id, String mqttHost, int mqttPort, String mqttTopic) {
    this.id = id;
    this.mqttPort = mqttPort;
    this.mqttHost = mqttHost;
    this.mqttTopic = mqttTopic;
  }

  @Override
  public MqttClient createMqttClient(Vertx vertx) {
    return MqttClient.create(vertx, new MqttClientOptions().setClientId(this.id));
  }

  @Override
  public Future<MqttConnAckMessage> startAndConnectMqttClient(Vertx vertx) {
    if (circuitBreaker == null) {
      this.circuitBreaker = getBreaker(vertx);
    }

    return this.circuitBreaker.execute(promise -> {
      if (mqttClient == null)
        mqttClient = createMqttClient(vertx);

      mqttClient.connect(mqttPort, mqttHost)
        .onFailure(error -> {
          LOGGER.error("MQTT " + error.getMessage());
          promise.fail("[" + error.getMessage() + "]");
        })
        .onSuccess(ok -> {
          mqttClient.publish(
              this.mqttTopic,
              Buffer.buffer(jsonValue().encode()),
              MqttQoS.AT_MOST_ONCE, false, false
            ).onSuccess(res -> {
                mqttClient.disconnect();
                promise.complete();
              }
            )
            .onFailure(err -> {
              LOGGER.error(err.getMessage(), err);
              mqttClient.disconnect();
              promise.fail("[" + err.getMessage() + "]");
            });
        });
    });
  }

  @Override
  public String getProtocol() {
    return "mqtt";
  }

  @Override
  public MqttClient getMqttClient() {
    return this.mqttClient;
  }

  @Override
  public LinkedList<Sensor> getSensors() {
    return this.sensors;
  }

  @Override
  public Device setSensors(List<Sensor> sensors) {
    this.sensors.addAll(sensors);
    return this;
  }

  @Override
  public String getId() {
    return this.id;
  }

  @Override
  public String getPosition() {
    return this.position;
  }

  @Override
  public Device setPosition(String value) {
    this.position = value;
    return this;
  }

  @Override
  public String getCategory() {
    return this.category;
  }

  @Override
  public Device setCategory(String value) {
    this.category = value;
    return this;
  }
}
