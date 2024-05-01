package communications;


import io.vertx.circuitbreaker.CircuitBreaker;
import io.vertx.circuitbreaker.CircuitBreakerOptions;
import io.vertx.core.Future;
import io.vertx.core.Vertx;
import io.vertx.mqtt.MqttClient;
import io.vertx.mqtt.messages.MqttConnAckMessage;

public interface Mqtt {
  default MqttClient createMqttClient(Vertx vertx) {
    return MqttClient.create(vertx);
  }

  Future<MqttConnAckMessage> startAndConnectMqttClient(Vertx vertx);

  default CircuitBreaker getBreaker(Vertx vertx) {
    return CircuitBreaker.create("my-mqtt-circuit-breaker", vertx,
      new CircuitBreakerOptions()
        .setMaxFailures(5)              // number of failure before opening the circuit
        .setTimeout(5_000)              // consider a failure if the operation does not succeed in time
        .setFallbackOnFailure(false)    // do we call the fallback on failure
        .setResetTimeout(10_000)        // time spent in open state before attempting to re-try
    );
  }

  String getProtocol();

  MqttClient getMqttClient();

}
