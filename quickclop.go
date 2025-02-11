package quickclop

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func processDirectory(rootDir string) error {
	return filepath.Walk(rootDir, walker)
}

func walker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("access path %s: %w", path, err)
	}

	if info.IsDir() || !strings.HasSuffix(path, ".go") {
		return nil
	}

	fmt.Printf("Processing file: %s\n", path)
	return processFile(path)
}

func processFile(path string) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return fmt.Errorf("解析文件失败: %w", err)
	}

	// 第一步：收集所有结构体定义
	structDefs := make(map[string]*ast.StructType)
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				structDefs[typeSpec.Name.Name] = structType
			}
		}
	}

	modified := false
	ast.Inspect(file, func(n ast.Node) bool {

		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if structType, ok := typeSpec.Type.(*ast.StructType); ok {
			fmt.Printf("处理结构体: %s (文件: %s), comment( %s)\n", typeSpec.Name.Name, path, typeSpec.Doc.Text())
			// 先检查注释是否符合要求
			if typeSpec.Doc == nil || typeSpec.Doc.Text() == "" {
				return true // 没有注释直接跳过
			}

			if !hasQuickClopComment(typeSpec.Doc) {
				return true // 注释不符合要求跳过
			}

			// 符合条件才记录结构体
			structDefs[typeSpec.Name.Name] = structType
			log.Printf("处理结构体: %s (文件: %s)", typeSpec.Name.Name, path)
		}

		modified = true
		return true
	})

	if modified {
		// return safeWriteFile(path, fset, file)
	}
	return nil
}

func parseField(f *ast.Field) fieldInfo {
	// 处理匿名字段
	if len(f.Names) == 0 {
		typeName := fmt.Sprintf("%s", f.Type)
		// 如果是嵌套结构体
		if structType := findStructDef(typeName, nil, ""); structType != nil {
			return fieldInfo{
				Name:       typeName,
				Type:       typeName,
				IsNested:   true,
				structType: structType,
			}
		}
	}

	info := fieldInfo{
		Name: f.Names[0].Name,
		Type: fmt.Sprintf("%s", f.Type),
	}

	if f.Tag != nil {
		tag := strings.Trim(f.Tag.Value, "`")
		clopTag := reflect.StructTag(tag).Get("clop")
		if clopTag != "" {
			parts := strings.Split(clopTag, ",")
			for _, part := range parts {
				kv := strings.SplitN(part, "=", 2)
				switch kv[0] {
				case "short":
					info.Short = kv[1]
				case "long":
					info.Long = kv[1]
				case "usage":
					info.Usage = kv[1]
				case "default":
					info.Default = kv[1]
				case "args":
					info.Args = true
				}
			}
		}
	}

	// Set default long name if not specified
	if info.Long == "" && info.Short != "" {
		info.Long = strings.ToLower(info.Name)
	}

	// 如果是嵌套结构体
	if structType := findStructDef(info.Type, nil, ""); structType != nil {
		info.IsNested = true
	}

	return info
}

func hasArgs(fields []fieldInfo) bool {
	for _, f := range fields {
		if f.Args {
			return true
		}
	}
	return false
}

// 入口
func Main(path string) {
	if path == "" {
		path = "."
	}
	processDirectory(path)
}
