package com.okmich.livemanning.gridmonitoring.services.streams;


import com.okmich.livemanning.gridmonitoring.config.CanonicalKeySerdes;
import com.okmich.livemanning.gridmonitoring.config.CanonicalValueSerdes;
import com.okmich.livemanning.gridmonitoring.data.entities.Device;
import com.okmich.livemanning.gridmonitoring.data.repo.DeviceRepository;
import io.confluent.kafka.streams.serdes.avro.SpecificAvroSerde;
import manning.devices.canonical.CanonicalKey;
import manning.devices.canonical.CanonicalValue;
import org.apache.kafka.common.serialization.Serde;
import org.apache.kafka.streams.StreamsBuilder;
import org.apache.kafka.streams.kstream.Consumed;
import org.apache.kafka.streams.kstream.KStream;

import java.util.HashMap;
import java.util.Map;

public class CanonicalToDbProcessor extends BaseProcessor<CanonicalKey, CanonicalValue> {

    private final DeviceRepository deviceRepository;

    public CanonicalToDbProcessor(String sourceTopic, StreamsBuilder streamBuilder, DeviceRepository deviceRepository) {
        super(sourceTopic, "", streamBuilder);
        this.deviceRepository = deviceRepository;
    }

    public KStream<CanonicalKey, CanonicalValue> getKStream() {
        KStream<CanonicalKey, CanonicalValue> stream = getStreamBuilder().
                stream(this.getSourceTopic(), Consumed.with(new CanonicalKeySerdes(), new CanonicalValueSerdes()));
        stream.foreach((canonicalKey, canonicalValue) -> {
            if (canonicalValue.getEvents().containsKey("charging")) {
                Map<String, Object> events = new HashMap<>();
                for (var entry : canonicalValue.getEvents().entrySet())
                    events.put(String.valueOf(entry.getKey()), entry.getValue());

                var device = new Device(
                        canonicalValue.getUuid().toString(),
                        events,
                        canonicalValue.getArrivalTimeMs(),
                        canonicalValue.getEventTimeMs()
                );
                deviceRepository.save(device);
            }
        });
        return stream;
    }
}
