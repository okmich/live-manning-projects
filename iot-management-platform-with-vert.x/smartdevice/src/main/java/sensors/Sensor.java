package sensors;

/*
The sine and cosine functions can be used to model fluctuations in temperature data throughout the year.
An equation that can be used to model these data is of the form:

  y = A cos B(x - C) + D,

where A,B,C,D, are constants, y is the temperature in °C and x is the month (1–12).

A = amplitude = (ymax - ymin)/2
B = 2π/12
C = units translated to the right
D = ymin + amplitude = units translated up
*/

import io.vertx.core.json.Json;
import io.vertx.core.json.JsonObject;

import java.time.LocalTime;

public interface Sensor {
  String getName();
  String getUnit();
  double getLevel(int t);

  default JsonObject jsonValue(){
    var localTime = LocalTime.now();
    var t = localTime.getSecond()  / 5 + 1;

    return JsonObject.of(getName(), JsonObject.of("unit", getUnit(), "value", getLevel(t)));
  }

  default double simulate(int t, double minValue, double maxValue) {
    var amplitude = (maxValue - minValue) / 2;
    var unitsTranslatedUp = minValue + amplitude;
    var minBound = 1;
    var maxBound = 5;
    var unitsTranslatedToTheRight = Math.floor(Math.random()*(maxBound-minBound+1)) + minBound;

    return amplitude * Math.cos((t-unitsTranslatedToTheRight) * Math.PI/6) + unitsTranslatedUp;
  }
}

