package structs

import (
	"fmt"
	"test001/pkg/common"
)

type Student struct {
	Base   common.BaseStruct
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Course string `json:"course"`
}

func (stu Student) SpeakEnglish() {
	fmt.Printf("english:学生%v学的是%v\n", stu.Name, stu.Course)
}

func (stu Student) SpeakChinese() {
	fmt.Printf("Chinese:id:%v,学生%v学的是%v,%v,%v\n", stu.Base.Id, stu.Name, stu.Course, stu.Base.CreateDate, stu.Base.UpdateDate)
}

func (stu Student) Speak() {
	fmt.Printf("ID为%v的学生%v是说%v的人才\n", stu.Base.Id, stu.Name, stu.Course)
}
