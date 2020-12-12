package structs

import (
	"test001/pkg/inter"
)

//人才
type Talent struct {
}

func (t Talent) TypeOfTalent(speak inter.SpeakInter) {
	speak.Speak()
}
