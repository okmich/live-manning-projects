package com.smarthome.broker;

import com.sun.tools.javac.Main;
import data.MongoStore;
import handlers.EndPointHandler;
import io.vertx.core.AbstractVerticle;
import io.vertx.core.Promise;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.core.net.PemKeyCertOptions;
import io.vertx.core.net.PemTrustOptions;
import io.vertx.mqtt.MqttServer;
import io.vertx.mqtt.MqttServerOptions;

import java.util.Optional;

public class MainVerticle extends AbstractVerticle {

  private final static Logger LOGGER = LoggerFactory.getLogger(MainVerticle.class);

  @Override
  public void stop(Promise<Void> stopPromise) throws Exception {
    stopPromise.complete();
  }

  @Override
  public void start(Promise<Void> startPromise) throws Exception {
    // Define the connection parameters to the MongoDB database (the connection string);
    var mongoPort = Integer.parseInt(Optional.ofNullable(System.getenv("MONGO_PORT")).orElse("27017"));
    var mongoHost = Optional.ofNullable(System.getenv("MONGO_HOST")).orElse("localhost");
    var mongoBaseName = Optional.ofNullable(System.getenv("MONGO_BASE_NAME")).orElse("smarthome_db");

    // Define the connection parameters to the MQTT server (the MQTT port).
    var mqttHost = Optional.ofNullable(System.getenv("MQTT_HOST")).orElse("localhost");
    var mqttPort = Integer.parseInt(Optional.ofNullable(System.getenv("MQTT_PORT")).orElse("1883"));
    var mqttCert = Optional.ofNullable(System.getenv("MQTT_CERT")).orElse("");
    var mqttKey = Optional.ofNullable(System.getenv("MQTT_KEY")).orElse("");

    //Create the MongoDB client.
    MongoStore.initialize(vertx, mongoHost, mongoPort, mongoBaseName);

    // Instantiate the MQTTServer.
    var mqttOptions = new MqttServerOptions().setHost(mqttHost).setPort(mqttPort)
      .setPemTrustOptions(
        new PemTrustOptions().addCertPath(mqttCert)
      )
      .setKeyCertOptions(
        new PemKeyCertOptions()
          .setKeyPath(mqttKey)
          .setCertPath(mqttCert)
      )
      .setSsl(true);;
    MqttServer mqttServer = MqttServer.create(vertx, mqttOptions);
    //Set the endpointHandler: mqttServer.endpointHandler
    mqttServer.endpointHandler(new EndPointHandler());

    //Finally, start the MQTT server (mqttServer.listen()).
    mqttServer.listen()
      .onFailure(error -> {
        System.out.println("MQTT " + error.getMessage());
      })
      .onSuccess(ok -> {
        System.out.println("MQTT broker started, listening on " + mqttPort);
      });
  }

}
