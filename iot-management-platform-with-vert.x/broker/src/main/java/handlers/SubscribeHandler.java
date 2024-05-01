package handlers;

import com.smarthome.broker.MainVerticle;
import data.Store;
import io.netty.handler.codec.mqtt.MqttQoS;
import io.vertx.core.Handler;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.mqtt.MqttEndpoint;
import io.vertx.mqtt.messages.MqttSubscribeMessage;
import models.MqttSubscription;

import java.util.ArrayList;

public class SubscribeHandler implements Handler<MqttSubscribeMessage> {

  private final static Logger LOGGER = LoggerFactory.getLogger(MainVerticle.class);
  private final MqttEndpoint mqttEndpoint;

    public SubscribeHandler(MqttEndpoint mqttEndpoint) {
        this.mqttEndpoint = mqttEndpoint;
    }

    @Override
  public void handle(MqttSubscribeMessage mqttSubscribeMessage) {
      var qsLevels = new ArrayList<MqttQoS>();

      mqttSubscribeMessage.topicSubscriptions().forEach(ts -> {

        String subKey = String.format("%s|%s", mqttEndpoint.clientIdentifier(), ts.topicName());
        Store.getMqttSubscriptions().put(
          subKey, new MqttSubscription(ts.topicName(), this.mqttEndpoint)
        );
        qsLevels.add(ts.qualityOfService());
      });
      mqttEndpoint.subscribeAcknowledge(mqttSubscribeMessage.messageId(), qsLevels);
  }
}
