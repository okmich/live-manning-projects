package com.okmich.livemanning.gridmonitoring.services.kafka;


import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.okmich.livemanning.gridmonitoring.data.entities.Device;
import com.okmich.livemanning.gridmonitoring.data.repo.DeviceRepository;
import manning.devices.raw.m1.RawRecord;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import java.io.IOException;
import java.util.Collections;
import java.util.Map;

@Component
public class RawDeviceStreamListener {
    private final ObjectMapper MAPPER = new ObjectMapper();
    @Autowired
    private DeviceRepository deviceRepository;

    @KafkaListener(topics = "${kafka.event.topic}", groupId = "main_group")
    public void listenForRawRecords(RawRecord record) {
        try {
            var jsonNode = recordToJson(record);
            var records = MAPPER.convertValue(jsonNode, Map.class);
            records.remove("device_id");
            var device = new Device(
                    record.getUuid().toString(),
                    records,
                    record.getArrivalTimeMs()
            );
            deviceRepository.save(device);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public JsonNode recordToJson(RawRecord record) throws IOException {
        return MAPPER.readTree(record.getBody().array());
    }
}
