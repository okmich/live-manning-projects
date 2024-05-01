@echo off
set HTTP_PORT=8081
set DEVICE_TYPE=http
set DEVICE_HOSTNAME=devices.home.smart
set DEVICE_LOCATION=bedroom
set DEVICE_ID=AX3345
set GATEWAY_TOKEN=smart.home
set GATEWAY_DOMAIN=gateway.home.smart
set GATEWAY_HTTP_PORT=9090
java -jar target\smartdevice-1.0.0-SNAPSHOT-fat.jar
