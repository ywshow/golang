package routine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ChannelStudy() {
	channel := make(chan int, 5)
	fmt.Println("未赋值前channel的长度：", len(channel))
	channel <- 10
	num := 30
	channel <- num
	channel <- 100
	channel <- 234
	channel <- 99
	fmt.Println("赋值后channel的长度：", len(channel))

	//取出channel，只能取出同等大小的次数（5）,取第六次报deadlock错误
	//channel为先进先出（FIFO）,取出后，长度减少，容量（cap）也会变（用for range循环），可继续存到channel
	//channel为动态，即取出一次，长度减一
	//关闭通道，关闭后只能读不能写入
	close(channel)
	for v := range channel {
		fmt.Printf("channel中的值为%v\n", v)
	}
}

type Person struct {
	Name    string
	Age     int
	Address string
}

func ChannelStudy01() {
	cha := make(chan Person, 10)
	for i := 0; i < 10; i++ {
		ps := &Person{
			Name:    "tome~" + strconv.Itoa(i),
			Age:     rand.Intn(100),
			Address: "北京~" + strconv.Itoa(i),
		}
		cha <- *ps
	}

	//在遍历时，，如果
	close(cha)
	for v := range cha {
		p := v
		str, _ := json.Marshal(p)
		fmt.Println(string(str))
	}
	//关闭通道，关闭后只能读不能写入
	//close(cha)
	fmt.Println("channel长度=", len(cha))
}

func WriteData(cha chan int, num int) {
	for i := 1; i <= num; i++ {
		cha <- i
		fmt.Println("写入数据：", i)
	}
	//执行完毕，关闭channel
	close(cha)
}

func ReadData(cha chan int, exitChan chan bool) {
	for {
		//当读不到东西，ok为false
		v, ok := <-cha
		if !ok {
			break
		}
		fmt.Println("读出的数据：", v)
	}
	//执行完毕，关闭channel
	exitChan <- true
	close(exitChan)
	//for v := range cha {
	//	fmt.Println("写入的数据有：",v)
	//}
}

func DemoOneIntiChannel(num int, numChan chan int) {
	for i := 1; i <= num; i++ {
		numChan <- i
	}
	close(numChan)
}

func DemoOneLogic(numChan chan int, exitChan chan bool, resChan chan string, num int) {

	for v := range numChan {
		sum := 0
		for i := 1; i <= v; i++ {
			sum += i
		}
		resChan <- strconv.Itoa(sum)

		if v == num {
			fmt.Println(">>>>>>", v)
			exitChan <- true
			close(resChan)
			close(exitChan)
		}

	}
}

type SliceStr []int

func (slice SliceStr) Len() int {
	return len(slice)
}

func (slice SliceStr) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice SliceStr) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func WriteDataToFile(randNum int, path string) {
	var slice []int
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= randNum; i++ {
		slice = append(slice, rand.Intn(1000))
	}
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件异常:", err)
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	bytes, _ := json.Marshal(slice)
	writer.Write(bytes)
	writer.Flush()
}

func Sort(path string, dstPath string) {
	fmt.Println("path=", path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件异常：", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	dataStr := ""
	for {
		str, err := reader.ReadString('\n')
		dataStr = str
		if err == io.EOF {
			break
		}
	}

	//排序
	data := []byte(dataStr)
	var slice SliceStr
	json.Unmarshal(data, &slice)
	fmt.Println("slice==", slice)
	sort.Sort(slice)

	fmt.Println("排序后的结果为：", slice)
	fileDst, errDst := os.OpenFile(dstPath, os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if errDst != nil {
		fmt.Println("打开文件异常")
	}
	defer fileDst.Close()
	writer := bufio.NewWriter(fileDst)
	sliceByte, _ := json.Marshal(slice)
	writer.Write(sliceByte)
	writer.Flush()
}

func WriteDataToFile01(randNum int, path string, pathChan chan string, index int, writeDoneChan chan bool) {
	var slice []int
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= randNum; i++ {
		slice = append(slice, rand.Intn(1000))
	}
	ind := strings.LastIndex(path, ".")
	if index == -1 {
		fmt.Println("文件路径不对")
		return
	}
	path = path[:ind] + strconv.Itoa(index) + path[ind:len(path)]
	fmt.Println(path)
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件异常:", err)
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	bytes, _ := json.Marshal(slice)
	writer.Write(bytes)
	writer.Flush()
	pathChan <- path
	writeDoneChan <- true
}

func Sort01(dstPath string, path string, exitChan chan bool, index int) {
	writeSort(dstPath, path, index)
	exitChan <- true
}

func writeSort(dstPath string, path string, index int) {
	fmt.Println("writeSort path=", path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件异常：", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	dataStr := ""
	for {
		str, err := reader.ReadString('\n')
		dataStr = str
		if err == io.EOF {
			break
		}
	}

	//排序
	data := []byte(dataStr)
	var slice SliceStr
	json.Unmarshal(data, &slice)
	//fmt.Println("slice==", slice)
	sort.Sort(slice)

	//fmt.Println("排序后的结果为：", slice)
	ind := strings.LastIndex(dstPath, ".")
	dstPath = dstPath[:ind] + strconv.Itoa(index) + dstPath[ind:len(dstPath)]
	fileDst, errDst := os.OpenFile(dstPath, os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if errDst != nil {
		fmt.Println("打开文件异常")
	}
	defer fileDst.Close()
	writer := bufio.NewWriter(fileDst)
	sliceByte, _ := json.Marshal(slice)
	writer.Write(sliceByte)
	writer.Flush()
}
