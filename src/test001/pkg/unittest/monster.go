package unittest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store(path string) (string, error) {
	flage, err := pathExist(path)
	var dstFile *os.File
	if !flage {
		if err == nil {
			file, errFile := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
			defer file.Close()
			if errFile != nil {
				return "", errFile
			}
			dstFile = file
		} else {
			return "", err
		}
	} else {
		file, errFile := os.OpenFile(path, os.O_APPEND, 0666)
		defer file.Close()
		if errFile != nil {
			return "", errFile
		}
		dstFile = file
	}

	write := bufio.NewWriter(dstFile)
	str, e := json.Marshal(&m)
	if e != nil {
		return "", e
	}
	count, errWrite := write.Write(str)
	if errWrite != nil {
		return "", errWrite
	}
	write.Flush()
	fmt.Println("总共写入：", count)
	return string(str), nil
}

func (m *Monster) Restore(path string) error {
	if flag, errPath := pathExist(path); !flag {
		return errPath
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	var result string
	for {
		res, errReader := reader.ReadString('\n')
		if errReader != nil {
			return errReader
		}
		result += res
	}
	json.Unmarshal([]byte(result), &m)
	return nil
}

//判断文件或者文件夹是否存在
//如果返回的错误为nil，说明文件或文件夹存在
//如果返回的错误类型用os.IsNotExist判断为true,说明文件或者文件夹不存在
//如果返回的错误为其他类型，则不确定是否存在
func pathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
