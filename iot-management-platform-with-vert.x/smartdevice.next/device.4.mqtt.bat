@echo off
set DEVICE_TYPE=mqtt
set DEVICE_LOCATION=attic
set DEVICE_ID=MQTT-AT-XX12
set MQTT_HOST=vert-x-mqtt-server
set MQTT_PORT=8883
set MQTT_TOPIC=house
set MQTT_KEY=certificates/vert-x-mqtt-server.key
set MQTT_CERT=certificates/vert-x-mqtt-server.crt
java -jar .\target\smartdevice-1.0.0-SNAPSHOT-fat.jar ;
