/**
 * Autogenerated by Avro
 *
 * DO NOT EDIT DIRECTLY
 */
package manning.devices.canonical;

import org.apache.avro.generic.GenericArray;
import org.apache.avro.specific.SpecificData;
import org.apache.avro.util.Utf8;
import org.apache.avro.message.BinaryMessageEncoder;
import org.apache.avro.message.BinaryMessageDecoder;
import org.apache.avro.message.SchemaStore;

@org.apache.avro.specific.AvroGenerated
public class CanonicalValue extends org.apache.avro.specific.SpecificRecordBase implements org.apache.avro.specific.SpecificRecord {
  private static final long serialVersionUID = 5597930688600105100L;


  public static final org.apache.avro.Schema SCHEMA$ = new org.apache.avro.Schema.Parser().parse("{\"type\":\"record\",\"name\":\"CanonicalValue\",\"namespace\":\"manning.devices.canonical\",\"fields\":[{\"name\":\"uuid\",\"type\":\"string\"},{\"name\":\"regionId\",\"type\":[\"null\",\"long\"]},{\"name\":\"arrival_time_ms\",\"type\":[\"null\",\"long\"]},{\"name\":\"event_time_ms\",\"type\":[\"null\",\"long\"]},{\"name\":\"events\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]}]}");
  public static org.apache.avro.Schema getClassSchema() { return SCHEMA$; }

  private static final SpecificData MODEL$ = new SpecificData();

  private static final BinaryMessageEncoder<CanonicalValue> ENCODER =
      new BinaryMessageEncoder<>(MODEL$, SCHEMA$);

  private static final BinaryMessageDecoder<CanonicalValue> DECODER =
      new BinaryMessageDecoder<>(MODEL$, SCHEMA$);

  /**
   * Return the BinaryMessageEncoder instance used by this class.
   * @return the message encoder used by this class
   */
  public static BinaryMessageEncoder<CanonicalValue> getEncoder() {
    return ENCODER;
  }

  /**
   * Return the BinaryMessageDecoder instance used by this class.
   * @return the message decoder used by this class
   */
  public static BinaryMessageDecoder<CanonicalValue> getDecoder() {
    return DECODER;
  }

  /**
   * Create a new BinaryMessageDecoder instance for this class that uses the specified {@link SchemaStore}.
   * @param resolver a {@link SchemaStore} used to find schemas by fingerprint
   * @return a BinaryMessageDecoder instance for this class backed by the given SchemaStore
   */
  public static BinaryMessageDecoder<CanonicalValue> createDecoder(SchemaStore resolver) {
    return new BinaryMessageDecoder<>(MODEL$, SCHEMA$, resolver);
  }

  /**
   * Serializes this CanonicalValue to a ByteBuffer.
   * @return a buffer holding the serialized data for this instance
   * @throws java.io.IOException if this instance could not be serialized
   */
  public java.nio.ByteBuffer toByteBuffer() throws java.io.IOException {
    return ENCODER.encode(this);
  }

  /**
   * Deserializes a CanonicalValue from a ByteBuffer.
   * @param b a byte buffer holding serialized data for an instance of this class
   * @return a CanonicalValue instance decoded from the given buffer
   * @throws java.io.IOException if the given bytes could not be deserialized into an instance of this class
   */
  public static CanonicalValue fromByteBuffer(
      java.nio.ByteBuffer b) throws java.io.IOException {
    return DECODER.decode(b);
  }

  private java.lang.CharSequence uuid;
  private java.lang.Long regionId;
  private java.lang.Long arrival_time_ms;
  private java.lang.Long event_time_ms;
  private java.util.Map<java.lang.CharSequence,java.lang.CharSequence> events;

  /**
   * Default constructor.  Note that this does not initialize fields
   * to their default values from the schema.  If that is desired then
   * one should use <code>newBuilder()</code>.
   */
  public CanonicalValue() {}

  /**
   * All-args constructor.
   * @param uuid The new value for uuid
   * @param regionId The new value for regionId
   * @param arrival_time_ms The new value for arrival_time_ms
   * @param event_time_ms The new value for event_time_ms
   * @param events The new value for events
   */
  public CanonicalValue(java.lang.CharSequence uuid, java.lang.Long regionId, java.lang.Long arrival_time_ms, java.lang.Long event_time_ms, java.util.Map<java.lang.CharSequence,java.lang.CharSequence> events) {
    this.uuid = uuid;
    this.regionId = regionId;
    this.arrival_time_ms = arrival_time_ms;
    this.event_time_ms = event_time_ms;
    this.events = events;
  }

  @Override
  public org.apache.avro.specific.SpecificData getSpecificData() { return MODEL$; }

  @Override
  public org.apache.avro.Schema getSchema() { return SCHEMA$; }

  // Used by DatumWriter.  Applications should not call.
  @Override
  public java.lang.Object get(int field$) {
    switch (field$) {
    case 0: return uuid;
    case 1: return regionId;
    case 2: return arrival_time_ms;
    case 3: return event_time_ms;
    case 4: return events;
    default: throw new IndexOutOfBoundsException("Invalid index: " + field$);
    }
  }

  // Used by DatumReader.  Applications should not call.
  @Override
  @SuppressWarnings(value="unchecked")
  public void put(int field$, java.lang.Object value$) {
    switch (field$) {
    case 0: uuid = (java.lang.CharSequence)value$; break;
    case 1: regionId = (java.lang.Long)value$; break;
    case 2: arrival_time_ms = (java.lang.Long)value$; break;
    case 3: event_time_ms = (java.lang.Long)value$; break;
    case 4: events = (java.util.Map<java.lang.CharSequence,java.lang.CharSequence>)value$; break;
    default: throw new IndexOutOfBoundsException("Invalid index: " + field$);
    }
  }

  /**
   * Gets the value of the 'uuid' field.
   * @return The value of the 'uuid' field.
   */
  public java.lang.CharSequence getUuid() {
    return uuid;
  }


  /**
   * Sets the value of the 'uuid' field.
   * @param value the value to set.
   */
  public void setUuid(java.lang.CharSequence value) {
    this.uuid = value;
  }

  /**
   * Gets the value of the 'regionId' field.
   * @return The value of the 'regionId' field.
   */
  public java.lang.Long getRegionId() {
    return regionId;
  }


  /**
   * Sets the value of the 'regionId' field.
   * @param value the value to set.
   */
  public void setRegionId(java.lang.Long value) {
    this.regionId = value;
  }

  /**
   * Gets the value of the 'arrival_time_ms' field.
   * @return The value of the 'arrival_time_ms' field.
   */
  public java.lang.Long getArrivalTimeMs() {
    return arrival_time_ms;
  }


  /**
   * Sets the value of the 'arrival_time_ms' field.
   * @param value the value to set.
   */
  public void setArrivalTimeMs(java.lang.Long value) {
    this.arrival_time_ms = value;
  }

  /**
   * Gets the value of the 'event_time_ms' field.
   * @return The value of the 'event_time_ms' field.
   */
  public java.lang.Long getEventTimeMs() {
    return event_time_ms;
  }


  /**
   * Sets the value of the 'event_time_ms' field.
   * @param value the value to set.
   */
  public void setEventTimeMs(java.lang.Long value) {
    this.event_time_ms = value;
  }

  /**
   * Gets the value of the 'events' field.
   * @return The value of the 'events' field.
   */
  public java.util.Map<java.lang.CharSequence,java.lang.CharSequence> getEvents() {
    return events;
  }


  /**
   * Sets the value of the 'events' field.
   * @param value the value to set.
   */
  public void setEvents(java.util.Map<java.lang.CharSequence,java.lang.CharSequence> value) {
    this.events = value;
  }

  /**
   * Creates a new CanonicalValue RecordBuilder.
   * @return A new CanonicalValue RecordBuilder
   */
  public static manning.devices.canonical.CanonicalValue.Builder newBuilder() {
    return new manning.devices.canonical.CanonicalValue.Builder();
  }

  /**
   * Creates a new CanonicalValue RecordBuilder by copying an existing Builder.
   * @param other The existing builder to copy.
   * @return A new CanonicalValue RecordBuilder
   */
  public static manning.devices.canonical.CanonicalValue.Builder newBuilder(manning.devices.canonical.CanonicalValue.Builder other) {
    if (other == null) {
      return new manning.devices.canonical.CanonicalValue.Builder();
    } else {
      return new manning.devices.canonical.CanonicalValue.Builder(other);
    }
  }

  /**
   * Creates a new CanonicalValue RecordBuilder by copying an existing CanonicalValue instance.
   * @param other The existing instance to copy.
   * @return A new CanonicalValue RecordBuilder
   */
  public static manning.devices.canonical.CanonicalValue.Builder newBuilder(manning.devices.canonical.CanonicalValue other) {
    if (other == null) {
      return new manning.devices.canonical.CanonicalValue.Builder();
    } else {
      return new manning.devices.canonical.CanonicalValue.Builder(other);
    }
  }

  /**
   * RecordBuilder for CanonicalValue instances.
   */
  @org.apache.avro.specific.AvroGenerated
  public static class Builder extends org.apache.avro.specific.SpecificRecordBuilderBase<CanonicalValue>
    implements org.apache.avro.data.RecordBuilder<CanonicalValue> {

    private java.lang.CharSequence uuid;
    private java.lang.Long regionId;
    private java.lang.Long arrival_time_ms;
    private java.lang.Long event_time_ms;
    private java.util.Map<java.lang.CharSequence,java.lang.CharSequence> events;

    /** Creates a new Builder */
    private Builder() {
      super(SCHEMA$, MODEL$);
    }

    /**
     * Creates a Builder by copying an existing Builder.
     * @param other The existing Builder to copy.
     */
    private Builder(manning.devices.canonical.CanonicalValue.Builder other) {
      super(other);
      if (isValidValue(fields()[0], other.uuid)) {
        this.uuid = data().deepCopy(fields()[0].schema(), other.uuid);
        fieldSetFlags()[0] = other.fieldSetFlags()[0];
      }
      if (isValidValue(fields()[1], other.regionId)) {
        this.regionId = data().deepCopy(fields()[1].schema(), other.regionId);
        fieldSetFlags()[1] = other.fieldSetFlags()[1];
      }
      if (isValidValue(fields()[2], other.arrival_time_ms)) {
        this.arrival_time_ms = data().deepCopy(fields()[2].schema(), other.arrival_time_ms);
        fieldSetFlags()[2] = other.fieldSetFlags()[2];
      }
      if (isValidValue(fields()[3], other.event_time_ms)) {
        this.event_time_ms = data().deepCopy(fields()[3].schema(), other.event_time_ms);
        fieldSetFlags()[3] = other.fieldSetFlags()[3];
      }
      if (isValidValue(fields()[4], other.events)) {
        this.events = data().deepCopy(fields()[4].schema(), other.events);
        fieldSetFlags()[4] = other.fieldSetFlags()[4];
      }
    }

    /**
     * Creates a Builder by copying an existing CanonicalValue instance
     * @param other The existing instance to copy.
     */
    private Builder(manning.devices.canonical.CanonicalValue other) {
      super(SCHEMA$, MODEL$);
      if (isValidValue(fields()[0], other.uuid)) {
        this.uuid = data().deepCopy(fields()[0].schema(), other.uuid);
        fieldSetFlags()[0] = true;
      }
      if (isValidValue(fields()[1], other.regionId)) {
        this.regionId = data().deepCopy(fields()[1].schema(), other.regionId);
        fieldSetFlags()[1] = true;
      }
      if (isValidValue(fields()[2], other.arrival_time_ms)) {
        this.arrival_time_ms = data().deepCopy(fields()[2].schema(), other.arrival_time_ms);
        fieldSetFlags()[2] = true;
      }
      if (isValidValue(fields()[3], other.event_time_ms)) {
        this.event_time_ms = data().deepCopy(fields()[3].schema(), other.event_time_ms);
        fieldSetFlags()[3] = true;
      }
      if (isValidValue(fields()[4], other.events)) {
        this.events = data().deepCopy(fields()[4].schema(), other.events);
        fieldSetFlags()[4] = true;
      }
    }

    /**
      * Gets the value of the 'uuid' field.
      * @return The value.
      */
    public java.lang.CharSequence getUuid() {
      return uuid;
    }


    /**
      * Sets the value of the 'uuid' field.
      * @param value The value of 'uuid'.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder setUuid(java.lang.CharSequence value) {
      validate(fields()[0], value);
      this.uuid = value;
      fieldSetFlags()[0] = true;
      return this;
    }

    /**
      * Checks whether the 'uuid' field has been set.
      * @return True if the 'uuid' field has been set, false otherwise.
      */
    public boolean hasUuid() {
      return fieldSetFlags()[0];
    }


    /**
      * Clears the value of the 'uuid' field.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder clearUuid() {
      uuid = null;
      fieldSetFlags()[0] = false;
      return this;
    }

    /**
      * Gets the value of the 'regionId' field.
      * @return The value.
      */
    public java.lang.Long getRegionId() {
      return regionId;
    }


    /**
      * Sets the value of the 'regionId' field.
      * @param value The value of 'regionId'.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder setRegionId(java.lang.Long value) {
      validate(fields()[1], value);
      this.regionId = value;
      fieldSetFlags()[1] = true;
      return this;
    }

    /**
      * Checks whether the 'regionId' field has been set.
      * @return True if the 'regionId' field has been set, false otherwise.
      */
    public boolean hasRegionId() {
      return fieldSetFlags()[1];
    }


    /**
      * Clears the value of the 'regionId' field.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder clearRegionId() {
      regionId = null;
      fieldSetFlags()[1] = false;
      return this;
    }

    /**
      * Gets the value of the 'arrival_time_ms' field.
      * @return The value.
      */
    public java.lang.Long getArrivalTimeMs() {
      return arrival_time_ms;
    }


    /**
      * Sets the value of the 'arrival_time_ms' field.
      * @param value The value of 'arrival_time_ms'.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder setArrivalTimeMs(java.lang.Long value) {
      validate(fields()[2], value);
      this.arrival_time_ms = value;
      fieldSetFlags()[2] = true;
      return this;
    }

    /**
      * Checks whether the 'arrival_time_ms' field has been set.
      * @return True if the 'arrival_time_ms' field has been set, false otherwise.
      */
    public boolean hasArrivalTimeMs() {
      return fieldSetFlags()[2];
    }


    /**
      * Clears the value of the 'arrival_time_ms' field.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder clearArrivalTimeMs() {
      arrival_time_ms = null;
      fieldSetFlags()[2] = false;
      return this;
    }

    /**
      * Gets the value of the 'event_time_ms' field.
      * @return The value.
      */
    public java.lang.Long getEventTimeMs() {
      return event_time_ms;
    }


    /**
      * Sets the value of the 'event_time_ms' field.
      * @param value The value of 'event_time_ms'.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder setEventTimeMs(java.lang.Long value) {
      validate(fields()[3], value);
      this.event_time_ms = value;
      fieldSetFlags()[3] = true;
      return this;
    }

    /**
      * Checks whether the 'event_time_ms' field has been set.
      * @return True if the 'event_time_ms' field has been set, false otherwise.
      */
    public boolean hasEventTimeMs() {
      return fieldSetFlags()[3];
    }


    /**
      * Clears the value of the 'event_time_ms' field.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder clearEventTimeMs() {
      event_time_ms = null;
      fieldSetFlags()[3] = false;
      return this;
    }

    /**
      * Gets the value of the 'events' field.
      * @return The value.
      */
    public java.util.Map<java.lang.CharSequence,java.lang.CharSequence> getEvents() {
      return events;
    }


    /**
      * Sets the value of the 'events' field.
      * @param value The value of 'events'.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder setEvents(java.util.Map<java.lang.CharSequence,java.lang.CharSequence> value) {
      validate(fields()[4], value);
      this.events = value;
      fieldSetFlags()[4] = true;
      return this;
    }

    /**
      * Checks whether the 'events' field has been set.
      * @return True if the 'events' field has been set, false otherwise.
      */
    public boolean hasEvents() {
      return fieldSetFlags()[4];
    }


    /**
      * Clears the value of the 'events' field.
      * @return This builder.
      */
    public manning.devices.canonical.CanonicalValue.Builder clearEvents() {
      events = null;
      fieldSetFlags()[4] = false;
      return this;
    }

    @Override
    @SuppressWarnings("unchecked")
    public CanonicalValue build() {
      try {
        CanonicalValue record = new CanonicalValue();
        record.uuid = fieldSetFlags()[0] ? this.uuid : (java.lang.CharSequence) defaultValue(fields()[0]);
        record.regionId = fieldSetFlags()[1] ? this.regionId : (java.lang.Long) defaultValue(fields()[1]);
        record.arrival_time_ms = fieldSetFlags()[2] ? this.arrival_time_ms : (java.lang.Long) defaultValue(fields()[2]);
        record.event_time_ms = fieldSetFlags()[3] ? this.event_time_ms : (java.lang.Long) defaultValue(fields()[3]);
        record.events = fieldSetFlags()[4] ? this.events : (java.util.Map<java.lang.CharSequence,java.lang.CharSequence>) defaultValue(fields()[4]);
        return record;
      } catch (org.apache.avro.AvroMissingFieldException e) {
        throw e;
      } catch (java.lang.Exception e) {
        throw new org.apache.avro.AvroRuntimeException(e);
      }
    }
  }

  @SuppressWarnings("unchecked")
  private static final org.apache.avro.io.DatumWriter<CanonicalValue>
    WRITER$ = (org.apache.avro.io.DatumWriter<CanonicalValue>)MODEL$.createDatumWriter(SCHEMA$);

  @Override public void writeExternal(java.io.ObjectOutput out)
    throws java.io.IOException {
    WRITER$.write(this, SpecificData.getEncoder(out));
  }

  @SuppressWarnings("unchecked")
  private static final org.apache.avro.io.DatumReader<CanonicalValue>
    READER$ = (org.apache.avro.io.DatumReader<CanonicalValue>)MODEL$.createDatumReader(SCHEMA$);

  @Override public void readExternal(java.io.ObjectInput in)
    throws java.io.IOException {
    READER$.read(this, SpecificData.getDecoder(in));
  }

  @Override protected boolean hasCustomCoders() { return true; }

  @Override public void customEncode(org.apache.avro.io.Encoder out)
    throws java.io.IOException
  {
    out.writeString(this.uuid);

    if (this.regionId == null) {
      out.writeIndex(0);
      out.writeNull();
    } else {
      out.writeIndex(1);
      out.writeLong(this.regionId);
    }

    if (this.arrival_time_ms == null) {
      out.writeIndex(0);
      out.writeNull();
    } else {
      out.writeIndex(1);
      out.writeLong(this.arrival_time_ms);
    }

    if (this.event_time_ms == null) {
      out.writeIndex(0);
      out.writeNull();
    } else {
      out.writeIndex(1);
      out.writeLong(this.event_time_ms);
    }

    if (this.events == null) {
      out.writeIndex(0);
      out.writeNull();
    } else {
      out.writeIndex(1);
      long size0 = this.events.size();
      out.writeMapStart();
      out.setItemCount(size0);
      long actualSize0 = 0;
      for (java.util.Map.Entry<java.lang.CharSequence, java.lang.CharSequence> e0: this.events.entrySet()) {
        actualSize0++;
        out.startItem();
        out.writeString(e0.getKey());
        java.lang.CharSequence v0 = e0.getValue();
        out.writeString(v0);
      }
      out.writeMapEnd();
      if (actualSize0 != size0)
      throw new java.util.ConcurrentModificationException("Map-size written was " + size0 + ", but element count was " + actualSize0 + ".");
    }

  }

  @Override public void customDecode(org.apache.avro.io.ResolvingDecoder in)
    throws java.io.IOException
  {
    org.apache.avro.Schema.Field[] fieldOrder = in.readFieldOrderIfDiff();
    if (fieldOrder == null) {
      this.uuid = in.readString(this.uuid instanceof Utf8 ? (Utf8)this.uuid : null);

      if (in.readIndex() != 1) {
        in.readNull();
        this.regionId = null;
      } else {
        this.regionId = in.readLong();
      }

      if (in.readIndex() != 1) {
        in.readNull();
        this.arrival_time_ms = null;
      } else {
        this.arrival_time_ms = in.readLong();
      }

      if (in.readIndex() != 1) {
        in.readNull();
        this.event_time_ms = null;
      } else {
        this.event_time_ms = in.readLong();
      }

      if (in.readIndex() != 1) {
        in.readNull();
        this.events = null;
      } else {
        long size0 = in.readMapStart();
        java.util.Map<java.lang.CharSequence,java.lang.CharSequence> m0 = this.events; // Need fresh name due to limitation of macro system
        if (m0 == null) {
          m0 = new java.util.HashMap<java.lang.CharSequence,java.lang.CharSequence>((int)size0);
          this.events = m0;
        } else m0.clear();
        for ( ; 0 < size0; size0 = in.mapNext()) {
          for ( ; size0 != 0; size0--) {
            java.lang.CharSequence k0 = null;
            k0 = in.readString(k0 instanceof Utf8 ? (Utf8)k0 : null);
            java.lang.CharSequence v0 = null;
            v0 = in.readString(v0 instanceof Utf8 ? (Utf8)v0 : null);
            m0.put(k0, v0);
          }
        }
      }

    } else {
      for (int i = 0; i < 5; i++) {
        switch (fieldOrder[i].pos()) {
        case 0:
          this.uuid = in.readString(this.uuid instanceof Utf8 ? (Utf8)this.uuid : null);
          break;

        case 1:
          if (in.readIndex() != 1) {
            in.readNull();
            this.regionId = null;
          } else {
            this.regionId = in.readLong();
          }
          break;

        case 2:
          if (in.readIndex() != 1) {
            in.readNull();
            this.arrival_time_ms = null;
          } else {
            this.arrival_time_ms = in.readLong();
          }
          break;

        case 3:
          if (in.readIndex() != 1) {
            in.readNull();
            this.event_time_ms = null;
          } else {
            this.event_time_ms = in.readLong();
          }
          break;

        case 4:
          if (in.readIndex() != 1) {
            in.readNull();
            this.events = null;
          } else {
            long size0 = in.readMapStart();
            java.util.Map<java.lang.CharSequence,java.lang.CharSequence> m0 = this.events; // Need fresh name due to limitation of macro system
            if (m0 == null) {
              m0 = new java.util.HashMap<java.lang.CharSequence,java.lang.CharSequence>((int)size0);
              this.events = m0;
            } else m0.clear();
            for ( ; 0 < size0; size0 = in.mapNext()) {
              for ( ; size0 != 0; size0--) {
                java.lang.CharSequence k0 = null;
                k0 = in.readString(k0 instanceof Utf8 ? (Utf8)k0 : null);
                java.lang.CharSequence v0 = null;
                v0 = in.readString(v0 instanceof Utf8 ? (Utf8)v0 : null);
                m0.put(k0, v0);
              }
            }
          }
          break;

        default:
          throw new java.io.IOException("Corrupt ResolvingDecoder.");
        }
      }
    }
  }
}










