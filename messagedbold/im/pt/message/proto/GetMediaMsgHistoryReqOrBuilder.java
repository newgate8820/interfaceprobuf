// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: hbase.proto

package im.pt.message.proto;

public interface GetMediaMsgHistoryReqOrBuilder extends
    // @@protoc_insertion_point(interface_extends:messagedbold.GetMediaMsgHistoryReq)
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
   * <code>int32 MaxMsgID = 4;</code>
   */
  int getMaxMsgID();

  /**
   * <code>repeated int32 Medias = 5;</code>
   */
  java.util.List<java.lang.Integer> getMediasList();
  /**
   * <code>repeated int32 Medias = 5;</code>
   */
  int getMediasCount();
  /**
   * <code>repeated int32 Medias = 5;</code>
   */
  int getMedias(int index);

  /**
   * <code>int32 Limit = 6;</code>
   */
  int getLimit();

  /**
   * <code>.messagedbold.DATABASETYPE DType = 7;</code>
   */
  int getDTypeValue();
  /**
   * <code>.messagedbold.DATABASETYPE DType = 7;</code>
   */
  im.pt.message.proto.DATABASETYPE getDType();
}
