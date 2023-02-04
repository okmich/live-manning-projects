package com.smarthome.gateway;

import discovery.DiscoveryManager;
import http.DevicesHealth;
import http.Registration;
import io.vertx.core.AbstractVerticle;
import io.vertx.core.Handler;
import io.vertx.core.Promise;
import io.vertx.core.Vertx;
import io.vertx.core.http.HttpServerOptions;
import io.vertx.core.net.PemKeyCertOptions;
import io.vertx.ext.web.Router;
import io.vertx.ext.web.client.WebClient;
import io.vertx.ext.web.handler.BodyHandler;
import io.vertx.servicediscovery.rest.ServiceDiscoveryRestEndpoint;
import mqtt.MqttManager;

import java.util.Optional;

public class MainVerticle extends AbstractVerticle {

  public static void main(String[] args) {
    Vertx vertx = Vertx.vertx();
    vertx.deployVerticle(new MainVerticle());
  }

  @Override
  public void stop(Promise<Void> stopPromise) throws Exception {

    stopPromise.complete();
  }

  @Override
  public void start(Promise<Void> startPromise) throws Exception {

    //Initialize a periodic task to read the values of the devices.
    vertx.setPeriodic(5000, id -> {
      //do the work here
    });
  }

}
