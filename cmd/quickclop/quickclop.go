package main

import (
	"log"
	"os"

	"github.com/antlabs/quickclop"
)

func main() {

	var opts quickclop.Options
	if err := opts.Parse(os.Args[1:]); err != nil {
		log.Fatalf("解析参数失败: %v", err)
	}

	if opts.Help {
		opts.Usage()
		return
	}

	// 如果没有提供任何参数，则使用 Main 函数自动处理当前目录下的所有文件
	if opts.InputFile == "" && opts.StructName == "" && !opts.Completion {
		// 使用当前目录
		path := "."
		log.Printf("未指定参数，将自动处理当前目录下的所有文件")
		quickclop.Main(path)
		return
	}

	// 检查必要参数
	if opts.InputFile == "" {
		log.Fatal("请指定输入文件路径 (-i)")
	}

	if opts.StructName == "" {
		log.Fatal("请指定结构体名称 (-s)")
	}

	// 调用 quickclop 库函数
	if err := quickclop.Generate(opts.InputFile, opts.OutputFile, opts.StructName, opts.ShellType, opts.Completion); err != nil {
		log.Fatalf("生成代码失败: %v", err)
	}
}
