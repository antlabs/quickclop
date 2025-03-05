package main

import (
	"fmt"
	"net"
	"os"
)

// :quickclop
// Config 演示 IP 地址类型的使用
type Config struct {
	// IP 地址类型
	ServerIP net.IP `clop:"--server-ip" usage:"服务器 IP 地址" default:"127.0.0.1" env:"SERVER_IP"`
	ClientIP *net.IP `clop:"--client-ip" usage:"客户端 IP 地址"`

	// IP 网络类型
	ServerNetwork net.IPNet `clop:"--server-network" usage:"服务器网络 CIDR" default:"192.168.1.0/24" env:"SERVER_NETWORK"`
	ClientNetwork *net.IPNet `clop:"--client-network" usage:"客户端网络 CIDR"`

	// 帮助信息
	Help bool `clop:"-h,--help" usage:"显示帮助信息"`
}

func main() {
	var config Config
	if err := config.Parse(os.Args[1:]); err != nil {
		fmt.Printf("解析参数错误: %v\n", err)
		return
	}

	if config.Help {
		config.Usage()
		return
	}

	// 显示 IP 地址信息
	fmt.Println("IP 地址信息:")
	fmt.Printf("  服务器 IP: %v\n", config.ServerIP)
	
	if config.ClientIP != nil {
		fmt.Printf("  客户端 IP: %v\n", *config.ClientIP)
	} else {
		fmt.Println("  客户端 IP: 未设置")
	}

	// 显示网络信息
	fmt.Println("\n网络信息:")
	fmt.Printf("  服务器网络: %v (IP: %v, 掩码: %v)\n", 
		config.ServerNetwork.String(),
		config.ServerNetwork.IP,
		config.ServerNetwork.Mask)
	
	if config.ClientNetwork != nil {
		fmt.Printf("  客户端网络: %v (IP: %v, 掩码: %v)\n", 
			config.ClientNetwork.String(),
			config.ClientNetwork.IP,
			config.ClientNetwork.Mask)
	} else {
		fmt.Println("  客户端网络: 未设置")
	}
}
