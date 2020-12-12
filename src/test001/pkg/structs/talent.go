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
			v.SpeakChinese()
			v.SpeakEnglish()
		}
	}
}
