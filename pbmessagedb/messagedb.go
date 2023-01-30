package messagedb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/siddontang/go-mysql/canal"
	"gitlab.chatserver.im/commonlib/utillib/potatolog"
)

/*
   Author: Dev0026
   创建日期：2017/11/6
   编辑日期：2017/11/6
   修改详细描述
*/

func (d *UserMsg) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *UserMsg) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (d *UserDialog) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *UserDialog) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (d *UserPingedDialog) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *UserPingedDialog) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (d *UserDialogAndMsg) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *UserDialogAndMsg) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (d *UserMsg) GetUserIdFromBinLog(logger potatolog.Logger, e *canal.RowsEvent, rows []interface{}) (int32, error) {
	return d.GetInt32FromBinLog(logger, e, rows, "user_id")
}

func (d *UserMsg) GetMsgIdFromBinLog(logger potatolog.Logger, e *canal.RowsEvent, rows []interface{}) (int32, error) {
	return d.GetInt32FromBinLog(logger, e, rows, "msg_id")
}

func (d *UserMsg) GetInt32FromBinLog(logger potatolog.Logger, e *canal.RowsEvent, rows []interface{}, colName string) (int32, error) {
	var int32Val int32
	idx := e.Table.FindColumn(strings.ToLower(colName))
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			int32Val = val
		} else {
			er := fmt.Errorf("userMsg with_id set val err")
			logger.Error(er)
			return int32Val, er
		}
	} else {
		er := fmt.Errorf("not found with_id")
		logger.Error(er)
		return int32Val, er
	}
	return int32Val, nil
}

func (d *UserMsg) GetUserIdAndMsgIdFromBinLog(logger potatolog.Logger, e *canal.RowsEvent, rows []interface{}) (int32, int32, error) {
	userId, err := d.GetUserIdFromBinLog(logger, e, rows)
	if nil != err {
		return 0, 0, err
	}
	msgId, err := d.GetMsgIdFromBinLog(logger, e, rows)
	if nil != err {
		return 0, 0, err
	}
	return userId, msgId, nil
}

// 返回 withID msgId pts
func (d *UserMsg) StrconvToDeleteReqByBinLog(logger potatolog.Logger, e *canal.RowsEvent, rows []interface{}) (int32, int32, int32, error) {
	var withId, msgId, pts int32
	idx := e.Table.FindColumn("with_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.WithId, withId = val, val
		} else {
			er := fmt.Errorf("userMsg with_id set val err")
			logger.Error(er)
			return withId, msgId, pts, er
		}
	} else {
		er := fmt.Errorf("not found with_id")
		logger.Error(er)
		return withId, msgId, pts, er
	}

	idx = e.Table.FindColumn("msg_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.MsgId, msgId = val, val
		} else {
			er := fmt.Errorf("userMsg msg_id set val err")
			logger.Error(er)
			return withId, msgId, pts, er
		}
	} else {
		er := fmt.Errorf("not found msg_id")
		logger.Error(er)
		return withId, msgId, pts, er
	}

	idx = e.Table.FindColumn("pts")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.Pts, pts = val, val
		} else {
			er := fmt.Errorf("userMsg pts set val err")
			logger.Error(er)
			return withId, msgId, pts, er
		}
	} else {
		er := fmt.Errorf("not found pts")
		logger.Error(er)
		return withId, msgId, pts, er
	}

	return withId, msgId, pts, nil
}

func (d *UserMsg) StrconvByBinLog(logger potatolog.Logger, e *canal.RowsEvent, rows []interface{}) {
	idx := e.Table.FindColumn("id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int64); ok {
			d.Id = val
		} else {
			er := fmt.Errorf("userMsg id set val err")
			logger.Error(er)
		}
	}

	idx = e.Table.FindColumn("user_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.UserId = val
		} else {
			er := fmt.Errorf("userMsg user_id set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("with_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.WithId = val
		} else {
			er := fmt.Errorf("userMsg with_id set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("with_id_type")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.WithIdType = val
		} else {
			er := fmt.Errorf("userMsg with_id_type set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("message_type")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.MessageType = val
		} else {
			er := fmt.Errorf("userMsg message_type set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("flags")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.Flags = val
		} else {
			er := fmt.Errorf("userMsg flags set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("out")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int8); ok {
			if int8(1) == val {
				d.Out = true
			}
		} else {
			er := fmt.Errorf("userMsg out set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("mentioned")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int8); ok {
			if int8(1) == val {
				d.Mentioned = true
			}
		} else {
			er := fmt.Errorf("userMsg mentioned set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("media_unread")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int8); ok {
			if int8(1) == val {
				d.MediaUnread = true
			}
		} else {
			er := fmt.Errorf("userMsg media_unread set val err")
			logger.Error(er, reflect.TypeOf(rows[idx]).Kind())
		}
	}
	idx = e.Table.FindColumn("silent")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int8); ok {
			if int8(1) == val {
				d.Silent = true
			}
		} else {
			er := fmt.Errorf("userMsg silent set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("post")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int8); ok {
			if int8(1) == val {
				d.Post = true
			}
		} else {
			er := fmt.Errorf("userMsg post set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("msg_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.MsgId = val
		} else {
			er := fmt.Errorf("userMsg msg_id set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("to_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.ToId = val
		} else {
			er := fmt.Errorf("userMsg to_id set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("pts")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.Pts = val
		} else {
			er := fmt.Errorf("userMsg pts set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("fwd_from")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].([]byte); ok {
			d.FwdFrom = val
		} else {
			er := fmt.Errorf("userMsg fwd_from set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("via_bot_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.ViaBotId = val
		} else {
			er := fmt.Errorf("userMsg via_bot_id set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("reply_to_msg_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.ReplyToMsgId = val
		} else {
			er := fmt.Errorf("userMsg reply_to_msg_id set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("date")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.Date = val
		} else {
			er := fmt.Errorf("userMsg date set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("message")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(string); ok {
			d.Message = val
		} else {
			er := fmt.Errorf("userMsg message set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("media")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int64); ok {
			d.Media = val
		} else {
			er := fmt.Errorf("userMsg media set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("media_data")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].([]byte); ok {
			d.MediaData = val
		} else {
			er := fmt.Errorf("userMsg media_data set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("reply_markup")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].([]byte); ok {
			d.ReplyMarkup = val
		} else {
			er := fmt.Errorf("userMsg reply_markup set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("entities")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].([]uint8); ok {
			vv := reflect.ValueOf(val)
			//var entity = make([][]byte, 0)
			err := json.Unmarshal(vv.Bytes(), &d.Entities)
			if nil != err {
				er := fmt.Errorf("json unmarshal usermsg entities err: %v", err)
				logger.Error(er)
			}
		} else {
			er := fmt.Errorf("userMsg entities set val err")
			logger.Error(er, reflect.TypeOf(rows[idx]).Elem().Kind())
		}
	}
	idx = e.Table.FindColumn("from_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.FromId = val
		} else {
			er := fmt.Errorf("userMsg from_id set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("views")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.Views = val
		} else {
			er := fmt.Errorf("userMsg views set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("edit_date")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.EditDate = val
		} else if val, ok := rows[idx].(int64); ok {
			d.EditDate = int32(val)
		} else {
			er := fmt.Errorf("userMsg edit_date set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("action")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].([]byte); ok {
			d.Action = val
		} else {
			er := fmt.Errorf("userMsg action set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("random_id")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int64); ok {
			d.RandomId = val
		} else {
			er := fmt.Errorf("userMsg random_id set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("uuid")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int64); ok {
			d.Uuid = val
		} else {
			er := fmt.Errorf("userMsg uuid set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("insert_date")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.InsertDate = val
		} else if val, ok := rows[idx].(int64); ok {
			d.InsertDate = int32(val)
		} else {
			er := fmt.Errorf("userMsg insert_date set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("update_date")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.UpdateDate = val
		} else if val, ok := rows[idx].(int64); ok {
			d.UpdateDate = int32(val)
		} else {
			er := fmt.Errorf("userMsg update_date set val err")
			logger.Error(er)
		}
	}
	idx = e.Table.FindColumn("encry")
	if idx >= 0 && rows[idx] != nil {
		if val, ok := rows[idx].(int32); ok {
			d.Encry = val
		} else {
			er := fmt.Errorf("userMsg encry set val err")
			logger.Error(er)
		}
	}
}
