@echo off
set MONGO_PORT=27017
set MONGO_HOST=mongodb-server
set MONGO_BASE_NAME=smarthome_db
set MQTT_HOST=mqtt-server
set MQTT_PORT=8883
set MQTT_KEY=certificates/vert-x-mqtt-server.key
set MQTT_CERT=certificates/vert-x-mqtt-server.crt
java -jar .\target\broker-1.0.0-SNAPSHOT-fat.jar
