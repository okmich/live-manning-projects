package com.okmich.liveproject.vertx.iot.gateway.smartdevice;

import devices.HttpDevice;
import io.vertx.core.AbstractVerticle;
import io.vertx.core.Promise;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.core.json.JsonObject;
import io.vertx.ext.web.Router;
import sensors.HumiditySensor;
import sensors.TemperatureSensor;
import sensors.eCO2Sensor;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

public class MainVerticle extends AbstractVerticle {

  private final Logger LOGGER = LoggerFactory.getLogger(MainVerticle.class);

  @Override
  public void start(Promise<Void> startPromise) throws Exception {
    String deviceType = Optional.ofNullable(System.getenv("DEVICE_TYPE")).orElse("http");
    String deviceHostName = Optional.ofNullable(System.getenv("DEVICE_HOSTNAME")).orElse("localhost");
    String deviceLocation = Optional.ofNullable(System.getenv("DEVICE_LOCATION")).orElse("living room");
    String deviceId = Optional.ofNullable(System.getenv("DEVICE_ID")).orElse(UUID.randomUUID().toString());
    String gatewayToken = Optional.ofNullable(System.getenv("GATEWAY_TOKEN")).orElse("smart.home");
    String gatewayHttpPort = Optional.ofNullable(System.getenv("GATEWAY_HTTP_PORT")).orElse("9090");
    int httpPort = Integer.parseInt(Optional.ofNullable(System.getenv("HTTP_PORT")).orElse("8081"));

    if (deviceType.equals("http")) {
      HttpDevice httpDevice = new HttpDevice(deviceId, deviceLocation, httpPort,
        List.of(new TemperatureSensor(), new HumiditySensor(), new eCO2Sensor())
      );

      var registrationRequest = httpDevice.createRegisterToGatewayRequest(
        vertx, deviceHostName, Integer.parseInt(gatewayHttpPort), false, gatewayToken);

      var registrationRequestPayload = JsonObject.of(
        "category", httpDevice.getCategory(),
        "id", httpDevice.getId(),
        "position", httpDevice.getPosition(),
        "host", httpDevice.getHostName(),
        "port", httpDevice.getPort()
      );

      registrationRequest.sendJsonObject(registrationRequestPayload).onSuccess(
        response -> {
          if (response.statusCode() == 200) {
            LOGGER.info("Device registered with gateway successfully ");
            httpDevice.setConnectedToGateway(true);
            startServer(startPromise, httpDevice, httpPort);
          } else {
            var errorMessage = "Device registration with gateway failed with response code " + response.statusCode();
            LOGGER.error(errorMessage);
            startPromise.fail(errorMessage);
          }
        }
      ).onFailure(
        error -> {
          var errorMessage = "Connection to the Gateway failed: " + error.getMessage();
          LOGGER.error(errorMessage);
          startPromise.fail(errorMessage);
        }
      );
    } else {
      throw new UnsupportedOperationException("To be implemented in another milestone");
    }
  }

  private void startServer(Promise<Void> startPromise, HttpDevice httpDevice, int httpPort) {
    Router router = httpDevice.createRouter(vertx);
    router.route("/").handler(routingContext -> {
        LOGGER.info(httpDevice.jsonValue());
        routingContext.json(httpDevice.jsonValue());
      }
    );

    httpDevice.createHttpServer(vertx, router)
      .listen(httpPort)
      .onFailure(error -> startPromise.fail(error.getCause()))
      .onSuccess(ok -> {
        startPromise.complete();
        System.out.println("Device: HTTP server started on port " + httpPort);
      });
  }
}
