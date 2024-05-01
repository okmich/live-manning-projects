package http;

import io.vertx.core.Handler;
import io.vertx.core.MultiMap;
import io.vertx.core.http.HttpServerRequest;
import io.vertx.core.impl.logging.Logger;
import io.vertx.core.impl.logging.LoggerFactory;
import io.vertx.core.json.JsonObject;
import io.vertx.ext.web.RoutingContext;
import io.vertx.servicediscovery.Record;
import io.vertx.servicediscovery.ServiceDiscovery;
import io.vertx.servicediscovery.Status;
import io.vertx.servicediscovery.types.HttpEndpoint;

import java.util.Optional;

public class Registration {

  private final Logger LOGGER = LoggerFactory.getLogger(Registration.class);

  private final ServiceDiscovery serviceDiscovery;
  private final String authToken;

  public Registration(ServiceDiscovery discovery, String validAuthToken) {
    this.serviceDiscovery = discovery;
    this.authToken = validAuthToken;
  }

  private RegistrationData marshalToObject(RoutingContext routingContext) {
    try {
      var jsonObject = routingContext.body().asJsonObject();
      return jsonObject.mapTo(RegistrationData.class);
    } catch (Exception ex) {
      LOGGER.error(ex.getMessage(), ex);
      return null;
    }
  }

  public void handlerRegister(RoutingContext routingContext) {
    HttpServerRequest request = routingContext.request();
    if (validateRequestHeader(request.headers())) {
      RegistrationData registrationData = marshalToObject(routingContext);
      if (registrationData != null) {
        Record record = HttpEndpoint.createRecord(registrationData.getId(),
            registrationData.getHost(), registrationData.getPort(), "/")
          .setMetadata(
            JsonObject.of("category", registrationData.getCategory(),
              "position", registrationData.getPosition())
          );
        this.serviceDiscovery.getRecord(rec -> rec.getName().equals(registrationData.getId()))
          .onSuccess(recObj -> {
            if (recObj == null) { // new, so we publish
              this.serviceDiscovery.publish(record)
                .onSuccess(this.onDiscoveryPublishSuccess(request))
                .onFailure(this.basicOnFailure());
            } else { // existing, so we update
              recObj.setStatus(Status.UP);
              this.serviceDiscovery.update(recObj)
                .onSuccess(this.onDiscoveryPublishSuccess(request))
                .onFailure(this.basicOnFailure());
            }
          })
          .onFailure(this.basicOnFailure());
      } else {
        request.response().setStatusCode(400).end(JsonObject.of("message", "Bad Request").encode());
      }
    } else {
      request.response().setStatusCode(403).end(JsonObject.of("message", "bad-token").encode());
    }
  }

  private boolean validateRequestHeader(MultiMap headers) {
    var optionalToken = Optional.ofNullable(headers.get("smart-token"));
    var token = optionalToken.orElse("");
    return token.equals(authToken);
  }

  private Handler<Record> onDiscoveryPublishSuccess(HttpServerRequest request) {
    return res ->
      request.response().setStatusCode(200).end(
        JsonObject.of("message", "Registration done").encode());
  }

  private Handler<Throwable> basicOnFailure() {
    return throwable -> {
      LOGGER.error(throwable.getMessage(), throwable);
    };
  }
}
