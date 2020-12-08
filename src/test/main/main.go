package main

import (
	"fmt"
	"sort"
	"test/slice"
)

func main() {
	fmt.Println(">>>")
	//结构体切片排序
	slice := slice.RandSliceStudent(10)
	sort.Sort(slice)
	for _, v := range slice {
		fmt.Println(v)
	}
}
