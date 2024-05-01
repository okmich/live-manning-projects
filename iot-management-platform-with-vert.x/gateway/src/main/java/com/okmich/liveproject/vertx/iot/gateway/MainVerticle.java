package com.okmich.liveproject.vertx.iot.gateway;

import discovery.DiscoveryManager;
import http.DevicesHealth;
import http.Registration;
import io.vertx.core.AbstractVerticle;
import io.vertx.core.Handler;
import io.vertx.core.Promise;
import io.vertx.core.http.HttpMethod;
import io.vertx.core.http.HttpServer;
import io.vertx.core.http.HttpServerRequest;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.core.json.JsonObject;
import io.vertx.ext.web.Router;
import io.vertx.ext.web.RoutingContext;
import io.vertx.ext.web.client.WebClient;
import io.vertx.ext.web.handler.BodyHandler;
import io.vertx.servicediscovery.rest.ServiceDiscoveryRestEndpoint;
import mqtt.MqttManager;

import java.util.Optional;

public class MainVerticle extends AbstractVerticle {
  private final Logger LOGGER = LoggerFactory.getLogger(MainVerticle.class);
    private Registration registration;

  @Override
  public void start(Promise<Void> startPromise) throws Exception {

    var mqttClientId = Optional.ofNullable(System.getenv("MQTT_CLIENT_ID")).orElse("gateway");
    var mqttPort = Integer.parseInt(Optional.ofNullable(System.getenv("MQTT_PORT")).orElse("1883"));
    var mqttHost = Optional.ofNullable(System.getenv("MQTT_HOST")).orElse("mqtt-server");
    var redisHost = Optional.ofNullable(System.getenv("REDIS_HOST")).orElse("redis-server");
    var redisPort = Integer.parseInt(Optional.ofNullable(System.getenv("REDIS_PORT")).orElse("6379"));
    var authenticationToken = Optional.ofNullable(System.getenv("GATEWAY_TOKEN")).orElse("secret");
    var gatewayHttPort = Integer.parseInt(Optional.ofNullable(System.getenv("GATEWAY_HTTP_PORT")).orElse("9090"));
    var healthCheckFreq = Integer.parseInt(Optional.ofNullable(System.getenv("CHECK_FREQUENCY")).orElse("10000"));

    //initialize the mqtt client
    MqttManager mqttManager = new MqttManager(mqttClientId, mqttHost, mqttPort);
    mqttManager.startAndConnectMqttClient(vertx)
      .onSuccess(success -> {
        var serviceDiscovery = new DiscoveryManager(redisHost, redisPort).getRedisServiceDiscoveryInstance(vertx);
        registration = new Registration(serviceDiscovery, authenticationToken);

        //initialize the router
        Router router = Router.router(vertx);
        router.route().handler(BodyHandler.create());
        router.post("/register").handler(registration::handlerRegister);
        router.route().handler(this::handleInvalidRequest);

        HttpServer server = vertx.createHttpServer();
        server.requestHandler(router);

        //initialize the service discovery
        ServiceDiscoveryRestEndpoint serviceDiscoveryRestEndpoint = ServiceDiscoveryRestEndpoint.create(router, serviceDiscovery);

        // Initialize a Periodic Task To Read the Values of the Devices.
        var webClient = WebClient.create(vertx);
        DevicesHealth devicesHealth = new DevicesHealth(serviceDiscovery, webClient);
        vertx.setPeriodic(healthCheckFreq, devicesHealth::handleDevicesHealthCheck);

        // Create and start the http server
        vertx.createHttpServer().requestHandler(router)
          .listen(gatewayHttPort)
          .onFailure(this.failOnError(startPromise))
          .onSuccess(httpServerOk -> {
            startPromise.complete();
            LOGGER.info("Gateway HTTP server started on port " + gatewayHttPort);
          });
      })
      .onFailure(this.failOnError(startPromise));
  }


  private void handleInvalidRequest(RoutingContext routingContext) {
    routingContext.response().setStatusCode(404).end(JsonObject.of("message", "Resource Not Found").encode());
  }

  private Handler<Throwable> failOnError(Promise<Void> promise){
    return throwable -> {
      LOGGER.error(throwable.getMessage(), throwable);
      promise.fail(throwable.getMessage());
    };
  }

  @Override
  public void stop() throws Exception {
  }

}
