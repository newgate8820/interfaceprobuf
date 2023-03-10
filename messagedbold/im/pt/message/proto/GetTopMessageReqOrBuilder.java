// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: hbase.proto

package im.pt.message.proto;

public interface GetTopMessageReqOrBuilder extends
    // @@protoc_insertion_point(interface_extends:messagedbold.GetTopMessageReq)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * 消息拥有者
   * </pre>
   *
   * <code>int32 UserID = 1;</code>
   */
  int getUserID();

  /**
   * <pre>
   * 消息接收者
   * </pre>
   *
   * <code>int32 WithID = 2;</code>
   */
  int getWithID();

  /**
   * <pre>
   * 对话框类型
   * </pre>
   *
   * <code>int32 WithType = 3;</code>
   */
  int getWithType();

  /**
   * <pre>
   * 数据库类型
   * </pre>
   *
   * <code>.messagedbold.DATABASETYPE DType = 4;</code>
   */
  int getDTypeValue();
  /**
   * <pre>
   * 数据库类型
   * </pre>
   *
   * <code>.messagedbold.DATABASETYPE DType = 4;</code>
   */
  im.pt.message.proto.DATABASETYPE getDType();
}
