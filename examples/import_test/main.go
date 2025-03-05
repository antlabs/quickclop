package main

import (
	"fmt"
	"net"
	"net/url"
)

// :quickclop
type TestImport struct {
	// IP 地址类型
	IP    net.IP  `clop:"--ip" default:"127.0.0.1" usage:"IP地址"`
	IPPtr *net.IP `clop:"--ip-ptr" usage:"IP地址指针"`

	// IPNet 类型
	IPNet    net.IPNet  `clop:"--ipnet" usage:"IP网络"`
	IPNetPtr *net.IPNet `clop:"--ipnet-ptr" usage:"IP网络指针"`

	// URL 类型
	URL    url.URL  `clop:"--url" default:"https://example.com" usage:"URL"`
	URLPtr *url.URL `clop:"--url-ptr" usage:"URL指针"`
}

func main() {
	var test TestImport

	// 这里会调用生成的代码来解析命令行参数
	// 实际运行时需要取消注释
	// quickclop.MustRun(&test)

	fmt.Printf("IP: %v\n", test.IP)
	if test.IPPtr != nil {
		fmt.Printf("IPPtr: %v\n", *test.IPPtr)
	}

	fmt.Printf("IPNet: %v\n", test.IPNet)
	if test.IPNetPtr != nil {
		fmt.Printf("IPNetPtr: %v\n", *test.IPNetPtr)
	}

	fmt.Printf("URL: %v\n", test.URL)
	if test.URLPtr != nil {
		fmt.Printf("URLPtr: %v\n", *test.URLPtr)
	}
}
