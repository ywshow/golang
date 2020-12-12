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

	//接口多态数组，包含类型断言（即判断是否为同类型以方便调用对应结构体有的方法）在TypeOfTalentWithArray中
	fmt.Printf("\n>>>接口多态数组>>>\n")
	var slice []inter.SpeakInter
	slice = make([]inter.SpeakInter, 1)
	slice = append(slice, stu)
	slice = append(slice, t)
	fmt.Println("ssssss", slice)
	talent.TypeOfTalentWithArray(slice)

	fmt.Println("slice", slice)

	//断言扩展
	//var interParam interface{}
	param := make([]interface{}, 0)
	fmt.Println("...", param)
	param1 := true
	var param2 byte = 'h'
	param3 := "你好"
	param4 := 120
	param5 := 4.3
	fmt.Printf("%T\n", param2)
	JudgeType(param1, param2, param3, param4, param5, stu, &stu)
}

func JudgeType(item ...interface{}) {
	//if param,ok:=
	fmt.Println(len(item))
	for _, v := range item {
		//判断v的类型
		switch v.(type) {
		case bool:
			fmt.Printf("v是%T类型，值是%v\n", v, v)
		case byte:
			fmt.Printf("v是%T类型，值是%c\n", v, v)
		case string:
			fmt.Printf("v是%T类型，值是%v\n", v, v)
		case int, int32, int16, int64:
			fmt.Printf("v是%T类型，值是%v\n", v, v)
		case float64:
			fmt.Printf("v是%T类型，值是%v\n", v, v)
		case structs.Student:
			fmt.Printf("v是%T类型，值是%v\n", v, v)
		case *structs.Student:
			fmt.Printf("v是%T类型，值是%v\n", v, v)
		default:
			fmt.Printf("v是%T类型，值是%v\n", v, v)
			fmt.Println("未知类型")
		}
	}

}
