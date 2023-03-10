// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: hbase.proto

package im.pt.message.proto;

/**
 * Protobuf type {@code messagedbold.GetUserMessageHistoryResp}
 */
public  final class GetUserMessageHistoryResp extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:messagedbold.GetUserMessageHistoryResp)
    GetUserMessageHistoryRespOrBuilder {
private static final long serialVersionUID = 0L;
  // Use GetUserMessageHistoryResp.newBuilder() to construct.
  private GetUserMessageHistoryResp(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GetUserMessageHistoryResp() {
    messages_ = java.util.Collections.emptyList();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private GetUserMessageHistoryResp(
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
          case 10: {
            if (!((mutable_bitField0_ & 0x00000001) == 0x00000001)) {
              messages_ = new java.util.ArrayList<im.pt.message.proto.UserMessage>();
              mutable_bitField0_ |= 0x00000001;
            }
            messages_.add(
                input.readMessage(im.pt.message.proto.UserMessage.parser(), extensionRegistry));
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
      if (((mutable_bitField0_ & 0x00000001) == 0x00000001)) {
        messages_ = java.util.Collections.unmodifiableList(messages_);
      }
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMessageHistoryResp_descriptor;
  }

  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMessageHistoryResp_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            im.pt.message.proto.GetUserMessageHistoryResp.class, im.pt.message.proto.GetUserMessageHistoryResp.Builder.class);
  }

  public static final int MESSAGES_FIELD_NUMBER = 1;
  private java.util.List<im.pt.message.proto.UserMessage> messages_;
  /**
   * <pre>
   * ????????????
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  public java.util.List<im.pt.message.proto.UserMessage> getMessagesList() {
    return messages_;
  }
  /**
   * <pre>
   * ????????????
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  public java.util.List<? extends im.pt.message.proto.UserMessageOrBuilder> 
      getMessagesOrBuilderList() {
    return messages_;
  }
  /**
   * <pre>
   * ????????????
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  public int getMessagesCount() {
    return messages_.size();
  }
  /**
   * <pre>
   * ????????????
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  public im.pt.message.proto.UserMessage getMessages(int index) {
    return messages_.get(index);
  }
  /**
   * <pre>
   * ????????????
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  public im.pt.message.proto.UserMessageOrBuilder getMessagesOrBuilder(
      int index) {
    return messages_.get(index);
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
    for (int i = 0; i < messages_.size(); i++) {
      output.writeMessage(1, messages_.get(i));
    }
    unknownFields.writeTo(output);
  }

  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    for (int i = 0; i < messages_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, messages_.get(i));
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
    if (!(obj instanceof im.pt.message.proto.GetUserMessageHistoryResp)) {
      return super.equals(obj);
    }
    im.pt.message.proto.GetUserMessageHistoryResp other = (im.pt.message.proto.GetUserMessageHistoryResp) obj;

    boolean result = true;
    result = result && getMessagesList()
        .equals(other.getMessagesList());
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
    if (getMessagesCount() > 0) {
      hash = (37 * hash) + MESSAGES_FIELD_NUMBER;
      hash = (53 * hash) + getMessagesList().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static im.pt.message.proto.GetUserMessageHistoryResp parseFrom(
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
  public static Builder newBuilder(im.pt.message.proto.GetUserMessageHistoryResp prototype) {
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
   * Protobuf type {@code messagedbold.GetUserMessageHistoryResp}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:messagedbold.GetUserMessageHistoryResp)
      im.pt.message.proto.GetUserMessageHistoryRespOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMessageHistoryResp_descriptor;
    }

    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMessageHistoryResp_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              im.pt.message.proto.GetUserMessageHistoryResp.class, im.pt.message.proto.GetUserMessageHistoryResp.Builder.class);
    }

    // Construct using im.pt.message.proto.GetUserMessageHistoryResp.newBuilder()
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
        getMessagesFieldBuilder();
      }
    }
    public Builder clear() {
      super.clear();
      if (messagesBuilder_ == null) {
        messages_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
      } else {
        messagesBuilder_.clear();
      }
      return this;
    }

    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return im.pt.message.proto.Hbase.internal_static_messagedbold_GetUserMessageHistoryResp_descriptor;
    }

    public im.pt.message.proto.GetUserMessageHistoryResp getDefaultInstanceForType() {
      return im.pt.message.proto.GetUserMessageHistoryResp.getDefaultInstance();
    }

    public im.pt.message.proto.GetUserMessageHistoryResp build() {
      im.pt.message.proto.GetUserMessageHistoryResp result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    public im.pt.message.proto.GetUserMessageHistoryResp buildPartial() {
      im.pt.message.proto.GetUserMessageHistoryResp result = new im.pt.message.proto.GetUserMessageHistoryResp(this);
      int from_bitField0_ = bitField0_;
      if (messagesBuilder_ == null) {
        if (((bitField0_ & 0x00000001) == 0x00000001)) {
          messages_ = java.util.Collections.unmodifiableList(messages_);
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.messages_ = messages_;
      } else {
        result.messages_ = messagesBuilder_.build();
      }
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
      if (other instanceof im.pt.message.proto.GetUserMessageHistoryResp) {
        return mergeFrom((im.pt.message.proto.GetUserMessageHistoryResp)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(im.pt.message.proto.GetUserMessageHistoryResp other) {
      if (other == im.pt.message.proto.GetUserMessageHistoryResp.getDefaultInstance()) return this;
      if (messagesBuilder_ == null) {
        if (!other.messages_.isEmpty()) {
          if (messages_.isEmpty()) {
            messages_ = other.messages_;
            bitField0_ = (bitField0_ & ~0x00000001);
          } else {
            ensureMessagesIsMutable();
            messages_.addAll(other.messages_);
          }
          onChanged();
        }
      } else {
        if (!other.messages_.isEmpty()) {
          if (messagesBuilder_.isEmpty()) {
            messagesBuilder_.dispose();
            messagesBuilder_ = null;
            messages_ = other.messages_;
            bitField0_ = (bitField0_ & ~0x00000001);
            messagesBuilder_ = 
              com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                 getMessagesFieldBuilder() : null;
          } else {
            messagesBuilder_.addAllMessages(other.messages_);
          }
        }
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
      im.pt.message.proto.GetUserMessageHistoryResp parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (im.pt.message.proto.GetUserMessageHistoryResp) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.util.List<im.pt.message.proto.UserMessage> messages_ =
      java.util.Collections.emptyList();
    private void ensureMessagesIsMutable() {
      if (!((bitField0_ & 0x00000001) == 0x00000001)) {
        messages_ = new java.util.ArrayList<im.pt.message.proto.UserMessage>(messages_);
        bitField0_ |= 0x00000001;
       }
    }

    private com.google.protobuf.RepeatedFieldBuilderV3<
        im.pt.message.proto.UserMessage, im.pt.message.proto.UserMessage.Builder, im.pt.message.proto.UserMessageOrBuilder> messagesBuilder_;

    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public java.util.List<im.pt.message.proto.UserMessage> getMessagesList() {
      if (messagesBuilder_ == null) {
        return java.util.Collections.unmodifiableList(messages_);
      } else {
        return messagesBuilder_.getMessageList();
      }
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public int getMessagesCount() {
      if (messagesBuilder_ == null) {
        return messages_.size();
      } else {
        return messagesBuilder_.getCount();
      }
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public im.pt.message.proto.UserMessage getMessages(int index) {
      if (messagesBuilder_ == null) {
        return messages_.get(index);
      } else {
        return messagesBuilder_.getMessage(index);
      }
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public Builder setMessages(
        int index, im.pt.message.proto.UserMessage value) {
      if (messagesBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureMessagesIsMutable();
        messages_.set(index, value);
        onChanged();
      } else {
        messagesBuilder_.setMessage(index, value);
      }
      return this;
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public Builder setMessages(
        int index, im.pt.message.proto.UserMessage.Builder builderForValue) {
      if (messagesBuilder_ == null) {
        ensureMessagesIsMutable();
        messages_.set(index, builderForValue.build());
        onChanged();
      } else {
        messagesBuilder_.setMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public Builder addMessages(im.pt.message.proto.UserMessage value) {
      if (messagesBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureMessagesIsMutable();
        messages_.add(value);
        onChanged();
      } else {
        messagesBuilder_.addMessage(value);
      }
      return this;
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public Builder addMessages(
        int index, im.pt.message.proto.UserMessage value) {
      if (messagesBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureMessagesIsMutable();
        messages_.add(index, value);
        onChanged();
      } else {
        messagesBuilder_.addMessage(index, value);
      }
      return this;
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public Builder addMessages(
        im.pt.message.proto.UserMessage.Builder builderForValue) {
      if (messagesBuilder_ == null) {
        ensureMessagesIsMutable();
        messages_.add(builderForValue.build());
        onChanged();
      } else {
        messagesBuilder_.addMessage(builderForValue.build());
      }
      return this;
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public Builder addMessages(
        int index, im.pt.message.proto.UserMessage.Builder builderForValue) {
      if (messagesBuilder_ == null) {
        ensureMessagesIsMutable();
        messages_.add(index, builderForValue.build());
        onChanged();
      } else {
        messagesBuilder_.addMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public Builder addAllMessages(
        java.lang.Iterable<? extends im.pt.message.proto.UserMessage> values) {
      if (messagesBuilder_ == null) {
        ensureMessagesIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, messages_);
        onChanged();
      } else {
        messagesBuilder_.addAllMessages(values);
      }
      return this;
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public Builder clearMessages() {
      if (messagesBuilder_ == null) {
        messages_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
      } else {
        messagesBuilder_.clear();
      }
      return this;
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public Builder removeMessages(int index) {
      if (messagesBuilder_ == null) {
        ensureMessagesIsMutable();
        messages_.remove(index);
        onChanged();
      } else {
        messagesBuilder_.remove(index);
      }
      return this;
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public im.pt.message.proto.UserMessage.Builder getMessagesBuilder(
        int index) {
      return getMessagesFieldBuilder().getBuilder(index);
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public im.pt.message.proto.UserMessageOrBuilder getMessagesOrBuilder(
        int index) {
      if (messagesBuilder_ == null) {
        return messages_.get(index);  } else {
        return messagesBuilder_.getMessageOrBuilder(index);
      }
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public java.util.List<? extends im.pt.message.proto.UserMessageOrBuilder> 
         getMessagesOrBuilderList() {
      if (messagesBuilder_ != null) {
        return messagesBuilder_.getMessageOrBuilderList();
      } else {
        return java.util.Collections.unmodifiableList(messages_);
      }
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public im.pt.message.proto.UserMessage.Builder addMessagesBuilder() {
      return getMessagesFieldBuilder().addBuilder(
          im.pt.message.proto.UserMessage.getDefaultInstance());
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public im.pt.message.proto.UserMessage.Builder addMessagesBuilder(
        int index) {
      return getMessagesFieldBuilder().addBuilder(
          index, im.pt.message.proto.UserMessage.getDefaultInstance());
    }
    /**
     * <pre>
     * ????????????
     * </pre>
     *
     * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
     */
    public java.util.List<im.pt.message.proto.UserMessage.Builder> 
         getMessagesBuilderList() {
      return getMessagesFieldBuilder().getBuilderList();
    }
    private com.google.protobuf.RepeatedFieldBuilderV3<
        im.pt.message.proto.UserMessage, im.pt.message.proto.UserMessage.Builder, im.pt.message.proto.UserMessageOrBuilder> 
        getMessagesFieldBuilder() {
      if (messagesBuilder_ == null) {
        messagesBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
            im.pt.message.proto.UserMessage, im.pt.message.proto.UserMessage.Builder, im.pt.message.proto.UserMessageOrBuilder>(
                messages_,
                ((bitField0_ & 0x00000001) == 0x00000001),
                getParentForChildren(),
                isClean());
        messages_ = null;
      }
      return messagesBuilder_;
    }
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFieldsProto3(unknownFields);
    }

    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:messagedbold.GetUserMessageHistoryResp)
  }

  // @@protoc_insertion_point(class_scope:messagedbold.GetUserMessageHistoryResp)
  private static final im.pt.message.proto.GetUserMessageHistoryResp DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new im.pt.message.proto.GetUserMessageHistoryResp();
  }

  public static im.pt.message.proto.GetUserMessageHistoryResp getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GetUserMessageHistoryResp>
      PARSER = new com.google.protobuf.AbstractParser<GetUserMessageHistoryResp>() {
    public GetUserMessageHistoryResp parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new GetUserMessageHistoryResp(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<GetUserMessageHistoryResp> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GetUserMessageHistoryResp> getParserForType() {
    return PARSER;
  }

  public im.pt.message.proto.GetUserMessageHistoryResp getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

