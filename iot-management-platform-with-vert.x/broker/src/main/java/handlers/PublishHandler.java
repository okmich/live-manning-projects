package handlers;

import com.smarthome.broker.MainVerticle;
import data.MongoStore;
import data.Store;
import io.vertx.core.Handler;
import io.vertx.core.buffer.Buffer;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.core.json.JsonObject;
import io.vertx.mqtt.MqttEndpoint;
import io.vertx.mqtt.messages.MqttPublishMessage;

import java.time.LocalDate;
import java.time.LocalTime;

public class PublishHandler implements Handler<MqttPublishMessage> {

  private final static Logger LOGGER = LoggerFactory.getLogger(MainVerticle.class);
  private final MqttEndpoint mqttEndpoint;

  public PublishHandler(MqttEndpoint mqttEndpoint) {
    this.mqttEndpoint = mqttEndpoint;
  }

  @Override
  public void handle(MqttPublishMessage mqttPublishMessage) {
    var message = new String(mqttPublishMessage.payload().getBytes());
    System.out.println(message);
    try {
      var jsonObject = JsonObject.of("topic", mqttPublishMessage.topicName(),
        "device", new JsonObject(message),
        "date", LocalDate.now().toString(),
        "hour", LocalTime.now().toString());

      MongoStore.getMongoClient().save("devices", jsonObject)
        .onFailure(error -> {
          LOGGER.error(error.getMessage(), error);
        });
    } catch (Exception ex) {
      LOGGER.error(ex.getMessage(), ex);
    }

    Store.getMqttSubscriptions().forEach((id, mqttSubscription) -> {
      if(mqttSubscription.getTopic().equals(mqttPublishMessage.topicName()) && mqttSubscription.getMqttEndpoint().isConnected()) {
        mqttSubscription.getMqttEndpoint().publish(
          mqttPublishMessage.topicName(),
          Buffer.buffer(message),
          mqttPublishMessage.qosLevel(),
          false, false
        );
      }
    });

    switch (mqttPublishMessage.qosLevel()) {
      case AT_LEAST_ONCE:
        mqttEndpoint.publishAcknowledge(mqttPublishMessage.messageId());
        break;
      case EXACTLY_ONCE:
        mqttEndpoint.publishReceived(mqttPublishMessage.messageId());
        break;
      case FAILURE:
        break;
    }

  }
}
