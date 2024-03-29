package com.okmich.livemanning.gridmonitoring.data;


import com.fasterxml.jackson.databind.JsonNode;
import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;

@Component
public class DeviceDB {

    private static final Map<String, JsonNode> deviceDb = new HashMap<>();

    public DeviceDB() {
    }

    public void setDeviceCharging(String deviceId, JsonNode payload) {
        deviceDb.put(deviceId, payload);
    }

    public JsonNode getDeviceCharging(String deviceId) {
        return deviceDb.get(deviceId);
    }
}
