package quickclop

import (
	"text/template"
)

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

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
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
{{ range .Fields }}{{ if not .IsNested }}{{ if or .Short .Long }}    {{ if .Short }}-{{ .Short }}{{ end }}{{ if and .Short .Long }}, {{ end }}{{ if .Long }}--{{ .Long }}{{ end }}    {{ .Usage }}
{{ end }}{{ end }}{{ end }}{{ if .HasArgsField }}
ARGS:
{{ range .Fields }}{{ if .Args }}    {{ .Name }}    {{ .Usage }}
{{ end }}{{ end }}{{ end }}` + "`" + `)
}
`))
