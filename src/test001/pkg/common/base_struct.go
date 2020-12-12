package common

import "time"

type BaseStruct struct {
	Id         string    `json:"id"`
	CreateDate time.Time `json:"createDate"`
	UpdateDate time.Time `json:"updateDate"`
	//是否有效：默认无效：false
	Invalid bool `json:"invalid"`
}
