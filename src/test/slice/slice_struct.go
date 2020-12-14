package slice

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name string
	Age  int
}

type StudentSlice []Student

func (slice StudentSlice) Len() int {
	return len(slice)
}

func (slice StudentSlice) Less(i, j int) bool {
	return slice[i].Age < slice[j].Age
}

func (slice StudentSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func RandSliceStudent(num int) (slice StudentSlice) {
	for i := 0; i < num; i++ {
		stu := Student{
			Name: fmt.Sprintf("姓名%v", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		slice = append(slice, stu)
	}
	return
}
