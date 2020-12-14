package main

import "file/code"

func main() {
	path := "C:\\Users\\yw\\Desktop\\678.jpg"
	pathDst := "C:\\Users\\yw\\Desktop\\456.jpg"
	code.WriteFileByPath(path)
	code.ReadFileByPath(path)

	//文件复制
	//pathDst:目标文件，即，把内容复制到此文件
	//path：源文件
	code.CopyFile(pathDst, path)
}
