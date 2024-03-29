package com.okmich.livemanning.gridmonitoring.controller;


import com.fasterxml.jackson.databind.ObjectMapper;
import com.okmich.livemanning.gridmonitoring.data.DeviceDB;
import com.okmich.livemanning.gridmonitoring.data.repo.DeviceRepository;
import jakarta.servlet.http.HttpServletRequest;
import manning.devices.raw.m1.RawRecord;
import org.apache.kafka.clients.producer.RecordMetadata;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.time.Instant;
import java.util.Map;
import java.util.concurrent.ExecutionException;

import static org.apache.commons.compress.utils.IOUtils.toByteArray;

@RestController
@RequestMapping(value = "/device", consumes = {MediaType.APPLICATION_OCTET_STREAM_VALUE, MediaType.APPLICATION_JSON_VALUE}, produces = MediaType.APPLICATION_JSON_VALUE)
public class DeviceController {

    private final KafkaTemplate<String, RawRecord> kafkaTemplate;
    private final DeviceRepository deviceRepository;
    @Value("${kafka.event.topic}")
    private String topic;

    public DeviceController(KafkaTemplate<String, RawRecord> kafkaTemplate, DeviceRepository deviceRepository) {
        this.kafkaTemplate = kafkaTemplate;
        this.deviceRepository = deviceRepository;
    }

    @PostMapping("/send/{uuid}")
    public ResponseEntity<Map<String, Object>> send(@PathVariable("uuid") String uuid, HttpServletRequest request)
            throws IOException, ExecutionException, InterruptedException {
        ByteBuffer body = ByteBuffer.wrap(toByteArray(request.getInputStream()));
        RawRecord payload = new RawRecord(uuid, Instant.now().toEpochMilli(), body);

        var resultSendFuture = kafkaTemplate.send(topic, payload);
        RecordMetadata recordMetadata = resultSendFuture.get().getRecordMetadata();
        return ResponseEntity.ok(Map.of(
                "offset", recordMetadata.offset(),
                "partition", recordMetadata.partition(),
                "topic", recordMetadata.topic(),
                "timestamp", recordMetadata.timestamp()
        ));
    }

    @GetMapping("/state")
    public ResponseEntity<Map<String, Object>> getStatus(@RequestParam(value = "uuid", required = true) String uuid) {
        var optionalEnt = this.deviceRepository.findById(uuid);
        return optionalEnt.map(device -> ResponseEntity.ok(device.getReadings())).orElseGet(() -> ResponseEntity.notFound().build());

    }
}
