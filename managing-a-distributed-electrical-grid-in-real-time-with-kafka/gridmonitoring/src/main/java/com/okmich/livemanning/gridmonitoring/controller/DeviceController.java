package com.okmich.livemanning.gridmonitoring.controller;


import com.okmich.livemanning.gridmonitoring.data.repo.DeviceRepository;
import jakarta.servlet.http.HttpServletRequest;
import manning.devices.raw.RawRecord;
import org.apache.kafka.clients.producer.RecordMetadata;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.StandardOpenOption;
import java.time.Instant;
import java.util.Map;
import java.util.concurrent.ExecutionException;

import static org.apache.commons.compress.utils.IOUtils.toByteArray;

@RestController
@RequestMapping(value = "/device", consumes = {MediaType.APPLICATION_OCTET_STREAM_VALUE, MediaType.APPLICATION_JSON_VALUE}, produces = MediaType.APPLICATION_JSON_VALUE)
public class DeviceController {

    private final KafkaTemplate<String, RawRecord> kafkaTemplate;
    private final DeviceRepository deviceRepository;
    @Value("${kafka.event.raw.topic}")
    private String rawTopic;
    @Value("${kafka.event.raw.slow.topic}")
    private String rawSlowTopic;
    @Value("${device.max-size}")
    private int maxSize;

    @Autowired
    public DeviceController(KafkaTemplate<String, RawRecord> kafkaTemplate, DeviceRepository deviceRepository) {
        this.kafkaTemplate = kafkaTemplate;
        this.deviceRepository = deviceRepository;
    }

    @PostMapping("/send/{uuid}")
    public ResponseEntity<Map<String, Object>> send(@PathVariable("uuid") String uuid, HttpServletRequest request)
            throws IOException, ExecutionException, InterruptedException {
        ByteBuffer byteBuffer = ByteBuffer.wrap(toByteArray(request.getInputStream()));
        var now = Instant.now().toEpochMilli();
        RawRecord rawRecord = null;
        if (maxSize >= byteBuffer.capacity())
            rawRecord = new RawRecord(uuid, now, byteBuffer, null);
        else {
            try {
                String fileName = String.format("%s-%d", uuid, now);
                var path = Paths.get(System.getProperty("user.dir"), "temp_files", fileName);
                Files.write(path, byteBuffer.array(), StandardOpenOption.TRUNCATE_EXISTING);
                rawRecord = new RawRecord(uuid, now, null, path.toFile().getAbsolutePath());
            } catch (Exception e) {
                e.printStackTrace();
                return ResponseEntity.internalServerError().body(Map.of("message", e.getMessage()));
            }
        }

        var resultSendFuture = kafkaTemplate.send(
                rawRecord.getBody() != null ? this.rawTopic : this.rawSlowTopic,
                uuid,
                rawRecord
        );
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
