package quickclop

import (
	"text/template"
)

// 解析模板
var tmpl = template.Must(template.New("clop").Parse(`
{{ if .HasSubcommands }}
// 子命令处理函数类型
type CommandFunc func(args []string) error

// 子命令信息
type SubcommandInfo struct {
	Name        string
	Description string
	Func        CommandFunc
}
{{ end }}

// {{ .StructName }}的Parse方法，用于解析命令行参数
func (c *{{ .StructName }}) Parse(args []string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}

	// 设置默认值
	{{ range .Fields }}{{ if not .IsNested }}{{ if .Default }}
	{{ if eq .Type "string" }}
	c.{{ .Name }} = "{{ .Default }}"
	{{ else if eq .Type "*string" }}
	defaultStr := "{{ .Default }}"
	c.{{ .Name }} = &defaultStr
	{{ else if eq .Type "int" }}
	if val, err := strconv.Atoi("{{ .Default }}"); err == nil {
		c.{{ .Name }} = val
	}
	{{ else if eq .Type "*int" }}
	if val, err := strconv.Atoi("{{ .Default }}"); err == nil {
		c.{{ .Name }} = &val
	}
	{{ else if eq .Type "float64" }}
	if val, err := strconv.ParseFloat("{{ .Default }}", 64); err == nil {
		c.{{ .Name }} = val
	}
	{{ else if eq .Type "*float64" }}
	if val, err := strconv.ParseFloat("{{ .Default }}", 64); err == nil {
		c.{{ .Name }} = &val
	}
	{{ else if eq .Type "bool" }}
	if val, err := strconv.ParseBool("{{ .Default }}"); err == nil {
		c.{{ .Name }} = val
	}
	{{ else if eq .Type "*bool" }}
	if val, err := strconv.ParseBool("{{ .Default }}"); err == nil {
		c.{{ .Name }} = &val
	}
	{{ else if eq .Type "[]string" }}
	c.{{ .Name }} = strings.Split("{{ .Default }}", ",")
	{{ else if eq .Type "time.Duration" }}
	if duration, err := time.ParseDuration("{{ .Default }}"); err == nil {
		c.{{ .Name }} = duration
	}
	{{ else if eq .Type "*time.Duration" }}
	if duration, err := time.ParseDuration("{{ .Default }}"); err == nil {
		c.{{ .Name }} = &duration
	}
	{{ end }}
	{{ end }}{{ end }}{{ end }}

	// 从环境变量中读取值
	{{ range .Fields }}{{ if not .IsNested }}{{ if .EnvVar }}
	if envVal, ok := os.LookupEnv("{{ .EnvVar }}"); ok {
		{{ if eq .Type "string" }}
		c.{{ .Name }} = envVal
		{{ else if eq .Type "*string" }}
		val := envVal
		c.{{ .Name }} = &val
		{{ else if eq .Type "int" }}
		if val, err := strconv.Atoi(envVal); err == nil {
			c.{{ .Name }} = val
		}
		{{ else if eq .Type "*int" }}
		if val, err := strconv.Atoi(envVal); err == nil {
			c.{{ .Name }} = &val
		}
		{{ else if eq .Type "float64" }}
		if val, err := strconv.ParseFloat(envVal, 64); err == nil {
			c.{{ .Name }} = val
		}
		{{ else if eq .Type "*float64" }}
		if val, err := strconv.ParseFloat(envVal, 64); err == nil {
			c.{{ .Name }} = &val
		}
		{{ else if eq .Type "bool" }}
		if val, err := strconv.ParseBool(envVal); err == nil {
			c.{{ .Name }} = val
		}
		{{ else if eq .Type "*bool" }}
		if val, err := strconv.ParseBool(envVal); err == nil {
			c.{{ .Name }} = &val
		}
		{{ else if eq .Type "[]string" }}
		c.{{ .Name }} = strings.Split(envVal, ",")
		{{ else if eq .Type "time.Duration" }}
		if duration, err := time.ParseDuration(envVal); err == nil {
			c.{{ .Name }} = duration
		}
		{{ else if eq .Type "*time.Duration" }}
		if duration, err := time.ParseDuration(envVal); err == nil {
			c.{{ .Name }} = &duration
		}
		{{ else if eq .Type "time.Time" }}
		if t, err := time.Parse(time.RFC3339, envVal); err == nil {
			c.{{ .Name }} = t
		}
		{{ else if eq .Type "*time.Time" }}
		if t, err := time.Parse(time.RFC3339, envVal); err == nil {
			c.{{ .Name }} = &t
		}
		{{ end }}
	}
	{{ end }}{{ end }}{{ end }}

	{{ if .HasSubcommands }}
	// 检查是否有子命令
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		subcommand := args[0]
		subcommandArgs := args[1:]

		switch subcommand {
		{{ range .Subcommands }}
		case "{{ .Name }}":
			return c.{{ .Name }}.Parse(subcommandArgs)
		{{ end }}
		default:
			return fmt.Errorf("未知子命令: %s", subcommand)
		}
	}
	{{ end }}

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
		{{ range .Fields }}{{ if not .IsNested }}{{ if or .Short .Long }}
		case {{ if .Short }}arg == "-{{ .Short }}"{{ end }}{{ if and .Short .Long }} || {{ end }}{{ if .Long }}arg == "--{{ .Long }}"{{ end }}:
			{{ if eq .Type "bool" }}
			c.{{ .Name }} = true
			{{ else if eq .Type "*bool" }}
			val := true
			c.{{ .Name }} = &val
			{{ else }}
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for {{ if .Long }}--{{ .Long }}{{ else }}-{{ .Short }}{{ end }}")
			}
			{{ if eq .Type "int" }}
			val, err := strconv.Atoi(args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for {{ if .Long }}--{{ .Long }}{{ else }}-{{ .Short }}{{ end }}: %v", err)
			}
			c.{{ .Name }} = val
			{{ else if eq .Type "*int" }}
			val, err := strconv.Atoi(args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for {{ if .Long }}--{{ .Long }}{{ else }}-{{ .Short }}{{ end }}: %v", err)
			}
			c.{{ .Name }} = &val
			{{ else if eq .Type "float64" }}
			val, err := strconv.ParseFloat(args[i+1], 64)
			if err != nil {
				return fmt.Errorf("invalid value for {{ if .Long }}--{{ .Long }}{{ else }}-{{ .Short }}{{ end }}: %v", err)
			}
			c.{{ .Name }} = val
			{{ else if eq .Type "*float64" }}
			val, err := strconv.ParseFloat(args[i+1], 64)
			if err != nil {
				return fmt.Errorf("invalid value for {{ if .Long }}--{{ .Long }}{{ else }}-{{ .Short }}{{ end }}: %v", err)
			}
			c.{{ .Name }} = &val
			{{ else if eq .Type "string" }}
			c.{{ .Name }} = args[i+1]
			{{ else if eq .Type "*string" }}
			val := args[i+1]
			c.{{ .Name }} = &val
			{{ else if eq .Type "[]string" }}
			c.{{ .Name }} = strings.Split(args[i+1], ",")
			{{ else if eq .Type "time.Duration" }}
			duration, err := time.ParseDuration(args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for {{ if .Long }}--{{ .Long }}{{ else }}-{{ .Short }}{{ end }}: %v", err)
			}
			c.{{ .Name }} = duration
			{{ else if eq .Type "*time.Duration" }}
			duration, err := time.ParseDuration(args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for {{ if .Long }}--{{ .Long }}{{ else }}-{{ .Short }}{{ end }}: %v", err)
			}
			c.{{ .Name }} = &duration
			{{ else if eq .Type "time.Time" }}
			t, err := time.Parse(time.RFC3339, args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for {{ if .Long }}--{{ .Long }}{{ else }}-{{ .Short }}{{ end }}: %v", err)
			}
			c.{{ .Name }} = t
			{{ else if eq .Type "*time.Time" }}
			t, err := time.Parse(time.RFC3339, args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for {{ if .Long }}--{{ .Long }}{{ else }}-{{ .Short }}{{ end }}: %v", err)
			}
			c.{{ .Name }} = &t
			{{ end }}
			i++
			{{ end }}
		{{ end }}{{ end }}{{ end }}
		{{ if .HasArgsField }}
		case !strings.HasPrefix(arg, "-"):
			{{ range .Fields }}{{ if .Args }}
			{{ if eq .Type "string" }}
			c.{{ .Name }} = arg
			{{ else if eq .Type "*string" }}
			val := arg
			c.{{ .Name }} = &val
			{{ else if eq .Type "[]string" }}
			c.{{ .Name }} = append(c.{{ .Name }}, arg)
			{{ end }}
			{{ end }}
			{{ end }}
		{{ end }}
		default:
			return fmt.Errorf("unknown option: %s", arg)
		}
	}

	return nil
}

func (c *{{ .StructName }}) Usage() {
	fmt.Println(` + "`" + `USAGE:
  myapp {{ if .HasSubcommands }}[SUBCOMMAND] {{ end }}[OPTIONS]{{ if .HasArgsField }} [ARGS]{{ end }}
{{ if .HasSubcommands }}
SUBCOMMANDS:
{{ range .Subcommands }}    {{ .Name }}    {{ .Usage }}
{{ end }}{{ end }}
OPTIONS:
    -c, --config    指定配置文件路径 (支持 JSON, YAML, TOML)
    -h, --help      显示帮助信息
{{ range .Fields }}{{ if not .IsNested }}{{ if or .Short .Long }}    {{ if .Short }}-{{ .Short }}{{ end }}{{ if and .Short .Long }}, {{ end }}{{ if .Long }}--{{ .Long }}{{ end }}    {{ .Usage }}{{ if .EnvVar }} [env: {{ .EnvVar }}]{{ end }}{{ if .Default }} (default: {{ .Default }}){{ end }}
{{ end }}{{ end }}{{ end }}{{ if .HasArgsField }}
ARGS:
{{ range .Fields }}{{ if .Args }}    {{ if .ArgName }}{{ .ArgName }}{{ else }}{{ .Name }}{{ end }}    {{ .Usage }}{{ if .Default }} (default: {{ .Default }}){{ end }}
{{ end }}{{ end }}{{ end }}` + "`" + `)
}

func (c *{{ .StructName }}) loadFromConfigFile(configFile string) error {
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
`))
