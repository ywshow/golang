package routine

import (
	"fmt"
	"strconv"
	"time"
)

//素数
func Primes(start int, end int) (slice []int) {
	for i := start; i <= end; i++ {
		if i%2 != 0 {
			slice = append(slice, i)
		}
	}
	return
}

//在主线程（可以理解成进程）中开启一个goroutine，该协程每秒输出一个hello word
//在主线程中也每隔一秒输出Heller go，输出10此后退出程序
//要求主线程和goroutine同时执行
//可以理解为：主方法是主线程，goroutine是协程，一个主线程包含多个协程

func PrintWord() {
	for i := 0; i < 10; i++ {
		fmt.Println("PrintWord():hello Word" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func PrintGolang() {
	for i := 0; i < 10; i++ {
		fmt.Println("PrintGolang():hello golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
