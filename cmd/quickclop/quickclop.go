package main

import (
	"github.com/antlabs/quickclop"
	"os"
)

func main() {
	// 使用命令行参数，如果没有提供则使用当前目录
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	quickclop.Main(path)
}
