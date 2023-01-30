// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: hbase.proto

package im.pt.message.proto;

/**
 * Protobuf type {@code messagedbold.GetUserMsgUUIDReq}
 */
public  final class GetUserMsgUUIDReq extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:messagedbold.GetUserMsgUUIDReq)
    GetUserMsgUUIDReqOrBuilder {
private static final long serialVersionUID = 0L;
  // Use GetUserMsgUUIDReq.newBuilder() to construct.
  private GetUserMsgUUIDReq(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GetUserMsgUUIDReq() {
    userID_ = 0;
    msgID_ = 0;
    dType_ = 0;
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private GetUserMsgUUIDReq(
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

            msgID_ = input.readInt32();
            break;
          }
          case 24: {
            int rawValue = input.readEnum();

            dType_ = rawValue;
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
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgUUIDReq_descriptor;
  }

  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgUUIDReq_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            im.pt.message.proto.GetUserMsgUUIDReq.class, im.pt.message.proto.GetUserMsgUUIDReq.Builder.class);
  }

  public static final int USERID_FIELD_NUMBER = 1;
  private int userID_;
  /**
   * <code>int32 UserID = 1;</code>
   */
  public int getUserID() {
    return userID_;
  }

  public static final int MSGID_FIELD_NUMBER = 2;
  private int msgID_;
  /**
   * <code>int32 MsgID = 2;</code>
   */
  public int getMsgID() {
    return msgID_;
  }

  public static final int DTYPE_FIELD_NUMBER = 3;
  private int dType_;
  /**
   * <pre>
   * 数据库类型
   * </pre>
   *
   * <code>.messagedbold.DATABASETYPE DType = 3;</code>
   */
  public int getDTypeValue() {
    return dType_;
  }
  /**
   * <pre>
   * 数据库类型
   * </pre>
   *
   * <code>.messagedbold.DATABASETYPE DType = 3;</code>
   */
  public im.pt.message.proto.DATABASETYPE getDType() {
    im.pt.message.proto.DATABASETYPE result = im.pt.message.proto.DATABASETYPE.valueOf(dType_);
    return result == null ? im.pt.message.proto.DATABASETYPE.UNRECOGNIZED : result;
  }

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
    if (userID_ != 0) {
      output.writeInt32(1, userID_);
    }
    if (msgID_ != 0) {
      output.writeInt32(2, msgID_);
    }
    if (dType_ != im.pt.message.proto.DATABASETYPE._.getNumber()) {
      output.writeEnum(3, dType_);
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
    if (msgID_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, msgID_);
    }
    if (dType_ != im.pt.message.proto.DATABASETYPE._.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(3, dType_);
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
    if (!(obj instanceof im.pt.message.proto.GetUserMsgUUIDReq)) {
      return super.equals(obj);
    }
    im.pt.message.proto.GetUserMsgUUIDReq other = (im.pt.message.proto.GetUserMsgUUIDReq) obj;

    boolean result = true;
    result = result && (getUserID()
        == other.getUserID());
    result = result && (getMsgID()
        == other.getMsgID());
    result = result && dType_ == other.dType_;
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
    hash = (37 * hash) + MSGID_FIELD_NUMBER;
    hash = (53 * hash) + getMsgID();
    hash = (37 * hash) + DTYPE_FIELD_NUMBER;
    hash = (53 * hash) + dType_;
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static im.pt.message.proto.GetUserMsgUUIDReq parseFrom(
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
  public static Builder newBuilder(im.pt.message.proto.GetUserMsgUUIDReq prototype) {
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
   * Protobuf type {@code messagedbold.GetUserMsgUUIDReq}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:messagedbold.GetUserMsgUUIDReq)
      im.pt.message.proto.GetUserMsgUUIDReqOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgUUIDReq_descriptor;
    }

    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgUUIDReq_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              im.pt.message.proto.GetUserMsgUUIDReq.class, im.pt.message.proto.GetUserMsgUUIDReq.Builder.class);
    }

    // Construct using im.pt.message.proto.GetUserMsgUUIDReq.newBuilder()
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

      msgID_ = 0;

      dType_ = 0;

      return this;
    }

    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMsgUUIDReq_descriptor;
    }

    public im.pt.message.proto.GetUserMsgUUIDReq getDefaultInstanceForType() {
      return im.pt.message.proto.GetUserMsgUUIDReq.getDefaultInstance();
    }

    public im.pt.message.proto.GetUserMsgUUIDReq build() {
      im.pt.message.proto.GetUserMsgUUIDReq result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    public im.pt.message.proto.GetUserMsgUUIDReq buildPartial() {
      im.pt.message.proto.GetUserMsgUUIDReq result = new im.pt.message.proto.GetUserMsgUUIDReq(this);
      result.userID_ = userID_;
      result.msgID_ = msgID_;
      result.dType_ = dType_;
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
      if (other instanceof im.pt.message.proto.GetUserMsgUUIDReq) {
        return mergeFrom((im.pt.message.proto.GetUserMsgUUIDReq)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(im.pt.message.proto.GetUserMsgUUIDReq other) {
      if (other == im.pt.message.proto.GetUserMsgUUIDReq.getDefaultInstance()) return this;
      if (other.getUserID() != 0) {
        setUserID(other.getUserID());
      }
      if (other.getMsgID() != 0) {
        setMsgID(other.getMsgID());
      }
      if (other.dType_ != 0) {
        setDTypeValue(other.getDTypeValue());
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
      im.pt.message.proto.GetUserMsgUUIDReq parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (im.pt.message.proto.GetUserMsgUUIDReq) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

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

    private int msgID_ ;
    /**
     * <code>int32 MsgID = 2;</code>
     */
    public int getMsgID() {
      return msgID_;
    }
    /**
     * <code>int32 MsgID = 2;</code>
     */
    public Builder setMsgID(int value) {
      
      msgID_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 MsgID = 2;</code>
     */
    public Builder clearMsgID() {
      
      msgID_ = 0;
      onChanged();
      return this;
    }

    private int dType_ = 0;
    /**
     * <pre>
     * 数据库类型
     * </pre>
     *
     * <code>.messagedbold.DATABASETYPE DType = 3;</code>
     */
    public int getDTypeValue() {
      return dType_;
    }
    /**
     * <pre>
     * 数据库类型
     * </pre>
     *
     * <code>.messagedbold.DATABASETYPE DType = 3;</code>
     */
    public Builder setDTypeValue(int value) {
      dType_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 数据库类型
     * </pre>
     *
     * <code>.messagedbold.DATABASETYPE DType = 3;</code>
     */
    public im.pt.message.proto.DATABASETYPE getDType() {
      im.pt.message.proto.DATABASETYPE result = im.pt.message.proto.DATABASETYPE.valueOf(dType_);
      return result == null ? im.pt.message.proto.DATABASETYPE.UNRECOGNIZED : result;
    }
    /**
     * <pre>
     * 数据库类型
     * </pre>
     *
     * <code>.messagedbold.DATABASETYPE DType = 3;</code>
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
     * <pre>
     * 数据库类型
     * </pre>
     *
     * <code>.messagedbold.DATABASETYPE DType = 3;</code>
     */
    public Builder clearDType() {
      
      dType_ = 0;
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


    // @@protoc_insertion_point(builder_scope:messagedbold.GetUserMsgUUIDReq)
  }

  // @@protoc_insertion_point(class_scope:messagedbold.GetUserMsgUUIDReq)
  private static final im.pt.message.proto.GetUserMsgUUIDReq DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new im.pt.message.proto.GetUserMsgUUIDReq();
  }

  public static im.pt.message.proto.GetUserMsgUUIDReq getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GetUserMsgUUIDReq>
      PARSER = new com.google.protobuf.AbstractParser<GetUserMsgUUIDReq>() {
    public GetUserMsgUUIDReq parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new GetUserMsgUUIDReq(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<GetUserMsgUUIDReq> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GetUserMsgUUIDReq> getParserForType() {
    return PARSER;
  }

  public im.pt.message.proto.GetUserMsgUUIDReq getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

