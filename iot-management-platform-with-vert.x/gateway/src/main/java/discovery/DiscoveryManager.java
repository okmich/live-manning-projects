package discovery;

import io.vertx.core.Vertx;
import io.vertx.core.json.JsonObject;
import io.vertx.servicediscovery.ServiceDiscovery;
import io.vertx.servicediscovery.ServiceDiscoveryOptions;

public class DiscoveryManager {

  private final String redisHost;
  private final int redisPort;

  public DiscoveryManager(String redisHost, int redisPort) {
    this.redisHost = redisHost;
    this.redisPort = redisPort;
  }


  public ServiceDiscovery getRedisServiceDiscoveryInstance(Vertx vertx) {
    var redisConnectionString = "redis://" + this.redisHost + ":" + this.redisPort;
    return ServiceDiscovery.create(vertx, new ServiceDiscoveryOptions()
      .setBackendConfiguration(
        new JsonObject()
          .put("connectionString", redisConnectionString)
          .put("key", "devices_records")
      ));
  }
}
