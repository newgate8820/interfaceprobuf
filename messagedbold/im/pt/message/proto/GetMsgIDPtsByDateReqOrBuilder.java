// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: hbase.proto

package im.pt.message.proto;

public interface GetMsgIDPtsByDateReqOrBuilder extends
    // @@protoc_insertion_point(interface_extends:messagedbold.GetMsgIDPtsByDateReq)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int32 UserID = 1;</code>
   */
  int getUserID();

  /**
   * <code>int32 Date = 2;</code>
   */
  int getDate();

  /**
   * <code>.messagedbold.DATABASETYPE DType = 3;</code>
   */
  int getDTypeValue();
  /**
   * <code>.messagedbold.DATABASETYPE DType = 3;</code>
   */
  im.pt.message.proto.DATABASETYPE getDType();
}
