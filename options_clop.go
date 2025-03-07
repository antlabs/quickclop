// Code generated by clop-gen; DO NOT EDIT.

package quickclop

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

// Options的Parse方法，用于解析命令行参数
func (c *Options) Parse(args []string) error {
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

		case arg == "-i":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -i")
			}

			c.InputFile = args[i+1]

			i++

		case arg == "-o":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -o")
			}

			c.OutputFile = args[i+1]

			i++

		case arg == "-s":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -s")
			}

			c.StructName = args[i+1]

			i++

		case arg == "--shell":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --shell")
			}

			c.ShellType = args[i+1]

			i++

		case arg == "--completion":

			c.Completion = true

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
    -i    输入文件路径
    -o    输出文件路径
    -s    结构体名称
    --shell    生成补全脚本的 shell 类型 (bash, zsh, fish)
    --completion    仅生成补全脚本
    -h, --help    显示帮助信息
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
