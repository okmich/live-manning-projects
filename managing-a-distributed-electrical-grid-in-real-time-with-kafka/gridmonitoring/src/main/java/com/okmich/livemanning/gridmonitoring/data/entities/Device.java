package com.okmich.livemanning.gridmonitoring.data.entities;

import io.hypersistence.utils.hibernate.type.json.JsonType;
import jakarta.persistence.*;
import org.hibernate.annotations.Type;

import java.io.Serializable;
import java.time.LocalDateTime;
import java.util.Map;
import java.util.Objects;

/**
 * @author okmic
 */
@Entity
@Table(name = "device_reading")
public class Device implements Serializable {

    private static final long serialVersionUID = 12468L;
    @Id
    @Column(name = "id", nullable = false)
    private String deviceId;
    @Type(JsonType.class)
    @Column(name = "metadata", columnDefinition = "jsonb")
    private Map<String, Object> readings;
    @Column(name = "arrival_millis", nullable = false)
    private Long arrivalTsMillis;
    @Column(name = "event_millis", nullable = false)
    private Long eventTsMillis;
    @Column(name = "ts", nullable = false)
    @Temporal(TemporalType.TIMESTAMP)
    private LocalDateTime ts;


    public Device() {
    }

    public Device(String deviceId, Map<String, Object> readings, long arriveTimeInMillis, long eventTimeInMillis) {
        this.deviceId = deviceId;
        this.readings = readings;
        this.arrivalTsMillis = arriveTimeInMillis;
        this.eventTsMillis = eventTimeInMillis;
        this.ts = LocalDateTime.now();
    }

    public String getDeviceId() {
        return deviceId;
    }

    public void setDeviceId(String deviceId) {
        this.deviceId = deviceId;
    }

    public Map<String, Object> getReadings() {
        return readings;
    }

    public void setReadings(Map<String, Object> readings) {
        this.readings = readings;
    }

    public Long getArrivalTsMillis() {
        return arrivalTsMillis;
    }

    public void setArrivalTsMillis(Long arrivalTsMillis) {
        this.arrivalTsMillis = arrivalTsMillis;
    }

    public Long getEventTsMillis() {
        return eventTsMillis;
    }

    public void setEventTsMillis(Long eventTsMillis) {
        this.eventTsMillis = eventTsMillis;
    }

    public LocalDateTime getTs() {
        return ts;
    }

    public void setTs(LocalDateTime ts) {
        this.ts = ts;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Device that = (Device) o;
        return Objects.equals(deviceId, that.deviceId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(deviceId);
    }
}
