package com.smarthome.webapp.handlers;

import data.AdminUser;
import data.MongoStore;
import io.vertx.core.Handler;
import io.vertx.core.json.JsonArray;
import io.vertx.core.json.JsonObject;
import io.vertx.rxjava3.ext.web.RoutingContext;
import jwt.JwtHelper;

public class RequestHandlers {

  public static Handler<RoutingContext> authenticateHandler(JwtHelper jwtHelper, AdminUser adminUser) {
    return routingContext -> {
      var payload = routingContext.body().asJsonObject();
      if (payload == null){
        routingContext
          .response().setStatusCode(400)
          .end(new JsonObject().put("message", "Bad Request").encode());
        return;
      }

      var user = payload.getString("username");
      var pwd = payload.getString("password");

      routingContext.response()
        .putHeader("Content-Type", "application/jwt");

      if (user.equals(adminUser.getUserName()) && pwd.equals(adminUser.getPassword())) {
        var token = jwtHelper.generateToken(user, new JsonObject().put("greetingMessage", "Welcome " + user));
        adminUser.setAuthenticated(true);
        routingContext
          .response().setStatusCode(200)
          .end(
            JsonObject.of(
              "token", token,
              "greetingMessage", "Welcome " + adminUser.getUserName()
            ).encode()
          );
      } else {
        adminUser.setAuthenticated(false);
        routingContext
          .response().setStatusCode(401)
          .send(new JsonObject().put("message", "Bad JWT Token").encode());
      }
    };
  }

  public static Handler<RoutingContext> sayHelloHandler() {
    return routingContext -> {
      String subject = routingContext.user().principal().getString("sub");

      var documents = new JsonArray();
      MongoStore.getLastDevicesMetricsFlowable(10)
        .subscribe(documents::add, throwable -> {
          System.out.println(throwable.getMessage());
          routingContext.response().setStatusCode(500);
          routingContext.json(new JsonObject().put("error", throwable.getMessage()));
        }, () -> {
          System.out.println("All documents found");
          System.out.println(documents);
          routingContext.json(JsonObject.of("subject",subject, "documents", documents));
        });
    };
  }

  public static Handler<RoutingContext> disconnectHandler() {
    return routingContext -> {
      routingContext.json(new JsonObject().put("message", "disconnected"));
    };
  }
}
