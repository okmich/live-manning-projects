@echo off
set GATEWAY_TOKEN=smart.home
set GATEWAY_SSL=false
set GATEWAY_HTTP_PORT=9090
set REDIS_HOST=redis-server
set REDIS_PORT=6379
set MQTT_HOST=vert-x-mqtt-server
set MQTT_PORT=1884
set MQTT_KEY=certificates/vert-x-mqtt-server.key
set MQTT_CERT=certificates/vert-x-mqtt-server.crt
set MQTT_TOPIC=house
java -jar .\target\gateway-1.0.0-SNAPSHOT-fat.jar
