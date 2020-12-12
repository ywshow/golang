package main

import (
	"test001/pkg/common"
	"test001/pkg/structs"
	"time"
)

func main() {

	talent := structs.Talent{}

	stu := structs.Student{
		Base: common.BaseStruct{
			Id:         "0000001",
			CreateDate: time.Now(),
			UpdateDate: time.Now(),
			Invalid:    true,
		},
		Name:   "小王",
		Age:    20,
		Course: "英语",
	}

	stu.SpeakEnglish()

	talent.TypeOfTalent(stu)

	t := structs.Translator{
		Name:      "张三",
		Translate: "中文",
	}
	t.SpeakChinese()

	talent.TypeOfTalent(t)
}
