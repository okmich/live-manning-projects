package com.okmich.livemanning.gridmonitoring.services.streams;


import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.okmich.livemanning.gridmonitoring.config.RawRecordSerdes;
import io.confluent.kafka.streams.serdes.avro.SpecificAvroSerde;
import manning.devices.canonical.CanonicalErrorValue;
import manning.devices.canonical.CanonicalKey;
import manning.devices.canonical.CanonicalValue;
import manning.devices.raw.RawRecord;
import org.apache.avro.specific.SpecificRecord;
import org.apache.kafka.common.serialization.Serde;
import org.apache.kafka.common.serialization.Serdes;
import org.apache.kafka.streams.KeyValue;
import org.apache.kafka.streams.StreamsBuilder;
import org.apache.kafka.streams.kstream.Consumed;
import org.apache.kafka.streams.kstream.KStream;
import org.apache.kafka.streams.kstream.KeyValueMapper;
import org.apache.kafka.streams.kstream.Produced;

import java.io.*;
import java.time.Instant;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.function.Function;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class RawToCanonicalProcessor extends BaseProcessor<String, RawRecord> {

    private final String dlqTopic;

    public RawToCanonicalProcessor(String sourceTopic, String targetTopic, String dlqTopic,
                                   StreamsBuilder streamBuilder) {
        super(sourceTopic, targetTopic, streamBuilder);
        this.dlqTopic = dlqTopic;
    }

    public KStream<String, RawRecord> getKStream() {
        KStream<String, RawRecord> rawStream = getStreamBuilder().stream(getSourceTopic(),
                Consumed.with(Serdes.String(), new RawRecordSerdes()));
        rawStream
                .flatMap(new RawToCanonicalMapper())
                .to((canonicalKey, specificRecord, recordContext) -> {
                    if (specificRecord instanceof CanonicalValue)
                        return getTargetTopic();
                    else
                        return dlqTopic;
                });
        return rawStream;
    }

    public static class RawToCanonicalMapper implements
            KeyValueMapper<String, RawRecord, Iterable<? extends KeyValue<CanonicalKey, SpecificRecord>>> {

        static final ObjectMapper MAPPER = new ObjectMapper();
        static final TypeReference<Map<String, String>> MAP_TYPE_REFERENCE =
                new TypeReference<Map<String, String>>() {
                };
        public RawToCanonicalMapper() {
        }

        @Override
        public Iterable<? extends KeyValue<CanonicalKey, SpecificRecord>> apply(String key, RawRecord rawRecord) {
            List<Map<String, String>> records = new ArrayList<>();
            List<KeyValue<CanonicalKey, SpecificRecord>> badRecords = new ArrayList<>();
            try (BufferedReader bufferedReader = getBufferedReaderForRecord(rawRecord)) {
                String line;
                while ((line = bufferedReader.readLine()) != null)
                    try{
                        records.add(MAPPER.readValue(line, MAP_TYPE_REFERENCE));
                    } catch (JsonProcessingException e) {
                        badRecords.add(new KeyValue<>(new CanonicalKey(rawRecord.getUuid()),
                                new CanonicalErrorValue(List.of(e.getMessage()), rawRecord.getBody())));
                    }
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
            var goodRecordStream = records.stream()
                    .map(getMapKeyValueFunction(key, rawRecord));

            return Stream.concat(goodRecordStream, badRecords.stream())
                    .filter(Objects::nonNull)
                    .collect(Collectors.toList());
        }

        public Function<Map<String, String>, KeyValue<CanonicalKey, SpecificRecord>> getMapKeyValueFunction(String key, RawRecord rawRecord) {
            return record -> {
                try {
                    Map<CharSequence, CharSequence> events = Map.copyOf(record);
                    var canonicalValue = new CanonicalValue(
                            rawRecord.getUuid(),
                            record.get("region") == null ? null : Long.valueOf(record.get("region")),
                            rawRecord.getArrivalTimeMs(),
                            Instant.now().toEpochMilli(),
                            events
                    );
                    return new KeyValue<>(new CanonicalKey(rawRecord.getUuid()), canonicalValue);
                } catch (Exception e) {
                    return new KeyValue<>(
                            new CanonicalKey(rawRecord.getUuid()),
                            new CanonicalErrorValue(
                                    List.of(e.getMessage()),
                                    rawRecord.getBody()
                            )
                    );
                }
            };
        }

        public BufferedReader getBufferedReaderForRecord(RawRecord rawRecord) throws FileNotFoundException {
            if (rawRecord.getBody() != null) {
                ByteArrayInputStream inputStream = new ByteArrayInputStream(rawRecord.getBody().array());
                return new BufferedReader(new InputStreamReader(inputStream));
            } else { //uses body_reference
                return new BufferedReader(new InputStreamReader(
                        new FileInputStream(String.valueOf(rawRecord.getBodyReference()))));
            }
        }
    }
}
