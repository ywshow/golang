package structs

import (
	"fmt"
	"test001/pkg/inter"
)

//人才
type Talent struct {
}

func (t Talent) TypeOfTalent(speak inter.SpeakInter) {
	speak.Speak()
}

func (t Talent) TypeOfTalentWithArray(slice []inter.SpeakInter) {
	fmt.Printf("len=%v\n", len(slice))
	for _, v := range slice {
		if v != nil {
			//类型断言，判断多态结构体当前循环的v是否是Translator结构体，是，则调用方法
			if v, ok := v.(Translator); ok {
				v.Salary()
			}
			v.SpeakChinese()
			v.SpeakEnglish()
		}
	}
}
