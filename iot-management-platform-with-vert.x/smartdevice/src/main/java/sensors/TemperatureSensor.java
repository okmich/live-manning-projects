package sensors;

public class TemperatureSensor implements Sensor {

  @Override
  public String getName() {
    return "temperature";
  }

  @Override
  public String getUnit() {
    return "Celsius";
  }

  @Override
  public double getLevel(int t) {
    return simulate(t, 0.0, 33.0);
  }
}
