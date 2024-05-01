package handlers;

import data.Store;
import io.vertx.core.Handler;
import io.vertx.mqtt.MqttEndpoint;

public class DisconnectHandler implements Handler<Void> {

  private final MqttEndpoint mqttEndpoint;

  public DisconnectHandler(MqttEndpoint mqttEndpoint){
    this.mqttEndpoint = mqttEndpoint;
  }

  @Override
  public void handle(Void unused) {
    Store.getMqttEndpoints().remove(mqttEndpoint.clientIdentifier());
  }
}
