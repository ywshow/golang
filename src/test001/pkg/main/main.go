package main

import (
	"fmt"
	"test001/pkg/common"
	"test001/pkg/inter"
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

	//接口多态数组
	fmt.Printf("\n>>>接口多态数组>>>\n")
	var slice []inter.SpeakInter
	slice = make([]inter.SpeakInter, 1)
	slice = append(slice, stu)
	slice = append(slice, t)
	fmt.Println("ssssss", slice)
	talent.TypeOfTalentWithArray(slice)
}
