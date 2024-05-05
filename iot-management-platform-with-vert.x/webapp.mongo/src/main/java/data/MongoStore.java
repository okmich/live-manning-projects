package data;

import io.reactivex.rxjava3.core.Flowable;
import io.vertx.core.json.JsonObject;
import io.vertx.ext.mongo.FindOptions;
import io.vertx.rxjava3.core.Vertx;
import io.vertx.rxjava3.core.streams.ReadStream;
import io.vertx.rxjava3.ext.mongo.MongoClient;
import io.vertx.rxjava3.ext.web.RoutingContext;

import java.util.NoSuchElementException;

public class MongoStore {

  private static MongoClient mongoClient;

  public static void initialize(Vertx vertx, String connectionString, String dataBaseName) {
    mongoClient = MongoClient.create(
      vertx,
      new JsonObject()
        .put("db_name",dataBaseName)
        .put("useObjectId", false)
        .put("connection_string", connectionString)
    );
  }

  private void completeFetchRequest(RoutingContext ctx, JsonObject json) {
    ctx.response()
      .putHeader("Content-Type", "application/json")
      .end(json.encode());
  }

  public static Flowable<JsonObject> getLastDevicesMetricsFlowable(int howMany) {
    JsonObject query = new JsonObject();
    var options = new FindOptions();
    options.setSort(new JsonObject().put("_id",-1));
    options.setLimit(howMany);

    ReadStream<JsonObject> devices = mongoClient.findBatchWithOptions("devices", query, options);
    return devices.toFlowable();
  }
}
