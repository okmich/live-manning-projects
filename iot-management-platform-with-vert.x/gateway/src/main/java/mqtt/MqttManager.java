package mqtt;

import io.vertx.circuitbreaker.CircuitBreaker;
import io.vertx.circuitbreaker.CircuitBreakerOptions;
import io.vertx.core.Future;
import io.vertx.core.Vertx;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.core.net.PemKeyCertOptions;
import io.vertx.core.net.PemTrustOptions;
import io.vertx.mqtt.MqttClient;
import io.vertx.mqtt.MqttClientOptions;
import io.vertx.mqtt.messages.MqttConnAckMessage;

public class MqttManager {

  private final Logger LOGGER = LoggerFactory.getLogger(MqttManager.class);

  private final String mqttHost;
  private final int mqttPort;
  private final MqttClientOptions mqttOptions;
  private MqttClient mqttClient;
  private CircuitBreaker breaker;

  public MqttManager(String mqttClientId, String mqttHost, int mqttPort, String mqttCert, String mqttKey) {
    this.mqttHost = mqttHost;
    this.mqttPort = mqttPort;
    this.mqttOptions = new MqttClientOptions()
      .setClientId(mqttClientId)
      .setPemTrustOptions(
        new PemTrustOptions().addCertPath(mqttCert)
      )
      .setHostnameVerificationAlgorithm("HTTPS")
      .setKeyCertOptions(
        new PemKeyCertOptions()
          .setKeyPath(mqttKey)
          .setCertPath(mqttCert)
      )
      .setSsl(true);

  }

  public MqttClient getMqttClient() {
    return mqttClient;
  }

  //   get a circuit breaker
  private CircuitBreaker getBreaker(Vertx vertx) {
    if (breaker == null) {
      breaker = CircuitBreaker.create("my-gateway-breaker", vertx,
        new CircuitBreakerOptions()
          .setMaxFailures(5) // number of failure before opening the circuit
          .setTimeout(10_000) // consider a failure if the operation does not succeed in time
          .setFallbackOnFailure(false) // do we call the fallback on failure
          .setResetTimeout(20_000) // time spent in open state before attempting to re-try
      );
    }
    return breaker;
  }

  // create and connect the MQTT client "in" a Circuit Breaker
  public Future<MqttConnAckMessage> startAndConnectMqttClient(Vertx vertx) {
    return getBreaker(vertx).execute(promise -> {

      mqttClient = MqttClient.create(vertx, this.mqttOptions).exceptionHandler(throwable -> {
        LOGGER.error(throwable.getMessage(), throwable);
      }).closeHandler(voidValue -> {
        LOGGER.error("Connection with broker is lost");
        startAndConnectMqttClient(vertx);
      });

      // some code executing with the breaker
      // the code reports failures or success on the given promise.
      // if this promise is marked as failed, the breaker increased the
      // number of failures

      // connect the mqttClient
      mqttClient.connect(mqttPort, mqttHost)
        .onFailure(error -> {
          LOGGER.error("MQTT " + error.getMessage());
          promise.fail("[" + error.getMessage() + "]");
        })
        .onSuccess(ok -> {
          LOGGER.error("Connection to the broker is ok");
          promise.complete();
        });
    });
  }
}

