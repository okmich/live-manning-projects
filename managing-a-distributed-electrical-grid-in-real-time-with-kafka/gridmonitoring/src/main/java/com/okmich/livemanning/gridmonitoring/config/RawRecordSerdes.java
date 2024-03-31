package com.okmich.livemanning.gridmonitoring.config;

import manning.devices.raw.RawRecord;
import org.apache.kafka.common.serialization.Deserializer;
import org.apache.kafka.common.serialization.Serdes;
import org.apache.kafka.common.serialization.Serializer;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.Map;

public class RawRecordSerdes extends Serdes.WrapperSerde<RawRecord> {

    public RawRecordSerdes() {
        super(new RawRecordSerializer(), new RawRecordDeserializer());
    }

    public static class RawRecordSerializer implements Serializer<RawRecord> {

        @Override
        public void configure(Map<String, ?> configs, boolean isKey) {
            Serializer.super.configure(configs, isKey);
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

    public static class RawRecordDeserializer implements Deserializer<RawRecord> {

        @Override
        public void configure(Map<String, ?> configs, boolean isKey) {
            Deserializer.super.configure(configs, isKey);
        }

        @Override
        public RawRecord deserialize(String topic, byte[] bytes) {
            if (bytes == null) return null;
            else {
                try {
                    return RawRecord.fromByteBuffer(ByteBuffer.wrap(bytes));
                } catch (IOException e) {
                    throw new RuntimeException(e);
                }
            }
        }
    }
}
