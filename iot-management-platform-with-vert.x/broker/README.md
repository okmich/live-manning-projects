# Broker

## Build

To build and package the gateway application, use the below command:
```bash
.\mvnw clean package
```

To tun the application without packaging it, use the command below:
```bash
.\mvnw clean compile exec:java
```

## Prerequisites before testing the MQTT Broker

**First**, add `vert-x-mqtt-server` to your `hosts` file. The line should look like something like that:

```
127.0.0.1  localhost redis-server mongodb-server mqtt-server fake-device gateway.home.smart devices.home.smart vert-x-mqtt-server
```

**Then**, you will need a MongoDb DataBase and a Redis Database

In the Docker Compose project, start the containers:

```bash
docker-compose up
```

## Start the new Broker

```batch
@echo off
set MONGO_PORT=27017
set MONGO_HOST=mongodb-server
set MONGO_BASE_NAME=smarthome_db
set MQTT_PORT=1884
java -jar .\target\broker-1.0.0-SNAPSHOT-fat.jar
```

> **Remark**: the broker is listening on `1884` to avoid collisions with the Mosquitto broker (or deactivate Mosquitto without Docker Compose)

## Start an MQTT Device

Go to the Devices project, and start a new MQTT device:
```bash
@echo off
set DEVICE_TYPE=mqtt
set DEVICE_LOCATION=garden
set DEVICE_ID=MQTT-GD-1968
set MQTT_HOST=vert-x-mqtt-server
set MQTT_PORT=1884
set MQTT_TOPIC=house
java -jar target/smartdevice-1.0.0-SNAPSHOT-fat.jar;
```

> **Remarks**:
> - precise the type of the device: `DEVICE_TYPE="mqtt"`
> - precise the appropriate MQTT port number: `MQTT_PORT=1884`
> - precise the "local DNS name" of the broker: `MQTT_HOST="vert-x-mqtt-server"`  (like the one in the `hosts` file)

At this moment, you will see the data from the MQTT device displayed in the terminal output of the MQTT broker

## Start the Gateway

Go to the Gateway project and launch the gateway:
```bash
GATEWAY_TOKEN="smart.home" \
GATEWAY_SSL="false" \
GATEWAY_HTTP_PORT=9090 \
REDIS_HOST="redis-server" \
REDIS_PORT=6379 \
MQTT_HOST="vert-x-mqtt-server" \
MQTT_PORT=1884 \
java -jar target/gateway-1.0.0-SNAPSHOT-fat.jar
```

## Start an HTTP device

Go to the Devices project, and start a new HTTP device:
```bash
HTTP_PORT="8081" \
DEVICE_TYPE="http" \
DEVICE_HOSTNAME="devices.home.smart" \
DEVICE_LOCATION="bedroom" \
DEVICE_ID="AX3345" \
GATEWAY_TOKEN="smart.home" \
GATEWAY_DOMAIN="gateway.home.smart" \
GATEWAY_HTTP_PORT=9090 \
java -jar target/smartdevice-1.0.0-SNAPSHOT-fat.jar ;
```

> **Remark**: precise the type of the device: `DEVICE_TYPE="http"`

At this moment, you will see the data from the HTTP device displayed in the terminal output of the MQTT broker
