package sensors;

public class eCO2Sensor implements Sensor {

  @Override
  public String getName() {
    return "eCO2";
  }

  @Override
  public String getUnit() {
    return "ppm";
  }

  @Override
  public double getLevel(int t) {
    return simulate(t, 900.0, 3000.0);
  }
}
