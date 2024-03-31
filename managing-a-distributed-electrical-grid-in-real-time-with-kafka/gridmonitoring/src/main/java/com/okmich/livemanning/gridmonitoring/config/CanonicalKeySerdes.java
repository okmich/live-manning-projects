package com.okmich.livemanning.gridmonitoring.config;

import manning.devices.canonical.CanonicalKey;
import org.apache.kafka.common.serialization.Deserializer;
import org.apache.kafka.common.serialization.Serdes;
import org.apache.kafka.common.serialization.Serializer;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.Map;

public class CanonicalKeySerdes extends Serdes.WrapperSerde<CanonicalKey> {

    public CanonicalKeySerdes() {
        super(new CanonicalKeySerializer(), new CanonicalKeyDeserializer());
    }

    public static class CanonicalKeySerializer implements Serializer<CanonicalKey> {

        @Override
        public void configure(Map<String, ?> configs, boolean isKey) {
            Serializer.super.configure(configs, isKey);
        }

        @Override
        public byte[] serialize(String topic, CanonicalKey canonicalKey) {
            if (canonicalKey == null)
                return null;

            try {
                return canonicalKey.toByteBuffer().array();
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        }
    }


    public static class CanonicalKeyDeserializer implements Deserializer<CanonicalKey> {

        @Override
        public void configure(Map<String, ?> configs, boolean isKey) {
            Deserializer.super.configure(configs, isKey);
        }

        @Override
        public CanonicalKey deserialize(String topic, byte[] bytes) {
            if (bytes == null) return null;
            else {
                try {
                    return CanonicalKey.fromByteBuffer(ByteBuffer.wrap(bytes));
                } catch (IOException e) {
                    throw new RuntimeException(e);
                }
            }
        }
    }
}
