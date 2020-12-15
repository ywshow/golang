package routine

import (
	"fmt"
	"runtime"
	"testing"
)

//查出1到20000的所有素数
func TestPrimes(t *testing.T) {

	slice := Primes(1, 20000)
	for _, v := range slice {
		t.Logf("1到20000的素数有：%v", v)
	}
	t.Logf("总共有%v个", len(slice))
}

//go routine
func TestPrintWordAngGolang(t *testing.T) {
	//开启协程，go：产生支线，主线可以继续往下走PrintGolang(),不然自能等PrintWord()执行完才走PrintGolang()
	go PrintWord()
	//主线程
	PrintGolang()
}

//go1.8后默认让程序运行在多个核上，无需再设置
func TestCpu(t *testing.T) {
	num := runtime.NumCPU()
	//设置可同时执行的最大cpu数
	//runtime.GOMAXPROCS(num-1)
	fmt.Println("逻辑cpu个数：", num)
}

func TestChannelStudy(t *testing.T) {
	ChannelStudy()
}

func TestChannelStudy01(t *testing.T) {
	ChannelStudy01()
}

//1)开启一个WriteData协程，向管道cha中写入数据
//2)开启一个ReadData协程，从管道cha中读取WriteData写入的数据
//3)主线程需要等待WriteData和ReadData协程都完成工作才能推出
//数据穿插显示即可
func TestWriteAndReadData(t *testing.T) {
	cha := make(chan int, 500)
	exitChan := make(chan bool, 1)
	num := 500
	go WriteData(cha, num)
	go ReadData(cha, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

//案例1：
//启动一个协程，将1-2000的数放入到一个channel中，
//启动8个协程，从channel读取数，比如（n）,并计算1+...+n的值，并存到另一个channel中
//遍历第二个channel，显示结果，如channel[1]=1...channel[10]=55...
func TestDemoOne(t *testing.T) {
	numChan := make(chan int, 5000)
	exitChan := make(chan bool, 1)
	resChan := make(chan string, 5000)
	num := 2000
	go DemoOneIntiChannel(num, numChan)
	for i := 0; i < 8; i++ {
		go DemoOneLogic(numChan, exitChan, resChan, num)
	}

	length := len(resChan)
	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println("计算结果值为：", v)
	}
	fmt.Println("长度为：", length)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

//案例2：
//开一个协程writeDataToFile,随机生成1000个数，存放到文件中
//当writeDataToFile完成写1000个数据到文件后，让sort协程从文件中读取，并完才排序，重新写入另一个文件中
//扩展：
//开10个协程writeDataToFile，每个协程随机生成1000个数据，存放到是个文件中
//当10个文件都生成了，让10个sort协程从10个文件读取，并完才排序，重新写入到10个结果文件

func TestWriteDataToFileAndSort(t *testing.T) {
	path := "C:\\Users\\yw\\Desktop\\write.txt"
	pathTo := "C:\\Users\\yw\\Desktop\\sort.txt"
	num := 1000
	WriteDataToFile(num, path)
	Sort(path, pathTo)
}

//扩展：
func TestWriteDataToFileAndSort01(t *testing.T) {
	path := "C:\\Users\\yw\\Desktop\\write.txt"
	pathTo := "C:\\Users\\yw\\Desktop\\sort.txt"
	num := 10
	//因为有是个协程写入（WriteDataToFile01），没法确定哪个协程最后结束，所有没法在Sort01方法里关闭channel,在WriteDoneChan中关闭channel
	writeDoneChan := make(chan bool, 4)
	//因为有10个协程排序，没法确定哪个协程最后结束，所有没法在Sort01方法里关闭channel,在exitChan中关闭channel
	exitChan := make(chan bool, 10)
	//存放路径的channel
	pathChan := make(chan string, num)
	for i := 1; i <= 10; i++ {
		go WriteDataToFile01(num, path, pathChan, i, writeDoneChan)
	}

	//写入完毕，关闭pathChan
	go func() {
		//直接抛出即可，当Sort01有协程没结束则等待结束抛出数据（true/false）
		for i := 1; i < 10; i++ {
			<-writeDoneChan
		}
		close(pathChan)
	}()
	for i := 1; i <= 10; i++ {
		str := <-pathChan
		go Sort01(pathTo, str, exitChan, i)
	}

	go func() {
		//直接抛出即可，当Sort01有协程没结束则等待结束抛出数据（true/false）
		for i := 1; i <= 10; i++ {
			<-exitChan
		}
		//关闭channel
		//因为Sort01中没产生新channel，所有这里不用关闭
	}()

}
