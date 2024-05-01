package handlers;

import com.smarthome.broker.MainVerticle;
import data.Store;
import io.vertx.core.Handler;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.mqtt.MqttEndpoint;

public class EndPointHandler implements Handler<MqttEndpoint>{

  private final static Logger LOGGER = LoggerFactory.getLogger(EndPointHandler.class);

  @Override
  public void handle(MqttEndpoint mqttEndpoint) {
    LOGGER.info("Received request to connect from " + mqttEndpoint.clientIdentifier());

    Store.getMqttEndpoints().put(mqttEndpoint.clientIdentifier(), mqttEndpoint);
    mqttEndpoint.accept();

    mqttEndpoint.disconnectHandler(new DisconnectHandler(mqttEndpoint));
    mqttEndpoint.publishHandler(new PublishHandler(mqttEndpoint));
    mqttEndpoint.subscribeHandler(new SubscribeHandler(mqttEndpoint));
  }
}
