@echo off
set DEVICE_TYPE=mqtt
set DEVICE_LOCATION=garden
set DEVICE_ID=MQTT-GD-8888
set MQTT_HOST=mqtt-server
set MQTT_PORT=1883
set MQTT_TOPIC=house
java -jar .\target\smartdevice-1.0.0-SNAPSHOT-fat.jar ;
