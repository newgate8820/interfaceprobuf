// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: hbase.proto

package im.pt.message.proto;

public interface SearchMessageRespOrBuilder extends
    // @@protoc_insertion_point(interface_extends:messagedbold.SearchMessageResp)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * 消息切片
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  java.util.List<im.pt.message.proto.UserMessage> 
      getMessagesList();
  /**
   * <pre>
   * 消息切片
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  im.pt.message.proto.UserMessage getMessages(int index);
  /**
   * <pre>
   * 消息切片
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  int getMessagesCount();
  /**
   * <pre>
   * 消息切片
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  java.util.List<? extends im.pt.message.proto.UserMessageOrBuilder> 
      getMessagesOrBuilderList();
  /**
   * <pre>
   * 消息切片
   * </pre>
   *
   * <code>repeated .messagedbold.UserMessage Messages = 1;</code>
   */
  im.pt.message.proto.UserMessageOrBuilder getMessagesOrBuilder(
      int index);
}
