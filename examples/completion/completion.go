package main

import (
	"fmt"
	"os"
)

// Options 定义命令行选项
// :quickclop
type Options struct {
	// 基本选项
	Help    bool `clop:"-h, --help" usage:"显示帮助信息"`
	Version bool `clop:"-v, --version" usage:"显示版本信息"`
	Debug   bool `clop:"-d, --debug" usage:"启用调试模式"`

	// 文件相关选项，使用 completion 标签指定补全类型
	Input  string `clop:"-i, --input" usage:"输入文件路径" completion:"file"`
	Output string `clop:"-o, --output" usage:"输出文件路径" completion:"file"`
	Dir    string `clop:"--dir" usage:"工作目录" completion:"dir"`

	// 自定义补全选项
	Format string `clop:"-f, --format" usage:"输出格式" completion:"custom" completion_values:"json yaml toml xml"`

	// 子命令
	Add    AddCmd    `clop:"subcmd" usage:"添加记录"`
	Remove RemoveCmd `clop:"subcmd" usage:"删除记录"`
}

// AddCmd 添加子命令
type AddCmd struct {
	Name  string `clop:"-n, --name" usage:"记录名称" completion:"custom" completion_values:"user group role"`
	Value string `clop:"-v, --value" usage:"记录值"`
	File  string `clop:"-f, --file" usage:"从文件导入" completion:"file"`
}

// RemoveCmd 删除子命令
type RemoveCmd struct {
	Name  string `clop:"-n, --name" usage:"记录名称" completion:"custom" completion_values:"user group role"`
	All   bool   `clop:"-a, --all" usage:"删除所有记录"`
	Force bool   `clop:"--force" usage:"强制删除，不提示确认"`
}

//go:generate go run ../../quickclop.go -i completion.go -o completion_clop.go -s Options

func main() {
	var opt Options
	if err := opt.Parse(os.Args[1:]); err != nil {
		fmt.Println(err)
		opt.Usage()
		os.Exit(1)
	}

	if opt.Help {
		opt.Usage()
		return
	}

	if opt.Version {
		fmt.Println("v1.0.0")
		return
	}

	fmt.Printf("选项: %+v\n", opt)
}
