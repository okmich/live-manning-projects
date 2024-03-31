package com.okmich.livemanning.gridmonitoring.config;


import org.apache.kafka.clients.admin.AdminClientConfig;
import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.core.KafkaAdmin;

import java.util.HashMap;
import java.util.Map;

@Configuration
public class KafkaTopicConfig {
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

    @Bean
    public KafkaAdmin kafkaAdmin() {
        Map<String, Object> configs = new HashMap<>();
        configs.put(AdminClientConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapAddress);
        return new KafkaAdmin(configs);
    }

    @Bean
    public NewTopic rawTopic() {
        return new NewTopic(rawTopic, 1, (short) 1);
    }

    @Bean
    public NewTopic rawSlowTopic() {
        return new NewTopic(rawSlowTopic, 1, (short) 1);
    }

    @Bean
    public NewTopic rawErrorTopic() {
        return new NewTopic(rawDlqTopic, 1, (short) 1);
    }

    @Bean
    public NewTopic canonicalTopic() {
        return new NewTopic(canonicalTopic, 1, (short) 1);
    }
}
