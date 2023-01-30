// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: hbase.proto

package im.pt.message.proto;

/**
 * Protobuf type {@code messagedbold.GetUserMsgListByUUIDSReq}
 */
public  final class GetUserMsgListByUUIDSReq extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:messagedbold.GetUserMsgListByUUIDSReq)
    GetUserMsgListByUUIDSReqOrBuilder {
private static final long serialVersionUID = 0L;
  // Use GetUserMsgListByUUIDSReq.newBuilder() to construct.
  private GetUserMsgListByUUIDSReq(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GetUserMsgListByUUIDSReq() {
    userID_ = 0;
    dType_ = 0;
    uUIDs_ = java.util.Collections.emptyList();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private GetUserMsgListByUUIDSReq(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          default: {
            if (!parseUnknownFieldProto3(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
          case 8: {

            userID_ = input.readInt32();
            break;
          }
          case 16: {
            int rawValue = input.readEnum();

            dType_ = rawValue;
            break;
          }
          case 24: {
            if (!((mutable_bitField0_ & 0x00000004) == 0x00000004)) {
              uUIDs_ = new java.util.ArrayList<java.lang.Long>();
              mutable_bitField0_ |= 0x00000004;
            }
            uUIDs_.add(input.readInt64());
            break;
          }
          case 26: {
            int length = input.readRawVarint32();
            int limit = input.pushLimit(length);
            if (!((mutable_bitField0_ & 0x00000004) == 0x00000004) && input.getBytesUntilLimit() > 0) {
              uUIDs_ = new java.util.ArrayList<java.lang.Long>();
              mutable_bitField0_ |= 0x00000004;
            }
            while (input.getBytesUntilLimit() > 0) {
              uUIDs_.add(input.readInt64());
            }
            input.popLimit(limit);
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      if (((mutable_bitField0_ & 0x00000004) == 0x00000004)) {
        uUIDs_ = java.util.Collections.unmodifiableList(uUIDs_);
      }
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgListByUUIDSReq_descriptor;
  }

  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgListByUUIDSReq_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            im.pt.message.proto.GetUserMsgListByUUIDSReq.class, im.pt.message.proto.GetUserMsgListByUUIDSReq.Builder.class);
  }

  private int bitField0_;
  public static final int USERID_FIELD_NUMBER = 1;
  private int userID_;
  /**
   * <code>int32 UserID = 1;</code>
   */
  public int getUserID() {
    return userID_;
  }

  public static final int DTYPE_FIELD_NUMBER = 2;
  private int dType_;
  /**
   * <code>.messagedbold.DATABASETYPE DType = 2;</code>
   */
  public int getDTypeValue() {
    return dType_;
  }
  /**
   * <code>.messagedbold.DATABASETYPE DType = 2;</code>
   */
  public im.pt.message.proto.DATABASETYPE getDType() {
    im.pt.message.proto.DATABASETYPE result = im.pt.message.proto.DATABASETYPE.valueOf(dType_);
    return result == null ? im.pt.message.proto.DATABASETYPE.UNRECOGNIZED : result;
  }

  public static final int UUIDS_FIELD_NUMBER = 3;
  private java.util.List<java.lang.Long> uUIDs_;
  /**
   * <code>repeated int64 UUIDs = 3;</code>
   */
  public java.util.List<java.lang.Long>
      getUUIDsList() {
    return uUIDs_;
  }
  /**
   * <code>repeated int64 UUIDs = 3;</code>
   */
  public int getUUIDsCount() {
    return uUIDs_.size();
  }
  /**
   * <code>repeated int64 UUIDs = 3;</code>
   */
  public long getUUIDs(int index) {
    return uUIDs_.get(index);
  }
  private int uUIDsMemoizedSerializedSize = -1;

  private byte memoizedIsInitialized = -1;
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    getSerializedSize();
    if (userID_ != 0) {
      output.writeInt32(1, userID_);
    }
    if (dType_ != im.pt.message.proto.DATABASETYPE._.getNumber()) {
      output.writeEnum(2, dType_);
    }
    if (getUUIDsList().size() > 0) {
      output.writeUInt32NoTag(26);
      output.writeUInt32NoTag(uUIDsMemoizedSerializedSize);
    }
    for (int i = 0; i < uUIDs_.size(); i++) {
      output.writeInt64NoTag(uUIDs_.get(i));
    }
    unknownFields.writeTo(output);
  }

  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (userID_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(1, userID_);
    }
    if (dType_ != im.pt.message.proto.DATABASETYPE._.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(2, dType_);
    }
    {
      int dataSize = 0;
      for (int i = 0; i < uUIDs_.size(); i++) {
        dataSize += com.google.protobuf.CodedOutputStream
          .computeInt64SizeNoTag(uUIDs_.get(i));
      }
      size += dataSize;
      if (!getUUIDsList().isEmpty()) {
        size += 1;
        size += com.google.protobuf.CodedOutputStream
            .computeInt32SizeNoTag(dataSize);
      }
      uUIDsMemoizedSerializedSize = dataSize;
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof im.pt.message.proto.GetUserMsgListByUUIDSReq)) {
      return super.equals(obj);
    }
    im.pt.message.proto.GetUserMsgListByUUIDSReq other = (im.pt.message.proto.GetUserMsgListByUUIDSReq) obj;

    boolean result = true;
    result = result && (getUserID()
        == other.getUserID());
    result = result && dType_ == other.dType_;
    result = result && getUUIDsList()
        .equals(other.getUUIDsList());
    result = result && unknownFields.equals(other.unknownFields);
    return result;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + USERID_FIELD_NUMBER;
    hash = (53 * hash) + getUserID();
    hash = (37 * hash) + DTYPE_FIELD_NUMBER;
    hash = (53 * hash) + dType_;
    if (getUUIDsCount() > 0) {
      hash = (37 * hash) + UUIDS_FIELD_NUMBER;
      hash = (53 * hash) + getUUIDsList().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static im.pt.message.proto.GetUserMsgListByUUIDSReq parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(im.pt.message.proto.GetUserMsgListByUUIDSReq prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code messagedbold.GetUserMsgListByUUIDSReq}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:messagedbold.GetUserMsgListByUUIDSReq)
      im.pt.message.proto.GetUserMsgListByUUIDSReqOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgListByUUIDSReq_descriptor;
    }

    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgListByUUIDSReq_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              im.pt.message.proto.GetUserMsgListByUUIDSReq.class, im.pt.message.proto.GetUserMsgListByUUIDSReq.Builder.class);
    }

    // Construct using im.pt.message.proto.GetUserMsgListByUUIDSReq.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    public Builder clear() {
      super.clear();
      userID_ = 0;

      dType_ = 0;

      uUIDs_ = java.util.Collections.emptyList();
      bitField0_ = (bitField0_ & ~0x00000004);
      return this;
    }

    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgListByUUIDSReq_descriptor;
    }

    public im.pt.message.proto.GetUserMsgListByUUIDSReq getDefaultInstanceForType() {
      return im.pt.message.proto.GetUserMsgListByUUIDSReq.getDefaultInstance();
    }

    public im.pt.message.proto.GetUserMsgListByUUIDSReq build() {
      im.pt.message.proto.GetUserMsgListByUUIDSReq result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    public im.pt.message.proto.GetUserMsgListByUUIDSReq buildPartial() {
      im.pt.message.proto.GetUserMsgListByUUIDSReq result = new im.pt.message.proto.GetUserMsgListByUUIDSReq(this);
      int from_bitField0_ = bitField0_;
      int to_bitField0_ = 0;
      result.userID_ = userID_;
      result.dType_ = dType_;
      if (((bitField0_ & 0x00000004) == 0x00000004)) {
        uUIDs_ = java.util.Collections.unmodifiableList(uUIDs_);
        bitField0_ = (bitField0_ & ~0x00000004);
      }
      result.uUIDs_ = uUIDs_;
      result.bitField0_ = to_bitField0_;
      onBuilt();
      return result;
    }

    public Builder clone() {
      return (Builder) super.clone();
    }
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return (Builder) super.setField(field, value);
    }
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return (Builder) super.clearField(field);
    }
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return (Builder) super.clearOneof(oneof);
    }
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return (Builder) super.setRepeatedField(field, index, value);
    }
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return (Builder) super.addRepeatedField(field, value);
    }
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof im.pt.message.proto.GetUserMsgListByUUIDSReq) {
        return mergeFrom((im.pt.message.proto.GetUserMsgListByUUIDSReq)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(im.pt.message.proto.GetUserMsgListByUUIDSReq other) {
      if (other == im.pt.message.proto.GetUserMsgListByUUIDSReq.getDefaultInstance()) return this;
      if (other.getUserID() != 0) {
        setUserID(other.getUserID());
      }
      if (other.dType_ != 0) {
        setDTypeValue(other.getDTypeValue());
      }
      if (!other.uUIDs_.isEmpty()) {
        if (uUIDs_.isEmpty()) {
          uUIDs_ = other.uUIDs_;
          bitField0_ = (bitField0_ & ~0x00000004);
        } else {
          ensureUUIDsIsMutable();
          uUIDs_.addAll(other.uUIDs_);
        }
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    public final boolean isInitialized() {
      return true;
    }

    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      im.pt.message.proto.GetUserMsgListByUUIDSReq parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (im.pt.message.proto.GetUserMsgListByUUIDSReq) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private int userID_ ;
    /**
     * <code>int32 UserID = 1;</code>
     */
    public int getUserID() {
      return userID_;
    }
    /**
     * <code>int32 UserID = 1;</code>
     */
    public Builder setUserID(int value) {
      
      userID_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 UserID = 1;</code>
     */
    public Builder clearUserID() {
      
      userID_ = 0;
      onChanged();
      return this;
    }

    private int dType_ = 0;
    /**
     * <code>.messagedbold.DATABASETYPE DType = 2;</code>
     */
    public int getDTypeValue() {
      return dType_;
    }
    /**
     * <code>.messagedbold.DATABASETYPE DType = 2;</code>
     */
    public Builder setDTypeValue(int value) {
      dType_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>.messagedbold.DATABASETYPE DType = 2;</code>
     */
    public im.pt.message.proto.DATABASETYPE getDType() {
      im.pt.message.proto.DATABASETYPE result = im.pt.message.proto.DATABASETYPE.valueOf(dType_);
      return result == null ? im.pt.message.proto.DATABASETYPE.UNRECOGNIZED : result;
    }
    /**
     * <code>.messagedbold.DATABASETYPE DType = 2;</code>
     */
    public Builder setDType(im.pt.message.proto.DATABASETYPE value) {
      if (value == null) {
        throw new NullPointerException();
      }
      
      dType_ = value.getNumber();
      onChanged();
      return this;
    }
    /**
     * <code>.messagedbold.DATABASETYPE DType = 2;</code>
     */
    public Builder clearDType() {
      
      dType_ = 0;
      onChanged();
      return this;
    }

    private java.util.List<java.lang.Long> uUIDs_ = java.util.Collections.emptyList();
    private void ensureUUIDsIsMutable() {
      if (!((bitField0_ & 0x00000004) == 0x00000004)) {
        uUIDs_ = new java.util.ArrayList<java.lang.Long>(uUIDs_);
        bitField0_ |= 0x00000004;
       }
    }
    /**
     * <code>repeated int64 UUIDs = 3;</code>
     */
    public java.util.List<java.lang.Long>
        getUUIDsList() {
      return java.util.Collections.unmodifiableList(uUIDs_);
    }
    /**
     * <code>repeated int64 UUIDs = 3;</code>
     */
    public int getUUIDsCount() {
      return uUIDs_.size();
    }
    /**
     * <code>repeated int64 UUIDs = 3;</code>
     */
    public long getUUIDs(int index) {
      return uUIDs_.get(index);
    }
    /**
     * <code>repeated int64 UUIDs = 3;</code>
     */
    public Builder setUUIDs(
        int index, long value) {
      ensureUUIDsIsMutable();
      uUIDs_.set(index, value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated int64 UUIDs = 3;</code>
     */
    public Builder addUUIDs(long value) {
      ensureUUIDsIsMutable();
      uUIDs_.add(value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated int64 UUIDs = 3;</code>
     */
    public Builder addAllUUIDs(
        java.lang.Iterable<? extends java.lang.Long> values) {
      ensureUUIDsIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, uUIDs_);
      onChanged();
      return this;
    }
    /**
     * <code>repeated int64 UUIDs = 3;</code>
     */
    public Builder clearUUIDs() {
      uUIDs_ = java.util.Collections.emptyList();
      bitField0_ = (bitField0_ & ~0x00000004);
      onChanged();
      return this;
    }
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFieldsProto3(unknownFields);
    }

    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:messagedbold.GetUserMsgListByUUIDSReq)
  }

  // @@protoc_insertion_point(class_scope:messagedbold.GetUserMsgListByUUIDSReq)
  private static final im.pt.message.proto.GetUserMsgListByUUIDSReq DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new im.pt.message.proto.GetUserMsgListByUUIDSReq();
  }

  public static im.pt.message.proto.GetUserMsgListByUUIDSReq getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GetUserMsgListByUUIDSReq>
      PARSER = new com.google.protobuf.AbstractParser<GetUserMsgListByUUIDSReq>() {
    public GetUserMsgListByUUIDSReq parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new GetUserMsgListByUUIDSReq(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<GetUserMsgListByUUIDSReq> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GetUserMsgListByUUIDSReq> getParserForType() {
    return PARSER;
  }

  public im.pt.message.proto.GetUserMsgListByUUIDSReq getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

