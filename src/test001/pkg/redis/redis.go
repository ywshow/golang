package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"test001/pkg/common"
)

func Deal(cmd string, args ...interface{}) {
	//连接redis
	conn, err := redis.Dial(common.SocketTcp, common.SocketUrl+common.RedisPort)
	if err != nil {
		fmt.Println("redis链接一号仓", err)
	}
	defer conn.Close()

	fmt.Println("链接成功：", conn)
	//_,er:=redis.Do(cmd,args)
	//if er != nil {
	//	fmt.Println("redis设置值异常：",er)
	//}

	sResult, err := conn.Do("set", "name", "ywshow0000001")
	if err != nil {
		fmt.Println("redis设置值异常：", err)
	}
	sStr, _ := redis.String(sResult, err)
	fmt.Println(sStr)

	result, gErr := conn.Do("get", "name")
	if gErr != nil {
		fmt.Println("redis获取值异常:", gErr)
	}
	str, _ := redis.String(result, gErr)
	fmt.Println("操作成功，获取到的值为：", str)
}
