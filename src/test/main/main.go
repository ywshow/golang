package main

import (
	"fmt"
	"sort"
	"test/slice"
)

func main() {
	fmt.Println(">>>")
	//结构体切片排序
	stu := slice.RandSliceStudent(10)
	sort.Sort(stu)
	for _, v := range stu {
		fmt.Println(v)
	}
}
