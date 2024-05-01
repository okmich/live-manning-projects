package http;

public class RegistrationData {

  private String category;
  private String id;
  private String position;
  private String host;
  private int port;

  public RegistrationData() {
  }

  public RegistrationData(String category, String host, String position, String ip, int port) {
    this.category = category;
    this.id = id;
    this.position = position;
    this.host = ip;
    this.port = port;
  }

  public String getCategory() {
    return category;
  }

  public void setCategory(String category) {
    this.category = category;
  }

  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }

  public String getPosition() {
    return position;
  }

  public void setPosition(String position) {
    this.position = position;
  }

  public String getHost() {
    return host;
  }

  public void setHost(String host) {
    this.host = host;
  }

  public int getPort() {
    return port;
  }

  public void setPort(int port) {
    this.port = port;
  }
}
