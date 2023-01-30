package pbcomm

import (
	"encoding/json"
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
