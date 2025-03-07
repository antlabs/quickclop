// Code generated by clop-gen; DO NOT EDIT.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Options的Parse方法，用于解析命令行参数
func (c *Options) Parse(args []string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}

	// 从环境变量中读取值

	if envVal := os.Getenv("UINT8_VALUE"); envVal != "" {

	}

	if envVal := os.Getenv("UINT8_PTR_VALUE"); envVal != "" {

	}

	if envVal := os.Getenv("UINT8_SLICE_VALUE"); envVal != "" {

		for _, s := range strings.Split(envVal, ",") {
			if val, err := strconv.ParseUint(s, 10, 8); err == nil {
				c.Uint8Slice = append(c.Uint8Slice, uint8(val))
			}
		}

	}

	if envVal := os.Getenv("UINT8_SLICE_PTR_VALUE"); envVal != "" {

	}

	// 检查是否指定了配置文件
	var configFile string
	for i := 0; i < len(args); i++ {
		if args[i] == "--config" || args[i] == "-c" {
			if i+1 < len(args) {
				configFile = args[i+1]
				// 移除 --config 和它的值，避免后续解析错误
				if i+2 < len(args) {
					args = append(args[:i], args[i+2:]...)
				} else {
					args = args[:i]
				}
				i--
				break
			}
		}
	}

	// 从配置文件加载选项
	if configFile != "" {
		if err := c.loadFromConfigFile(configFile); err != nil {
			return fmt.Errorf("加载配置文件失败: %w", err)
		}
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--help" || arg == "-h":
			c.Usage()
			os.Exit(0)

		case arg == "-u" || arg == "--uint8":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -u or --uint8")
			}

			val, err := strconv.ParseUint(args[i+1], 10, 8)
			if err != nil {
				return fmt.Errorf("invalid value for -u or --uint8: %v", err)
			}
			c.Uint8 = uint8(val)

			i++

		case arg == "--uint8-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --uint8-ptr")
			}

			val, err := strconv.ParseUint(args[i+1], 10, 8)
			if err != nil {
				return fmt.Errorf("invalid value for --uint8-ptr: %v", err)
			}
			val8 := uint8(val)
			c.Uint8Ptr = &val8

			i++

		case arg == "--uint8-slice":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --uint8-slice")
			}

			for _, s := range strings.Split(args[i+1], ",") {
				if val, err := strconv.ParseUint(s, 10, 8); err == nil {
					c.Uint8Slice = append(c.Uint8Slice, uint8(val))
				}
			}

			i++

		case arg == "--uint8-slice-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --uint8-slice-ptr")
			}

			i++

		case arg == "-h" || arg == "--help":

			c.Help = true

		default:
			return fmt.Errorf("unknown option: %s", arg)
		}
	}

	return nil
}

// UsageTmpl 是用于生成 Usage 函数的模板
func (c *Options) Usage() {
	fmt.Println(`USAGE:
  myapp [OPTIONS]

OPTIONS:
    -c, --config    指定配置文件路径 (支持 JSON, YAML, TOML)
    -h, --help      显示帮助信息
    -u, --uint8    A uint8 value [env: UINT8_VALUE]
    --uint8-ptr    A uint8 pointer value [env: UINT8_PTR_VALUE]
    --uint8-slice    A slice of uint8 values [env: UINT8_SLICE_VALUE]
    --uint8-slice-ptr    A pointer to a slice of uint8 values [env: UINT8_SLICE_PTR_VALUE]
    -h, --help    Show help information
`)
}

func (c *Options) loadFromConfigFile(configFile string) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 根据文件扩展名选择解析方式
	ext := filepath.Ext(configFile)
	switch strings.ToLower(ext) {
	case ".json":
		return json.Unmarshal(data, c)
	case ".yaml", ".yml":
		return yaml.Unmarshal(data, c)
	case ".toml":
		_, err := toml.Decode(string(data), c)
		return err
	default:
		return fmt.Errorf("不支持的配置文件格式: %s", ext)
	}
}
