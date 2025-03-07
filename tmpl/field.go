package tmpl

import "go/ast"

// FieldInfo 存储字段信息
type FieldInfo struct {
	Name       string // 字段名称
	Type       string // 字段类型
	Tag        string // 原始标签
	Short      string // 短选项，如 -h
	Long       string // 长选项，如 --help
	Usage      string // 使用说明
	Default    string // 默认值
	Args       bool   // 是否为位置参数
	Required   bool   // 是否必需
	IsNested   bool   // 是否为嵌套结构体
	ParseFunc  string // 解析函数名
	StructType *ast.StructType `json:"-" template:"-"` // 在模板渲染时忽略
	CmdName    string
	ArgName    string // 新增字段，用于存储参数名称
	EnvVar     string // 新增字段，用于存储环境变量名
	ConfigKey  string // 配置文件中的键名
	Completion string // 补全脚本的类型，例如 "file", "dir", "custom"
}
