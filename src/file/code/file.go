package code

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFileByPath(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件异常：", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, readErr := reader.ReadString('\n')
		fmt.Print(str)
		if readErr == io.EOF {
			break
		}
	}
	fmt.Println()
}

func WriteFileByPath(path string) {
	//os.O_WRONLY|os.O_CREATE只读没有则创建
	//os.O_APPEND：写的内容追加到文件最后
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("打开文件异常：", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("什么东西")
	writer.Flush()
}

//判断文件或者文件夹是否存在
//如果返回的错误为nil，说明文件或文件夹存在
//如果返回的错误类型用os.IsNotExist判断为true,说明文件或者文件夹不存在
//如果返回的错误为其他类型，则不确定是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func openOrCreateFile(path string, flag int, perm os.FileMode) (*os.File, error) {
	file, errFile := os.OpenFile(path, flag, perm)
	if errFile != nil {
		return file, errFile
	}
	return file, nil
}

func CopyFile(dst string, src string) {
	flagDst, err := PathExist(dst)
	var dstFile *os.File
	var srcFile *os.File
	//目标文件检测
	if !flagDst {
		if err == nil {
			dstFileTmp, errFile := openOrCreateFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
			defer dstFileTmp.Close()
			if errFile != nil {
				fmt.Println("其他错误")
				return
			}
			dstFile = dstFileTmp
		} else {
			fmt.Println("其他错误")
			return
		}
	} else {
		dstFileTmp, errFile := openOrCreateFile(dst, os.O_WRONLY, 0666)
		defer dstFileTmp.Close()
		if errFile != nil {
			fmt.Println("其他错误")
			return
		}
		dstFile = dstFileTmp
	}

	//源文件检测
	flagSrc, errSrc := PathExist(src)
	if !flagSrc {
		if errSrc == nil {
			fmt.Println("文件不存在")
		} else {
			fmt.Println("其他错误")
		}
		return
	}
	srcFile, errSrcFile := os.OpenFile(src, os.O_RDONLY, 0666)
	defer srcFile.Close()
	if errSrcFile != nil {
		return
	}

	bufDst := bufio.NewWriter(dstFile)
	bufSrc := bufio.NewReader(srcFile)
	//for {
	//	str, err := bufSrc.ReadString('\n')
	//	fmt.Println("str=", str)
	//	if err == io.EOF {
	//		break
	//	}
	//}
	written, errCopy := io.Copy(bufDst, bufSrc)
	if errCopy != nil {
		fmt.Println("errCopy:", errCopy)
	}
	fmt.Println("written:", written)
}
