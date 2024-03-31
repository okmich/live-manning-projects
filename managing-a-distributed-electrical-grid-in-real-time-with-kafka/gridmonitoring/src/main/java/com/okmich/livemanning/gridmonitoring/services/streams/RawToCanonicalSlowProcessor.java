package com.okmich.livemanning.gridmonitoring.services.streams;


import com.okmich.livemanning.gridmonitoring.config.CanonicalKeySerdes;
import com.okmich.livemanning.gridmonitoring.config.CanonicalValueSerdes;
import com.okmich.livemanning.gridmonitoring.config.RawRecordSerdes;
import manning.devices.canonical.CanonicalKey;
import manning.devices.raw.RawRecord;
import org.apache.avro.specific.SpecificRecord;
import org.apache.kafka.common.serialization.Serdes;
import org.apache.kafka.streams.StreamsBuilder;
import org.apache.kafka.streams.kstream.Consumed;
import org.apache.kafka.streams.kstream.KStream;
import org.apache.kafka.streams.kstream.Produced;

public class RawToCanonicalSlowProcessor extends BaseProcessor<String, RawRecord> {


    public RawToCanonicalSlowProcessor(String sourceTopic, String targetTopic,
                                       StreamsBuilder streamBuilder) {
        super(sourceTopic, targetTopic, streamBuilder);
    }

    public KStream<String, RawRecord> getKStream() {
        KStream<String, RawRecord> rawStream = getStreamBuilder().stream(getSourceTopic(),
                Consumed.with(Serdes.String(), new RawRecordSerdes()));
        rawStream
                .flatMap(new RawToCanonicalProcessor.RawToCanonicalMapper())
                .to(getTargetTopic());
        return rawStream;
    }
}
