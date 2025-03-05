
// Code2 generated by clop-gen; DO NOT EDIT.

package main
// 包含子命令的结构体

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"github.com/BurntSushi/toml"
	"path/filepath"
)

// Options的Parse方法，用于解析命令行参数
func (c *Options) Parse(args []string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}

	
	// 检查是否有子命令
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		subcommand := args[0]
		subcommandArgs := args[1:]

		switch subcommand {
		
		case "Add":
			return c.Add.Parse(subcommandArgs)
		
		case "Remove":
			return c.Remove.Parse(subcommandArgs)
		
		default:
			return fmt.Errorf("未知子命令: %s", subcommand)
		}
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
		
		case arg == "-h, --help":
			
			c.Help = true
			
		
		case arg == "-v, --version":
			
			c.Version = true
			
		
		case arg == "-d, --debug":
			
			c.Debug = true
			
		
		case arg == "-i, --input":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -i, --input")
			}
			
			c.Input = args[i+1]
			
			i++
			
		
		case arg == "-o, --output":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -o, --output")
			}
			
			c.Output = args[i+1]
			
			i++
			
		
		case arg == "--dir":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --dir")
			}
			
			c.Dir = args[i+1]
			
			i++
			
		
		case arg == "-f, --format":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -f, --format")
			}
			
			c.Format = args[i+1]
			
			i++
			
		
		
		default:
			return fmt.Errorf("unknown option: %s", arg)
		}
	}

	return nil
}

func (c *Options) Usage() {
	fmt.Println(`USAGE:
  myapp [SUBCOMMAND] [OPTIONS]

SUBCOMMANDS:
    Add    添加记录
    Remove    删除记录

OPTIONS:
    -c, --config    指定配置文件路径 (支持 JSON, YAML, TOML)
    -h, --help      显示帮助信息
    -h, --help    显示帮助信息
    -v, --version    显示版本信息
    -d, --debug    启用调试模式
    -i, --input    输入文件路径
    -o, --output    输出文件路径
    --dir    工作目录
    -f, --format    输出格式
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

