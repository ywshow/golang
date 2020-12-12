package structs

import (
	"fmt"
	"test001/pkg/common"
)

type Translator struct {
	Base common.BaseStruct
	Name string `json:"name"`
	//翻译的语言
	Translate string `json:"translate"`
}

func (t Translator) SpeakEnglish() {
	fmt.Printf("englist:翻译者%v翻译的是%v\n", t.Name, t.Translate)
}

func (t Translator) SpeakChinese() {
	fmt.Printf("chinese:翻译者%v翻译的是%v\n", t.Name, t.Translate)
}

func (t Translator) Speak() {
	fmt.Printf("ID为%v的%v是%v翻译的人才\n", t.Base.Id, t.Name, t.Translate)
}

func (t Translator) Salary() {
	fmt.Println("翻译人员的工资为300/天")
}
