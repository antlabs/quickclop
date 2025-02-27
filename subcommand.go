package quickclop

import (
	"go/ast"
	"go/token"
	"strings"
)

// 解析子命令标签
func parseSubcommandTag(tag string) (string, string) {
	// 去掉首尾的反引号
	tag = strings.Trim(tag, "`")
	
	// 解析标签
	tags := parseTag(tag)
	
	// 获取subcmd标签
	if subcmdTag, ok := tags["subcmd"]; ok {
		parts := strings.Split(subcmdTag, ",")
		if len(parts) > 0 {
			name := strings.TrimSpace(parts[0])
			desc := ""
			if len(parts) > 1 {
				desc = strings.TrimSpace(parts[1])
			}
			return name, desc
		}
	}
	
	return "", ""
}

// 检查是否有子命令
func hasSubcommands(fields []FieldInfo) bool {
	for _, field := range fields {
		if field.IsNested {
			return true
		}
	}
	return false
}

// 解析子命令
func parseSubcommands(structType *ast.StructType, file *ast.File, fset *token.FileSet) []SubcommandInfo {
	var subcommands []SubcommandInfo
	
	// 遍历字段
	for _, field := range structType.Fields.List {
		// 检查是否有标签
		if field.Tag == nil {
			continue
		}
		
		// 解析标签
		name, desc := parseSubcommandTag(field.Tag.Value)
		if name == "" {
			continue
		}
		
		// 获取字段名
		var fieldName string
		if len(field.Names) > 0 {
			fieldName = field.Names[0].Name
		} else {
			// 匿名字段，使用类型名作为字段名
			switch t := field.Type.(type) {
			case *ast.Ident:
				fieldName = t.Name
			case *ast.SelectorExpr:
				fieldName = t.Sel.Name
			case *ast.StarExpr:
				if ident, ok := t.X.(*ast.Ident); ok {
					fieldName = ident.Name
				}
			}
		}
		
		// 获取字段类型
		var typeName string
		switch t := field.Type.(type) {
		case *ast.Ident:
			typeName = t.Name
		case *ast.SelectorExpr:
			if x, ok := t.X.(*ast.Ident); ok {
				typeName = x.Name + "." + t.Sel.Name
			}
		case *ast.StarExpr:
			if ident, ok := t.X.(*ast.Ident); ok {
				typeName = "*" + ident.Name
			}
		}
		
		// 查找结构体定义
		structDef := findStructDef(typeName, file, fset.Position(field.Pos()).Filename, "")
		if structDef == nil {
			continue
		}
		
		// 解析子命令结构体字段
		var cmdFields []FieldInfo
		for _, field := range structDef.Fields.List {
			if len(field.Names) > 0 {
				info := parseField(field, file, fset)
				cmdFields = append(cmdFields, info)
			}
		}
		
		// 添加子命令信息
		subcommands = append(subcommands, SubcommandInfo{
			Name:        fieldName,
			Description: desc,
			Fields:      cmdFields,
			StructType:  structDef,
		})
	}
	
	return subcommands
}

// 生成子命令代码
func generateSubcommandCode(subcommand SubcommandInfo, outputFile string) error {
	// 这里将使用与主命令类似的模板，但添加子命令处理逻辑
	return nil
}

// 子命令信息
type SubcommandInfo struct {
	Name        string
	Description string
	Fields      []FieldInfo
	StructType  *ast.StructType
}
