package com.smarthome.webapp.handlers;

import data.AdminUser;
import io.vertx.core.Handler;
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
      System.out.println(routingContext.user().principal());
      String subject = routingContext.user().principal().getString("sub");
      String greetingMessage = routingContext.user().principal().getString("greetingMessage");
      routingContext.json(new JsonObject().put("greetingMessage", greetingMessage).put("subject",subject));
    };
  }

  public static Handler<RoutingContext> disconnectHandler() {
    return routingContext -> {
      routingContext.json(new JsonObject().put("message", "disconnected"));
    };
  }
}
