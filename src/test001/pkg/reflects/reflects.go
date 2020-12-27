package reflects

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address string  `json:"address"`
	Score   float64 `json:"score"`
}

func (stu Student) PrintInfo() {
	fmt.Printf("姓名为：%v,年龄为：%v,住址为：%v,成绩为：%v\n", stu.Name, stu.Age, stu.Address, stu.Score)
}

func (stu Student) SumScore(num1, num2 float64) float64 {
	return num1 + num2
}

func (stu Student) InitStudent(name string, age int, address string, score float64) Student {
	return Student{
		Name:    name,
		Age:     age,
		Address: address,
		Score:   score,
	}
}

func TheReflectOfStruct(stu interface{}) {
	value := reflect.ValueOf(stu)
	typeInfo := reflect.TypeOf(stu)
	kind := value.Kind()

	if kind != reflect.Struct {
		panic("参数需为结构体类型")
	}
	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		fmt.Printf("第%v个字段的名称为%v,标签为%v,值为：%v\n", i, field.Name, field.Tag.Get("json"), value.Field(i))
	}

	for i := 0; i < typeInfo.NumMethod(); i++ {
		fmt.Printf("第%v个方法为：%v\n", i, typeInfo.Method(i).Name)
	}

	//函数调用赋值
	//val[0]获取的函数为函数字母ascII码排序
	//所以第0个为InitStudent函数,第一个为PrintInfo函数，第二个为SumScore函数
	var par []reflect.Value
	par = append(par, reflect.ValueOf("jack"))
	par = append(par, reflect.ValueOf(20))
	par = append(par, reflect.ValueOf("shanghai"))
	par = append(par, reflect.ValueOf(90.0))
	val := value.Method(0).Call(par)
	vInter := val[0].Interface()
	s := vInter.(Student)
	s.PrintInfo()

	var scoreSlice []reflect.Value
	scoreSlice = append(scoreSlice, reflect.ValueOf(s.Score))
	scoreSlice = append(scoreSlice, reflect.ValueOf(40.5))
	scoreVal := value.Method(2).Call(scoreSlice)
	fmt.Println("计算分数为：", scoreVal[0].Float())
}

func SReflects01(inter interface{}) (result interface{}) {
	value := reflect.ValueOf(inter)

	fmt.Printf("value的type为：%v,kind为:%v\n", value.Type(), value.Kind())

	return value.Interface()
}
