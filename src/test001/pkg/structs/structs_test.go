package structs

import (
	"testing"
)

//go test -v：测试所有
//go test -v structs_test.go student.go：测试某个文件的案例
//go test -v -test.run TestSpeakEnglish：测试某个犯法的案例
func TestSpeakEnglish(t *testing.T) {
	stu := Student{
		Name:   "tom",
		Course: "英语",
		Age:    19,
	}
	stu.SpeakEnglish()
}

func TestSpeakEnglishForTrans(t *testing.T) {
	tran := Translator{
		Name:      "jack",
		Translate: "中文",
	}
	tran.SpeakEnglish()
}
