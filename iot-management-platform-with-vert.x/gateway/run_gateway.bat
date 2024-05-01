@echo off
set GATEWAY_TOKEN=smart.home
set GATEWAY_SSL=false
set REDIS_HOST=redis-server
set MQTT_HOST=mqtt-server
java -jar .\target\gateway-1.0.0-SNAPSHOT-fat.jar
