package com.okmich.liveproject.vertx.iot.gateway.http;

import io.vertx.core.http.HttpServerRequest;
import io.vertx.ext.web.RoutingContext;

public class Registration {

  public void handlerRegister(RoutingContext routingContext) {
    HttpServerRequest request = routingContext.request();

    request.response().setStatusCode(200).end();
  }
}
