package sensors;

public class HumiditySensor implements Sensor {

  @Override
  public String getName() {
    return "humidity";
  }

  @Override
  public String getUnit() {
    return "%";
  }

  @Override
  public double getLevel(int t) {
    return simulate(t, 30.0, 60.0);
  }
}
