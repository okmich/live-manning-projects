package com.smarthome.webapp;

import data.AdminUser;
import data.MongoStore;
import io.reactivex.rxjava3.core.Completable;
import io.vertx.ext.bridge.PermittedOptions;
import io.vertx.ext.web.handler.sockjs.SockJSBridgeOptions;
import io.vertx.rxjava3.core.AbstractVerticle;
import io.vertx.rxjava3.core.http.HttpServer;
import io.vertx.rxjava3.ext.web.Router;
import io.vertx.rxjava3.ext.web.handler.BodyHandler;
import io.vertx.rxjava3.ext.web.handler.StaticHandler;
import io.vertx.rxjava3.ext.web.handler.sockjs.SockJSHandler;
import jwt.JwtHelper;

import java.util.List;
import java.util.Optional;

import static com.smarthome.webapp.handlers.RequestHandlers.*;

public class MainVerticle extends AbstractVerticle {

  private String validUserName;
  private String validPassword;

  private String sockJsAddress;

  private SockJSBridgeOptions bridgeOptions;

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
    var staticPath = Optional.ofNullable(System.getenv("STATIC_PATH")).orElse("/*");
    sockJsAddress = Optional.ofNullable(System.getenv("SOCK_JS_ADDRESS")).orElse("service.message");

    // Initialize the connection to the MongoDb database
    MongoStore.initialize(vertx, "mongodb://" + mongoHost + ":" + mongoPort, mongoDb);

    AdminUser adminUser = new AdminUser(validUserName, validPassword);

    JwtHelper jwtHelper = JwtHelper.getInstance(vertx, List.of(publicKeyPath, privateKeyPath));
    var jwtHandler = jwtHelper.getHandler();

    // =============== Begin SockJS ===============
//    Router sockJsRouter = Router.router(vertx);
    SockJSHandler sockJSHandler = SockJSHandler.create(vertx);
    SockJSBridgeOptions bridgeOptions = new SockJSBridgeOptions().addOutboundPermitted(
      new PermittedOptions().setAddress(sockJsAddress)
    ).addInboundPermitted(
      new PermittedOptions().setAddress(sockJsAddress)
    );
    Router sockJsRouter = sockJSHandler.bridge(bridgeOptions, bridgeEvent -> {
      System.out.println("websocket event:" + bridgeEvent.type());
      bridgeEvent.complete(true);
    });
    // =============== End Of SockJS ===============

    // Serving static resources
    var staticHandler = StaticHandler.create();
    staticHandler.setCachingEnabled(false);

    var router = Router.router(vertx);

    router.route().handler(BodyHandler.create());
    router.post("/authenticate").handler(authenticateHandler(jwtHelper, adminUser));
    router.get("/say-hello").handler(jwtHandler).handler(sayHelloHandler());
    router.get("/disconnect").handler(jwtHandler).handler(disconnectHandler());
    router.get(staticPath).handler(staticHandler);
    router.route("/eventbus/*").subRouter(sockJsRouter);

    // =============== Start the http server  ===============
    this.httpserver = vertx.createHttpServer().requestHandler(router);
    return httpserver
      .rxListen(httpPort)
      .doOnSuccess(ok -> {
        System.out.println("HTTP server started on port " + httpPort);
        startStreaming(3000);
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

  private void startStreaming(int delay) {
    vertx.setPeriodic(delay, longValue -> {
      MongoStore.getLastDevicesMetricsFlowable(5)
        .subscribe(doc -> {
          vertx.eventBus().publish(sockJsAddress, doc.encodePrettily());
        }, throwable -> {
          System.out.println(throwable.getMessage());
        });
    });
  }
}
