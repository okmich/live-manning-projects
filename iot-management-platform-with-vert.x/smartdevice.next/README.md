# Smartdevice

## Build

To package your application:
```bash
.\mvnw clean package
```

To run your application:
```bash
.\mvnw clean compile exec:java
```

## Before starting the device (prerequisites)

You'll need of an MQTT Broker.

In the Docker Compose project, start the containers:

```bash
docker-compose up
```

Once the containers started (you'll get a running Mosquitto MQTT Broker), in a terminal, run a Mosquitto client to listen on the `house` topic:

```bash
docker exec -it mqtt-server /bin/sh
mosquitto_sub -h localhost -t house/#
```

## Start a device:

```bash
@echo off
set DEVICE_TYPE="mqtt" \
set DEVICE_LOCATION="attic" \
set DEVICE_ID="MQTT-AT-XX12" \
set MQTT_HOST="mqtt-server" \
set MQTT_PORT=1883 \
set MQTT_TOPIC="house" \
java -jar .\target\smartdevice-1.0.0-SNAPSHOT-fat.jar ;
```

