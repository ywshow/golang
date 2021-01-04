package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"test001/pkg/common"
)

func Input01() {
	conn, err := net.Dial(common.SocketTcp, common.SocketUrl+common.SocketPort)
	if err != nil {
		fmt.Println("客户端获取链接异常")
		return
	}
	//defer conn.Close()
	//os.Stdin:获取从终端输入的信息

	for {
		fmt.Printf("请输入信息：\n")
		reader := bufio.NewReader(os.Stdin)
		str, strErr := reader.ReadString('\n')
		if strErr != nil {
			fmt.Println("从终端读取信息异常")
		}
		//str:="你好啊"
		fmt.Println("读取到的输入信息为：", str)

		if strings.Trim(str, " \r\n") == "quit" {
			return
		}
		n, conErr := conn.Write([]byte(str))
		if conErr != nil {
			fmt.Println("客户端链接服务器异常", conErr)
		}
		fmt.Printf("客户端写入%v个字节到服务器\n", n)
	}
}
