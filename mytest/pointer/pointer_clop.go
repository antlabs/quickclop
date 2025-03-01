// Code generated by clop-gen; DO NOT EDIT.

package pointer

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"github.com/BurntSushi/toml"
	"path/filepath"
)


// Code generated by clop-gen; DO NOT EDIT.

// testPointers的Parse方法，用于解析命令行参数
func (c *testPointers) Parse(args []string) error {
	if len(args) == 0 {
		args = os.Args[1:]
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
		
		case arg == "-s" || arg == "--string":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --string")
			}
			
			val := args[i+1]
			c.StringPtr = &val
			
			i++
			
		
		case arg == "-i" || arg == "--int":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --int")
			}
			
			val, err := strconv.Atoi(args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for --int: %v", err)
			}
			c.IntPtr = &val
			
			i++
			
		
		case arg == "-f" || arg == "--float":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --float")
			}
			
			val, err := strconv.ParseFloat(args[i+1], 64)
			if err != nil {
				return fmt.Errorf("invalid value for --float: %v", err)
			}
			c.FloatPtr = &val
			
			i++
			
		
		case arg == "-b" || arg == "--bool":
			
			val := true
			c.BoolPtr = &val
			
		
		
		case !strings.HasPrefix(arg, "-"):
			
			
			
			
			
			
			val := arg
			c.StringArg = &val
			
			
			
		
		default:
			return fmt.Errorf("unknown option: %s", arg)
		}
	}

	return nil
}

func (c *testPointers) Usage() {
	fmt.Println(`USAGE:
  myapp [OPTIONS] [ARGS]

OPTIONS:
    -c, --config    指定配置文件路径 (支持 JSON, YAML, TOML)
    -h, --help      显示帮助信息
    -s, --string    A string pointer
    -i, --int    An integer pointer
    -f, --float    A float pointer
    -b, --bool    A boolean pointer

ARGS:
    stringarg    A string argument as pointer
`)
}

func (c *testPointers) loadFromConfigFile(configFile string) error {
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
