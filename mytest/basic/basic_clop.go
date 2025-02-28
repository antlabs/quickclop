// Code generated by clop-gen; DO NOT EDIT.

package basic

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"github.com/BurntSushi/toml"
	"path/filepath"
)




// pcurl的Parse方法，用于解析命令行参数
func (c *pcurl) Parse(args []string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}

	// 设置默认值
	

	// 从环境变量中读取值
	

	

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
		
		case arg == "-X" || arg == "--request":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --request")
			}
			
			c.Method = args[i+1]
			
			i++
			
		
		case arg == "-G" || arg == "--get":
			
			c.Get = true
			
		
		case arg == "-H" || arg == "--header":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --header")
			}
			
			c.Header = strings.Split(args[i+1], ",")
			
			i++
			
		
		case arg == "-d" || arg == "--data":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --data")
			}
			
			c.Data = args[i+1]
			
			i++
			
		
		case arg == "--data-raw":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --data-raw")
			}
			
			c.DataRaw = args[i+1]
			
			i++
			
		
		case arg == "-F" || arg == "--form":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --form")
			}
			
			c.Form = strings.Split(args[i+1], ",")
			
			i++
			
		
		case arg == "--url":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --url")
			}
			
			c.URL = args[i+1]
			
			i++
			
		
		case arg == "-L" || arg == "--location":
			
			c.Location = true
			
		
		case arg == "--data-urlencode":
			
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --data-urlencode")
			}
			
			c.DataUrlencode = strings.Split(args[i+1], ",")
			
			i++
			
		
		case arg == "--compressed":
			
			c.Compressed = true
			
		
		case arg == "-i" || arg == "--include":
			
			c.Include = true
			
		
		case arg == "-k" || arg == "--insecure":
			
			c.Insecure = true
			
		
		
		case !strings.HasPrefix(arg, "-"):
			
			
			
			
			
			
			
			
			c.URL2 = arg
			
			
			
			
			
			
			
			
			
		
		default:
			return fmt.Errorf("unknown option: %s", arg)
		}
	}

	return nil
}

func (c *pcurl) Usage() {
	fmt.Println(`USAGE:
  myapp [OPTIONS] [ARGS]

OPTIONS:
    -c, --config    指定配置文件路径 (支持 JSON, YAML, TOML)
    -h, --help      显示帮助信息
    -X, --request    Specify request command to use
    -G, --get    Put the post data in the URL and use GET
    -H, --header    Pass custom header(s) to server
    -d, --data    HTTP POST data
    --data-raw    HTTP POST data, '@' allowed
    -F, --form    Specify multipart MIME data
    --url    URL to work with
    -L, --location    Follow redirects
    --data-urlencode    HTTP POST data url encoded
    --compressed    Request compressed response
    -i, --include    Include the HTTP response headers in the output. The HTTP response headers can include things like server name, cookies, date of the document, HTTP version and more.
    -k, --insecure    Allow insecure server connections when using SSL

ARGS:
    url2    url2
`)
}

func (c *pcurl) loadFromConfigFile(configFile string) error {
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
