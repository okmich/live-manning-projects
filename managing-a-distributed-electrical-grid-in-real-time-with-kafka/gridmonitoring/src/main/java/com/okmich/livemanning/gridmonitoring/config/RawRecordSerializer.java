package com.okmich.livemanning.gridmonitoring.config;

import manning.devices.raw.m1.RawRecord;
import org.apache.kafka.common.serialization.Serializer;

import java.io.IOException;
import java.util.Map;

public class RawRecordSerializer implements Serializer<RawRecord> {

    @Override
    public void configure(Map<String, ?> configs, boolean isKey) {
        Serializer.super.configure(configs, isKey);
        try {
            Class.forName("manning.devices.raw.m1.RawRecord");
        } catch (ClassNotFoundException e) {
            throw new RuntimeException(e);
        }
    }

    @Override
    public byte[] serialize(String topic, RawRecord rawRecord) {
        if (rawRecord == null || rawRecord.getBody() == null)
            return null;

        try {
            return rawRecord.toByteBuffer().array();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
