// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: hbase.proto

package im.pt.message.proto;

public interface DeleteUserMsgByUserIDWithIDWithTypeReqOrBuilder extends
    // @@protoc_insertion_point(interface_extends:messagedbold.DeleteUserMsgByUserIDWithIDWithTypeReq)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int32 UserID = 1;</code>
   */
  int getUserID();

  /**
   * <code>int32 WithID = 2;</code>
   */
  int getWithID();

  /**
   * <code>int32 WithType = 3;</code>
   */
  int getWithType();

  /**
   * <code>.messagedbold.DATABASETYPE DType = 4;</code>
   */
  int getDTypeValue();
  /**
   * <code>.messagedbold.DATABASETYPE DType = 4;</code>
   */
  im.pt.message.proto.DATABASETYPE getDType();
}