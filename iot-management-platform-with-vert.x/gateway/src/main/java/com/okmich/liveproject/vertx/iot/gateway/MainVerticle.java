package com.okmich.liveproject.vertx.iot.gateway;

import com.okmich.liveproject.vertx.iot.gateway.http.Registration;
import io.vertx.core.AbstractVerticle;
import io.vertx.core.Handler;
import io.vertx.core.Promise;
import io.vertx.core.http.HttpMethod;
import io.vertx.core.http.HttpServer;
import io.vertx.core.http.HttpServerRequest;
import io.vertx.ext.web.Router;
import io.vertx.ext.web.RoutingContext;

public class MainVerticle extends AbstractVerticle {

  private Registration registration = new Registration();

  @Override
  public void start(Promise<Void> startPromise) throws Exception {
    vertx.setPeriodic();

    HttpServer server = vertx.createHttpServer();

    Router router = Router.router(vertx);
    router.post("/register").handler(registration::handlerRegister);
    router.route().handler(this::handleInvalidRequest);

    server.requestHandler(router);

    server.listen(8888);
    System.out.println("Server up. Awaiting request ...");
  }


  private void handleInvalidRequest(RoutingContext routingContext) {
    routingContext.response().setStatusCode(404).end();
  }

  @Override
  public void stop() throws Exception {
  }

}
