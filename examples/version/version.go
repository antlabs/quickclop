package main

import (
	"fmt"
)

// :quickclop version:"1.0.0" description:"版本信息示例" build:"2025-03-01"
type Options struct {
	Name    string `clop:"-n,--name" usage:"名称"`
	Age     int    `clop:"-a,--age" usage:"年龄"`
	Version bool   `clop:"-v,--version" usage:"显示版本信息"`
}

func main() {
	opts := &Options{}
	if err := opts.Parse(nil); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", opts.Name, opts.Age)
}
