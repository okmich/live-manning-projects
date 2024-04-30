package devices;

import communications.Http;
import io.vertx.core.json.JsonObject;
import sensors.Sensor;

import java.util.LinkedList;
import java.util.List;

public class HttpDevice implements Device, Http {

  private boolean connectedToGateway;

  private final String position;
  private final LinkedList<Sensor> sensors;
  private final int port;
  private final String id;
  private final String category = "http";

  public HttpDevice(String id, String position, int port, List<Sensor> sensors) {
    this.id = id;
    this.position = position;
    this.sensors = new LinkedList<>(sensors);
    this.port = port;
    this.connectedToGateway = false;
  }

  @Override
  public boolean isConnectedToGateway() {
    return this.connectedToGateway;
  }

  @Override
  public void setConnectedToGateway(boolean value) {
    this.connectedToGateway = value;
  }

  @Override
  public int getPort() {
    return this.port;
  }

  @Override
  public LinkedList<Sensor> getSensors() {
    return this.sensors;
  }

  @Override
  public String getId() {
    return this.id;
  }

  @Override
  public String getPosition() {
    return this.position;
  }

  @Override
  public String getCategory() {
    return null;
  }

}
