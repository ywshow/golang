package reflects

import (
	"fmt"
	"testing"
)

func TestReflects01(t *testing.T) {
	num := 10.5
	inter := SReflects01(num)
	f := inter.(float64)

	fmt.Println("f>>>>", f)
}

//使用反射来便利结构体字段，调用结构体方法，获取结构体标签
func TestTheReflectOfStruct(t *testing.T) {
	stu := Student{
		Name:    "tom",
		Age:     30,
		Address: "beijing",
		Score:   80,
	}
	TheReflectOfStruct(stu)
	fmt.Println(stu)
}
