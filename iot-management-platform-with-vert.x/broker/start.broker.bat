@echo off
set MONGO_PORT=27017
set MONGO_HOST=mongodb-server
set MONGO_BASE_NAME=smarthome_db
set MQTT_HOST=mqtt-server
set MQTT_PORT=1884
java -jar .\target\broker-1.0.0-SNAPSHOT-fat.jar
