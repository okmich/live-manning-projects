package com.okmich.livemanning.gridmonitoring.config;

import manning.devices.canonical.CanonicalValue;
import org.apache.kafka.common.serialization.Deserializer;
import org.apache.kafka.common.serialization.Serdes;
import org.apache.kafka.common.serialization.Serializer;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.Map;

public class CanonicalValueSerdes extends Serdes.WrapperSerde<CanonicalValue> {

    public CanonicalValueSerdes() {
        super(new CanonicalValueSerializer(), new CanonicalValueDeserializer());
    }

    public static class CanonicalValueSerializer implements Serializer<CanonicalValue> {

        @Override
        public void configure(Map<String, ?> configs, boolean isKey) {
            Serializer.super.configure(configs, isKey);
        }

        @Override
        public byte[] serialize(String topic, CanonicalValue CanonicalValue) {
            if (CanonicalValue == null)
                return null;

            try {
                return CanonicalValue.toByteBuffer().array();
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        }
    }


    public static class CanonicalValueDeserializer implements Deserializer<CanonicalValue> {

        @Override
        public void configure(Map<String, ?> configs, boolean isKey) {
            Deserializer.super.configure(configs, isKey);
        }

        @Override
        public CanonicalValue deserialize(String topic, byte[] bytes) {
            if (bytes == null) return null;
            else {
                try {
                    return CanonicalValue.fromByteBuffer(ByteBuffer.wrap(bytes));
                } catch (IOException e) {
                    throw new RuntimeException(e);
                }
            }
        }
    }
}
