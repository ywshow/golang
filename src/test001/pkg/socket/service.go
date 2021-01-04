package socket

import (
	"fmt"
	"io"
	"net"
	"strings"
	"test001/pkg/common"
)

//服务端处理流程
//监听端口
//接收客户端tcp链接，建立客户端和服务器端的链接
//创建goroutine,处理该链接的请求（通常客户端会通过链接发送请求包）

func Deal(conn net.Conn) {
	fmt.Printf("处理%v的链接信息\n", conn.RemoteAddr().String())
	defer conn.Close()
	for {
		slice := make([]byte, 1027)
		n, err := conn.Read(slice)
		//err要么出错，要么读到最后一行，
		//return
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完毕")
			} else {
				fmt.Printf("读取%v的信息异常：%v\n", conn.RemoteAddr().String(), err)
			}
			return
		}
		//只读slice[:n]长度到n的数据
		fmt.Printf("读到%v的信息为:%v\n", conn.RemoteAddr().String(), string(slice[:n]))
	}
}

func ListenInfo() {
	listen, err := net.Listen(common.SocketTcp, common.SocketUrl+common.SocketPort)
	if err != nil {
		fmt.Printf("服务器监听异常,%v", err)
	}
	defer listen.Close()
	//不退出监听
	for {
		fmt.Println("等待客户端输入...")
		conn, errAcc := listen.Accept()
		if errAcc != nil {
			fmt.Println("等待输入异常")
		}
		go Deal(conn)
	}
}

//客户端对话
func Conversation() {
	listen, err := net.Listen(common.SocketTcp, common.SocketUrl+common.SocketPort)
	if err != nil {
		fmt.Printf("服务器监听异常,%v", err)
	}
	defer listen.Close()
	//不退出监听
	for {
		fmt.Println("等待客户端输入...")
		conn, errAcc := listen.Accept()
		if errAcc != nil {
			fmt.Println("等待输入异常")
		}
		go ConversationDeal(conn)
	}
}

var addMap = make(map[string]interface{})

func ConversationDeal(conn net.Conn) {
	address := conn.RemoteAddr().String()
	fmt.Printf("处理%v的链接信息\n", address)
	fmt.Println("当前地址：", address)
	//defer conn.Close()
	addMap[address] = address
	fmt.Println("map=", addMap)
	for k, v := range addMap {
		if k != address {
			fmt.Println("v====:", v)
			str := v.(string)
			t := strings.LastIndex(str, ":")
			port := str[t+1 : len(str)]
			ipAddr, _ := net.ResolveTCPAddr(common.SocketTcp, common.SocketUrl+port)
			clientConn, err := net.DialTCP(common.SocketTcp, nil, ipAddr)
			if err != nil {
				fmt.Println("与另一客户端对话失败", err)
			}
			fmt.Println(clientConn.RemoteAddr().String())
			slice := make([]byte, 1024)
			clientConn.Read(slice)
			clientConn.Close()
		}
	}
}
