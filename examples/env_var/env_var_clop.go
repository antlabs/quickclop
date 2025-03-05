
// Code2 generated by clop-gen; DO NOT EDIT.

package main
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

// Config的Parse方法，用于解析命令行参数
func (c *Config) Parse(args []string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}

	// 从环境变量中读取值

	if envVal := os.Getenv("APP_DEBUG"); envVal != "" {
		
		if val, err := strconv.ParseBool(envVal); err == nil {
			c.Debug = val
		}
		
	}

	if envVal := os.Getenv("APP_PORT"); envVal != "" {
		
		if val, err := strconv.Atoi(envVal); err == nil {
			c.Port = val
		}
		
	}

	if envVal := os.Getenv("APP_HOST"); envVal != "" {
		
		c.Host = envVal
		
	}

	if envVal := os.Getenv("APP_LOG_LEVEL"); envVal != "" {
		
		c.LogLevel = envVal
		
	}

	if envVal := os.Getenv("APP_DATA_DIR"); envVal != "" {
		
		c.DataDir = envVal
		
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
		
		case arg == "-d" || arg == "--debug":
			
			c.Debug = true
			
		
		case arg == "-p" || arg == "--port":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --port")
			}
			
			val, err := strconv.Atoi(args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for --port: %v", err)
			}
			c.Port = val
			
			i++
			
		
		case arg == "-h" || arg == "--host":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --host")
			}
			
			c.Host = args[i+1]
			
			i++
			
		
		case arg == "--log-level":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --log-level")
			}
			
			c.LogLevel = args[i+1]
			
			i++
			
		
		
		case !strings.HasPrefix(arg, "-"):
			
			
			
			
			
			
			c.DataDir = arg
			
			
			
		
		default:
			return fmt.Errorf("unknown option: %s", arg)
		}
	}

	return nil
}

func (c *Config) Usage() {
	fmt.Println(`USAGE:
  myapp [OPTIONS] [ARGS]

OPTIONS:
    -c, --config    指定配置文件路径 (支持 JSON, YAML, TOML)
    -h, --help      显示帮助信息
    -d, --debug    启用调试模式 [env: APP_DEBUG]
    -p, --port    服务端口 [env: APP_PORT]
    -h, --host    服务主机 [env: APP_HOST]
    --log-level    日志级别 [env: APP_LOG_LEVEL]

ARGS:
    data-dir    数据目录
`)
}

func (c *Config) loadFromConfigFile(configFile string) error {
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

