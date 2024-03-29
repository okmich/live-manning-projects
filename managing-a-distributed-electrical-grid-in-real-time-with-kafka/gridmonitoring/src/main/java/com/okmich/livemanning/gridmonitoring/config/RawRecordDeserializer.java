package com.okmich.livemanning.gridmonitoring.config;

import manning.devices.raw.m1.RawRecord;
import org.apache.kafka.common.header.Headers;
import org.apache.kafka.common.serialization.Deserializer;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.Map;

public class RawRecordDeserializer implements Deserializer<RawRecord> {

    @Override
    public void configure(Map<String, ?> configs, boolean isKey) {
        Deserializer.super.configure(configs, isKey);
        try {
            Class.forName("manning.devices.raw.m1.RawRecord");
        } catch (ClassNotFoundException e) {
            throw new RuntimeException(e);
        }
    }

    @Override
    public RawRecord deserialize(String topic, byte[] bytes) {
        ByteBuffer byteBuffer = ByteBuffer.wrap(bytes);
        return this.deserialize(topic, null, byteBuffer);
    }

    @Override
    public RawRecord deserialize(String topic, Headers headers, byte[] bytes) {
        return this.deserialize(topic, headers, ByteBuffer.wrap(bytes));
    }

    @Override
    public RawRecord deserialize(String topic, Headers headers, ByteBuffer bytes) {
        try {
            return RawRecord.fromByteBuffer(bytes);
        } catch (IOException e) {
            return null;
        }
    }
}
