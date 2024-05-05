package com.smarthome.webapp;

import com.smarthome.webapp.handlers.RequestHandlers;
import data.AdminUser;
import data.MongoStore;
import io.reactivex.rxjava3.core.Completable;
import io.vertx.rxjava3.core.AbstractVerticle;
import io.vertx.rxjava3.core.http.HttpServer;
import io.vertx.rxjava3.ext.web.Router;
import io.vertx.rxjava3.ext.web.handler.BodyHandler;
import jwt.JwtHelper;

import java.util.List;
import java.util.Optional;

public class MainVerticle extends AbstractVerticle {

  private String validUserName;
  private String validPassword;

  private HttpServer httpserver;

  public Completable rxStart() {
    var validUserName = Optional.ofNullable(System.getenv("ADMIN_NAME")).orElse("root");
    var validPassword = Optional.ofNullable(System.getenv("ADMIN_PASSWORD")).orElse("admin");
    var httpPort = Integer.parseInt(Optional.ofNullable(System.getenv("HTTP_PORT")).orElse("8080"));
    var privateKeyPath = Optional.ofNullable(System.getenv("PRIVATE_KEY_PATH")).orElse("./private_key.pem");
    var publicKeyPath = Optional.ofNullable(System.getenv("PUBLIC_KEY_PATH")).orElse("./public_key.pem");
    var mongoHost = Optional.ofNullable(System.getenv("MONGO_HOST")).orElse("mongo-server");
    var mongoPort = Integer.parseInt(Optional.ofNullable(System.getenv("MONGO_PORT")).orElse("27017"));
    var mongoDb = Optional.ofNullable(System.getenv("MONGO_DB")).orElse("smarthome_db");

    // Initialize the connection to the MongoDb database
    MongoStore.initialize(vertx, "mongodb://"+mongoHost+":"+mongoPort, mongoDb);

    AdminUser adminUser = new AdminUser(validUserName, validPassword);

    JwtHelper jwtHelper = JwtHelper.getInstance(vertx, List.of(publicKeyPath, privateKeyPath));
    var jwtHandler = jwtHelper.getHandler();
    var router = Router.router(vertx);

    router.route().handler(BodyHandler.create());
    router.post("/authenticate").handler(RequestHandlers.authenticateHandler(jwtHelper, adminUser));
    router.get("/say-hello").handler(jwtHandler).handler(RequestHandlers.sayHelloHandler());
    router.get("/disconnect").handler(jwtHandler).handler(RequestHandlers.disconnectHandler());

    this.httpserver = vertx.createHttpServer().requestHandler(router);
    return httpserver
      .rxListen(httpPort)
      .doOnSuccess(ok -> {
        System.out.println("HTTP server started on port " + httpPort);
      })
      .doOnError(error -> {
        System.out.println(error.getCause().getMessage());
      })
      .ignoreElement();
  }

  public Completable rxStop() {
    if (this.httpserver != null) {
      this.httpserver.rxClose();
    }
    return super.rxStop();
  }

}
