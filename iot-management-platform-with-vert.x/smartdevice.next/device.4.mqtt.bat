@echo off
set DEVICE_TYPE=mqtt
set DEVICE_LOCATION=attic
set DEVICE_ID=MQTT-AT-XX12
set MQTT_HOST=mqtt-server
set MQTT_PORT=1883
set MQTT_TOPIC=house
java -jar .\target\smartdevice-1.0.0-SNAPSHOT-fat.jar ;
