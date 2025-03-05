package main

import (
	"fmt"
)

// Config 演示各种切片类型的支持
// :quickclop
type Config struct {
	// 整数类型切片
	IntSlice   []int   `clop:"--ints" usage:"整数切片" default:"1,2,3"`
	Int8Slice  []int8  `clop:"--int8s" usage:"int8切片" default:"1,2,3"`
	Int16Slice []int16 `clop:"--int16s" usage:"int16切片" default:"1,2,3"`
	Int32Slice []int32 `clop:"--int32s" usage:"int32切片" default:"1,2,3"`
	Int64Slice []int64 `clop:"--int64s" usage:"int64切片" default:"1,2,3"`

	// 无符号整数类型切片
	UintSlice   []uint   `clop:"--uints" usage:"uint切片" default:"1,2,3"`
	Uint8Slice  []uint8  `clop:"--uint8s" usage:"uint8切片" default:"1,2,3"`
	Uint16Slice []uint16 `clop:"--uint16s" usage:"uint16切片" default:"1,2,3"`
	Uint32Slice []uint32 `clop:"--uint32s" usage:"uint32切片" default:"1,2,3"`
	Uint64Slice []uint64 `clop:"--uint64s" usage:"uint64切片" default:"1,2,3"`

	// 浮点数类型切片
	Float32Slice []float32 `clop:"--float32s" usage:"float32切片" default:"1.1,2.2,3.3"`
	Float64Slice []float64 `clop:"--float64s" usage:"float64切片" default:"1.1,2.2,3.3"`

	// 布尔类型切片
	BoolSlice []bool `clop:"--bools" usage:"布尔切片" default:"true,false,true"`

	// 字节和字符类型切片
	ByteSlice []byte `clop:"--bytes" usage:"字节切片" default:"abc"`
	RuneSlice []rune `clop:"--runes" usage:"字符切片" default:"你好世界"`

	// 字符串切片 (已支持)
	StringSlice []string `clop:"--strings" usage:"字符串切片" default:"a,b,c"`
}

func main() {
	var config Config

	// 解析命令行参数
	if err := config.Parse(nil); err != nil {
		fmt.Printf("解析参数错误: %v\n", err)
		return
	}

	// 打印解析结果
	fmt.Println("整数类型切片:")
	fmt.Printf("  []int:    %v\n", config.IntSlice)
	fmt.Printf("  []int8:   %v\n", config.Int8Slice)
	fmt.Printf("  []int16:  %v\n", config.Int16Slice)
	fmt.Printf("  []int32:  %v\n", config.Int32Slice)
	fmt.Printf("  []int64:  %v\n", config.Int64Slice)

	fmt.Println("无符号整数类型切片:")
	fmt.Printf("  []uint:   %v\n", config.UintSlice)
	fmt.Printf("  []uint8:  %v\n", config.Uint8Slice)
	fmt.Printf("  []uint16: %v\n", config.Uint16Slice)
	fmt.Printf("  []uint32: %v\n", config.Uint32Slice)
	fmt.Printf("  []uint64: %v\n", config.Uint64Slice)

	fmt.Println("浮点数类型切片:")
	fmt.Printf("  []float32: %v\n", config.Float32Slice)
	fmt.Printf("  []float64: %v\n", config.Float64Slice)

	fmt.Println("布尔类型切片:")
	fmt.Printf("  []bool:    %v\n", config.BoolSlice)

	fmt.Println("字节和字符类型切片:")
	fmt.Printf("  []byte:    %s\n", string(config.ByteSlice))
	fmt.Printf("  []rune:    %s\n", string(config.RuneSlice))

	fmt.Println("字符串切片:")
	fmt.Printf("  []string:  %v\n", config.StringSlice)
}
