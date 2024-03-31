package com.okmich.livemanning.gridmonitoring.config;


import com.okmich.livemanning.gridmonitoring.data.repo.DeviceRepository;
import com.okmich.livemanning.gridmonitoring.services.streams.CanonicalToDbProcessor;
import com.okmich.livemanning.gridmonitoring.services.streams.RawToCanonicalProcessor;
import com.okmich.livemanning.gridmonitoring.services.streams.RawToCanonicalSlowProcessor;
import manning.devices.canonical.CanonicalKey;
import manning.devices.canonical.CanonicalValue;
import manning.devices.raw.RawRecord;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.common.serialization.Serdes;
import org.apache.kafka.streams.StreamsBuilder;
import org.apache.kafka.streams.kstream.KStream;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.annotation.EnableKafkaStreams;
import org.springframework.kafka.annotation.KafkaStreamsDefaultConfiguration;
import org.springframework.kafka.config.KafkaStreamsConfiguration;

import java.util.HashMap;
import java.util.Map;

import static org.apache.kafka.streams.StreamsConfig.*;

@EnableKafkaStreams
@Configuration
public class KafkaStreamsConfig {

    @Value("${spring.kafka.bootstrap-servers}")
    private String bootstrapAddress;
    @Value("${kafka.event.raw.topic}")
    private String rawTopic;
    @Value("${kafka.event.raw.slow.topic}")
    private String rawSlowTopic;
    @Value("${kafka.event.raw_dlq.topic}")
    private String rawDlqTopic;
    @Value("${kafka.event.canonical.topic}")
    private String canonicalTopic;

    @Bean(name = KafkaStreamsDefaultConfiguration.DEFAULT_STREAMS_CONFIG_BEAN_NAME)
    public KafkaStreamsConfiguration kStreamsConfig() {
        Map<String, Object> props = new HashMap<>();
        props.put(APPLICATION_ID_CONFIG, "device-streams-app");
        props.put(BOOTSTRAP_SERVERS_CONFIG, bootstrapAddress);
        props.put(DEFAULT_KEY_SERDE_CLASS_CONFIG, CanonicalKeySerdes.class.getName());
        props.put(DEFAULT_VALUE_SERDE_CLASS_CONFIG, CanonicalValueSerdes.class.getName());
        props.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "latest");

        return new KafkaStreamsConfiguration(props);
    }

    @Bean
    public KStream<String, RawRecord> rawToCanonicalKStream(StreamsBuilder kStreamBuilder) {
        return new RawToCanonicalProcessor(
                rawTopic, canonicalTopic,
                rawDlqTopic,
                kStreamBuilder
        ).getKStream();
    }

    @Bean
    public KStream<String, RawRecord> rawSlowToCanonicalKStream(StreamsBuilder kStreamBuilder) {
        return new RawToCanonicalSlowProcessor(
                rawSlowTopic, canonicalTopic,
                kStreamBuilder
        ).getKStream();
    }

    @Bean
    public KStream<CanonicalKey, CanonicalValue> canonicalToDatabaseKStream(StreamsBuilder kStreamBuilder,
                                                                            DeviceRepository deviceRepository) {
        return new CanonicalToDbProcessor(canonicalTopic, kStreamBuilder, deviceRepository).getKStream();
    }
}
