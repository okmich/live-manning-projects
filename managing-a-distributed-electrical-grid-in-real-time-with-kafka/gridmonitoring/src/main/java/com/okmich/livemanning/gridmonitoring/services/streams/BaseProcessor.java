package com.okmich.livemanning.gridmonitoring.services.streams;

import org.apache.kafka.streams.StreamsBuilder;
import org.apache.kafka.streams.kstream.KStream;

public abstract class BaseProcessor<K, V> {

    private final String sourceTopic;
    private final String targetTopic;
    private final StreamsBuilder streamBuilder;

    public BaseProcessor(String sourceTopic, String targetTopic, StreamsBuilder streamBuilder) {
        this.sourceTopic = sourceTopic;
        this.targetTopic = targetTopic;
        this.streamBuilder = streamBuilder;
    }

    public abstract KStream<K, V> getKStream();

    public String getSourceTopic() {
        return sourceTopic;
    }

    public String getTargetTopic() {
        return targetTopic;
    }

    public StreamsBuilder getStreamBuilder() {
        return streamBuilder;
    }
}
