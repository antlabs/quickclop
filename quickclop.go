package quickclop

import (
	"bytes"
	"embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	_ "embed"

	"github.com/antlabs/quickclop/tmpl"
)

type FieldInfo = tmpl.FieldInfo

//go:embed tmpl/*
var tmplFS embed.FS

// Initialize templates
var clopTemplate = template.Must(template.ParseFS(tmplFS, "tmpl/*.tmpl"))

// 解析字段标签
func parseTag(tag string) map[string]string {
	// 去除首尾的反引号
	tag = strings.Trim(tag, "`")

	// 解析标签
	result := make(map[string]string)
	for tag != "" {
		// 跳过前导空格
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		// 查找键
		i = 0
		for i < len(tag) && tag[i] != ':' {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i+1] != '"' {
			break
		}
		key := tag[:i]
		tag = tag[i+1:]

		// 查找值
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' && i+1 < len(tag) {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		value := tag[1:i]
		tag = tag[i+1:]

		// 跳过后续空格和逗号
		i = 0
		for i < len(tag) && (tag[i] == ' ' || tag[i] == ',') {
			i++
		}
		tag = tag[i:]

		result[key] = value
	}

	return result
}

// 检查是否有位置参数
func hasArgs(fields []FieldInfo) bool {
	for _, f := range fields {
		if f.Args {
			return true
		}
	}
	return false
}

// Main 是主入口函数
func Main(path string) {
	// 获取当前目录下的所有 Go 文件
	files, err := getGoFiles(path)
	if err != nil {
		log.Fatalf("获取 Go 文件失败: %v", err)
	}

	// 处理每个文件
	for _, file := range files {
		fmt.Printf("Processing file: %s\n", file)
		processFile(file)
	}
}

// 处理单个文件
func processFile(filePath string) {
	// 解析文件
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		log.Printf("解析文件 %s 失败: %v", filePath, err)
		return
	}

	// 获取包名
	packageName := file.Name.Name

	// 生成输出文件路径
	outputFile := generateOutputFilePath(filePath)

	// 检查是否有需要处理的结构体
	var structsToProcess []struct {
		Name string
		Type *ast.StructType
	}

	// 查找带有 :quickclop 注释的结构体
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

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// 检查是否有 :quickclop 注释
			if genDecl.Doc != nil {
				hasQuickClop := false
				for _, comment := range genDecl.Doc.List {
					if strings.Contains(comment.Text, ":quickclop") {
						hasQuickClop = true
						break
					}
				}

				if hasQuickClop {
					structName := typeSpec.Name.Name
					fmt.Printf("处理结构体: %s (文件: %s), comment(%s)\n", structName, filePath, genDecl.Doc.Text())
					log.Printf("处理结构体: %s (文件: %s)", structName, filePath)

					// 添加到待处理列表
					structsToProcess = append(structsToProcess, struct {
						Name string
						Type *ast.StructType
					}{
						Name: structName,
						Type: structType,
					})
				}
			}
		}
	}

	// 如果没有需要处理的结构体，直接返回
	if len(structsToProcess) == 0 {
		return
	}

	// 创建或清空输出文件
	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("创建输出文件失败: %v", err)
		return
	}
	defer f.Close() // 只创建文件，不写入内容

	// 处理每个结构体
	for _, s := range structsToProcess {
		// 生成代码
		err := generateCode(s.Name, s.Type, outputFile, packageName, file, fset)
		if err != nil {
			log.Printf("生成代码失败: %v", err)
			return
		}
	}

	log.Printf("生成代码文件: %s", outputFile)
}

// 获取指定路径下的所有 Go 文件
func getGoFiles(path string) ([]string, error) {
	// 检查路径是否为文件
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// 如果是文件，直接返回
	if !fileInfo.IsDir() {
		if strings.HasSuffix(path, ".go") {
			return []string{path}, nil
		}
		return nil, fmt.Errorf("不是 Go 文件: %s", path)
	}

	// 如果是目录，递归获取所有 Go 文件
	var files []string
	err = filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(p, ".go") && !strings.HasSuffix(p, "_clop.go") {
			files = append(files, p)
		}
		return nil
	})

	return files, err
}

// 生成输出文件路径
func generateOutputFilePath(inputPath string) string {
	dir := filepath.Dir(inputPath)
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	name := base[:len(base)-len(ext)]
	return filepath.Join(dir, name+"_clop.go")
}

// 解析字段
func parseField(field *ast.Field, file *ast.File, fset *token.FileSet) FieldInfo {
	info := FieldInfo{}

	// 获取字段名
	if len(field.Names) > 0 {
		info.Name = field.Names[0].Name
	} else {
		// 匿名字段，使用类型名作为字段名
		switch t := field.Type.(type) {
		case *ast.Ident:
			info.Name = t.Name
		case *ast.SelectorExpr:
			info.Name = t.Sel.Name
		case *ast.StarExpr:
			if ident, ok := t.X.(*ast.Ident); ok {
				info.Name = ident.Name
			}
		}
	}

	// 获取字段类型
	switch t := field.Type.(type) {
	case *ast.Ident:
		info.Type = t.Name
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			info.Type = x.Name + "." + t.Sel.Name
		}
	case *ast.StarExpr:
		switch xt := t.X.(type) {
		case *ast.Ident:
			info.Type = "*" + xt.Name
		case *ast.SelectorExpr:
			if x, ok := xt.X.(*ast.Ident); ok {
				info.Type = "*" + x.Name + "." + xt.Sel.Name
			}
		}
	case *ast.ArrayType:
		if ident, ok := t.Elt.(*ast.Ident); ok {
			info.Type = "[]" + ident.Name
		}
	}

	// 获取字段标签
	if field.Tag != nil {
		tag := field.Tag.Value
		// 去掉首尾的反引号
		tag = strings.Trim(tag, "`")

		// 检查是否有clop标签
		if strings.Contains(tag, "clop:") {
			// 提取clop标签内容
			clopParts := strings.Split(tag, "clop:")
			if len(clopParts) > 1 {
				clopTag := clopParts[1]
				if strings.Contains(clopTag, "\"") {
					clopTagParts := strings.Split(clopTag, "\"")
					if len(clopTagParts) > 1 {
						clopTag = clopTagParts[1]
					}
				}

				// 解析短选项和长选项
				if strings.Contains(clopTag, "-") && !strings.HasPrefix(clopTag, "args") {
					// 支持分号和逗号作为分隔符
					var parts []string
					if strings.Contains(clopTag, ";") {
						parts = strings.Split(clopTag, ";")
					} else if strings.Contains(clopTag, ",") {
						parts = strings.Split(clopTag, ",")
					} else {
						parts = []string{clopTag}
					}
					
					for _, part := range parts {
						part = strings.TrimSpace(part)
						if strings.HasPrefix(part, "-") && !strings.HasPrefix(part, "--") {
							// 短选项
							info.Short = strings.TrimPrefix(part, "-")
						} else if strings.HasPrefix(part, "--") {
							// 长选项
							info.Long = strings.TrimPrefix(part, "--")
						}
					}
				} else if strings.HasPrefix(clopTag, "args") {
					// 位置参数
					info.Args = true
					if strings.Contains(clopTag, "=") {
						argsParts := strings.Split(clopTag, "=")
						if len(argsParts) > 1 {
							info.ArgName = argsParts[1]
						}
					}
				}
			}

			// 解析usage
			if strings.Contains(tag, "usage:") {
				usageParts := strings.Split(tag, "usage:")
				if len(usageParts) > 1 {
					usageTag := usageParts[1]
					if strings.Contains(usageTag, "\"") {
						usageTagParts := strings.Split(usageTag, "\"")
						if len(usageTagParts) > 1 {
							info.Usage = usageTagParts[1]
						}
					}
				}
			}

			// 解析default
			if strings.Contains(tag, "default:") {
				defaultParts := strings.Split(tag, "default:")
				if len(defaultParts) > 1 {
					defaultTag := defaultParts[1]
					if strings.Contains(defaultTag, "\"") {
						defaultTagParts := strings.Split(defaultTag, "\"")
						if len(defaultTagParts) > 1 {
							info.Default = defaultTagParts[1]
						}
					}
				}
			}

			// 解析required
			if strings.Contains(tag, "required") {
				info.Required = true
			}
		}

		// 解析环境变量标签
		if strings.Contains(tag, "env:") {
			envParts := strings.Split(tag, "env:")
			if len(envParts) > 1 {
				envTag := envParts[1]
				if strings.Contains(envTag, "\"") {
					envTagParts := strings.Split(envTag, "\"")
					if len(envTagParts) > 1 {
						info.EnvVar = envTagParts[1]
					}
				}
			}
		}

		// 解析配置文件标签
		if strings.Contains(tag, "config:") {
			configParts := strings.Split(tag, "config:")
			if len(configParts) > 1 {
				configTag := configParts[1]
				if strings.Contains(configTag, "\"") {
					configTagParts := strings.Split(configTag, "\"")
					if len(configTagParts) > 1 {
						info.ConfigKey = configTagParts[1]
					}
				}
			}
		}

		// 解析补全脚本标签
		if strings.Contains(tag, "completion:") {
			completionParts := strings.Split(tag, "completion:")
			if len(completionParts) > 1 {
				completionTag := completionParts[1]
				if strings.Contains(completionTag, "\"") {
					completionTagParts := strings.Split(completionTag, "\"")
					if len(completionTagParts) > 1 {
						info.Completion = completionTagParts[1]
					}
				}
			}
		}

		// 获取subcmd标签
		if strings.Contains(tag, "subcmd:") {
			subcmdParts := strings.Split(tag, "subcmd:")
			if len(subcmdParts) > 1 {
				subcmdTag := subcmdParts[1]
				if strings.Contains(subcmdTag, "\"") {
					subcmdTagParts := strings.Split(subcmdTag, "\"")
					if len(subcmdTagParts) > 1 {
						info.IsNested = true
						info.CmdName = subcmdTagParts[1]
					}
				}
			}
		}
	}

	// 检查是否是嵌套结构体
	if info.Type != "" && !isBasicType(info.Type) && !isTimeType(info.Type) && !isURLType(info.Type) && !isIPType(info.Type) {
		// 查找结构体定义
		structDef := findStructDef(info.Type, file, fset.Position(field.Pos()).Filename, "")
		if structDef != nil {
			info.StructType = structDef
			info.IsNested = true
		}
	}

	return info
}

// 检查是否是基本类型
func isBasicType(typeName string) bool {
	basicTypes := map[string]bool{
		"bool":      true,
		"int":       true,
		"int8":      true,
		"int16":     true,
		"int32":     true,
		"int64":     true,
		"uint":      true,
		"uint8":     true,
		"uint16":    true,
		"uint32":    true,
		"uint64":    true,
		"float32":   true,
		"float64":   true,
		"string":    true,
		"byte":      true,
		"rune":      true,
		"*bool":     true,
		"*int":      true,
		"*int8":     true,
		"*int16":    true,
		"*int32":    true,
		"*int64":    true,
		"*uint":     true,
		"*uint8":    true,
		"*uint16":   true,
		"*uint32":   true,
		"*uint64":   true,
		"*float32":  true,
		"*float64":  true,
		"*string":   true,
		"*byte":     true,
		"*rune":     true,
		"[]string":  true,
		"[]int":     true,
		"[]int8":    true,
		"[]int16":   true,
		"[]int32":   true,
		"[]int64":   true,
		"[]uint":    true,
		"[]uint8":   true,
		"[]uint16":  true,
		"[]uint32":  true,
		"[]uint64":  true,
		"[]float32": true,
		"[]float64": true,
		"[]bool":    true,
		"[]byte":    true,
		"[]rune":    true,
	}

	return basicTypes[typeName]
}

// 检查是否是时间类型
func isTimeType(typeName string) bool {
	timeTypes := map[string]bool{
		"time.Time":      true,
		"time.Duration":  true,
		"*time.Time":     true,
		"*time.Duration": true,
	}

	return timeTypes[typeName]
}

// 检查是否是 URL 类型
func isURLType(typeName string) bool {
	urlTypes := map[string]bool{
		"url.URL":  true,
		"*url.URL": true,
	}

	return urlTypes[typeName]
}

// 检查是否是 IP 地址类型
func isIPType(typeName string) bool {
	ipTypes := map[string]bool{
		"net.IP":     true,
		"*net.IP":    true,
		"net.IPNet":  true,
		"*net.IPNet": true,
	}

	return ipTypes[typeName]
}

// 解析结构体上的版本信息标签
func parseVersionTag(file *ast.File) (string, string, string) {
	var version, description, buildInfo string

	// 遍历文件中的注释
	for _, comment := range file.Comments {
		for _, line := range comment.List {
			if strings.Contains(line.Text, ":quickclop") && strings.Contains(line.Text, "version:") {
				// 解析版本标签
				versionMatch := regexp.MustCompile(`version:"([^"]*)"`)
				if matches := versionMatch.FindStringSubmatch(line.Text); len(matches) > 1 {
					version = matches[1]
				}

				// 解析描述标签
				descMatch := regexp.MustCompile(`description:"([^"]*)"`)
				if matches := descMatch.FindStringSubmatch(line.Text); len(matches) > 1 {
					description = matches[1]
				}

				// 解析构建信息标签
				buildMatch := regexp.MustCompile(`build:"([^"]*)"`)
				if matches := buildMatch.FindStringSubmatch(line.Text); len(matches) > 1 {
					buildInfo = matches[1]
				}
			}
		}
	}

	return version, description, buildInfo
}

// 检查是否有版本标志
func hasVersionFlag(fields []FieldInfo) bool {
	for _, field := range fields {
		if (field.Type == "bool" || field.Type == "*bool") &&
			(field.Name == "Version" || field.Name == "version") &&
			(field.Short == "v" || field.Long == "version") {
			return true
		}
	}
	return false
}

// 生成代码
func generateCode(structName string, structType *ast.StructType, outputFile string, packageName string, file *ast.File, fset *token.FileSet) error {
	// 解析结构体字段
	var fields []FieldInfo
	for _, field := range structType.Fields.List {
		info := parseField(field, file, fset)
		if !info.IsNested {
			fields = append(fields, info)
		}
	}

	// 检查是否有子命令
	var subcommands []struct {
		Name   string
		Usage  string
		Fields []FieldInfo
	}

	for _, field := range structType.Fields.List {
		info := parseField(field, file, fset)
		if info.IsNested && info.StructType != nil {
			// 解析子命令的字段
			var subFields []FieldInfo
			for _, subField := range info.StructType.Fields.List {
				subInfo := parseField(subField, file, fset)
				subInfo.CmdName = info.Name
				subFields = append(subFields, subInfo)
			}

			// 添加子命令
			subcommands = append(subcommands, struct {
				Name   string
				Usage  string
				Fields []FieldInfo
			}{
				Name:   info.Name,
				Usage:  info.Usage,
				Fields: subFields,
			})
		}
	}

	// 解析版本信息
	version, description, buildInfo := parseVersionTag(file)
	hasVersionFlagValue := hasVersionFlag(fields)

	// 分析所有结构体的字段类型，确定需要导入的包
	needFmt := false
	needOs := false
	needStrconv := false
	needStrings := true
	needTime := false
	needJson := false
	needYaml := false
	needToml := false
	needFilepath := false
	needNet := false
	needUrl := false

	for _, field := range fields {
		// 检查字段类型，确定需要导入的包
		if strings.Contains(field.Type, "time.") {
			needTime = true
		}

		// 检查是否是 net 包类型
		if strings.Contains(field.Type, "net.") {
			needNet = true
		}

		// 检查是否是 url 包类型
		if strings.Contains(field.Type, "url.") {
			needUrl = true
		}

		// 数值类型需要 strconv 包（用于命令行参数解析）
		if strings.Contains(field.Type, "int") ||
			strings.Contains(field.Type, "float") ||
			strings.Contains(field.Type, "uint") {
			needStrconv = true
		}

		// 布尔类型通常不需要strconv，除非有默认值或环境变量
		if strings.Contains(field.Type, "bool") && (field.Default != "" || field.EnvVar != "") {
			needStrconv = true
		}

		// 字符串和切片类型可能需要额外的 strings 包功能
		if strings.Contains(field.Type, "string") ||
			strings.Contains(field.Type, "[]") {
			// needStrings 已经在全局设置为 true
		}

		// 检查是否有环境变量或默认值，这些也需要相应的包
		if field.EnvVar != "" {
			needOs = true

			if strings.Contains(field.Type, "time.") {
				needTime = true
			}

			if strings.Contains(field.Type, "net.") {
				needNet = true
			}

			if strings.Contains(field.Type, "url.") {
				needUrl = true
			}
		}

		if field.Default != "" {
			if strings.Contains(field.Type, "time.") {
				needTime = true
			}

			if strings.Contains(field.Type, "net.") {
				needNet = true
			}

			if strings.Contains(field.Type, "url.") {
				needUrl = true
			}
		}

		// 配置文件支持需要的包
		if field.ConfigKey != "" || field.Long != "" || field.Short != "" {
			needJson = true // 支持JSON配置文件
			needYaml = true // 支持YAML配置文件
			needToml = true // 支持TOML配置文件
		}
	}

	// 所有类型都可能需要这些包
	needFmt = true      // 用于错误信息和Usage输出
	needOs = true       // 用于os.Args和os.Exit
	needFilepath = true // 用于filepath.Ext
	needStrings = true  // 用于 strings.ToLower 和其他字符串操作

	// 检查是否需要生成配置文件加载功能
	for _, field := range fields {
		if field.Short == "c" || field.Long == "config" {
			// 配置文件加载需要的包
			needFilepath = true // 用于 filepath.Ext
			needFmt = true      // 用于 fmt.Errorf
			needOs = true       // 用于 os.ReadFile
			needJson = true     // 用于 json.Unmarshal
			needYaml = true     // 用于 yaml.Unmarshal
			needToml = true     // 用于 toml.Decode
			break
		}
	}

	// 准备模板数据
	data := struct {
		StructName     string
		Fields         []FieldInfo
		HasArgsField   bool
		HasSubcommands bool
		Subcommands    []struct {
			Name   string
			Usage  string
			Fields []FieldInfo
		}
		// 版本信息相关字段
		HasVersionFlag bool
		Version        string
		Description    string
		BuildInfo      string
		AppName        string
		PackageName    string
		// 导入包相关字段
		NeedFmt      bool
		NeedOs       bool
		NeedStrconv  bool
		NeedStrings  bool
		NeedTime     bool
		NeedJson     bool
		NeedYaml     bool
		NeedToml     bool
		NeedFilepath bool
		NeedNet      bool
		NeedUrl      bool
	}{
		StructName:     structName,
		Fields:         fields,
		HasArgsField:   hasArgs(fields),
		HasSubcommands: len(subcommands) > 0,
		Subcommands:    subcommands,
		// 版本信息相关字段赋值
		HasVersionFlag: hasVersionFlagValue,
		Version:        version,
		Description:    description,
		BuildInfo:      buildInfo,
		AppName:        structName,
		PackageName:    packageName,
		// 导入包相关字段赋值
		NeedFmt:      needFmt,
		NeedOs:       needOs,
		NeedStrconv:  needStrconv,
		NeedStrings:  needStrings,
		NeedTime:     needTime,
		NeedJson:     needJson,
		NeedYaml:     needYaml,
		NeedToml:     needToml,
		NeedFilepath: needFilepath,
		NeedNet:      needNet,
		NeedUrl:      needUrl,
	}

	// 打开输出文件进行写入（不是追加）
	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("打开输出文件失败: %w", err)
	}
	defer f.Close()

	var out bytes.Buffer
	// 执行模板
	if err := clopTemplate.ExecuteTemplate(&out, "main", data); err != nil {
		return fmt.Errorf("生成代码失败: %w", err)
	}

	formatCode, err := format.Source(out.Bytes())
	if err != nil {
		return fmt.Errorf("格式化代码失败: %w", err)
	}
	// 写入文件

	if _, err := f.Write(formatCode); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	return nil
}

func extractTagValue(tag string, key string) string {
	keyPrefix := key + ":"
	parts := strings.Split(tag, " ")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, keyPrefix) {
			return strings.TrimPrefix(part, keyPrefix)
		}
	}
	return ""
}

// 生成补全脚本
func generateCompletionScript(structName string, structType *ast.StructType, outputDir string, packageName string, file *ast.File, fset *token.FileSet, shellType string) error {
	// 解析结构体字段
	var fields []FieldInfo
	for _, field := range structType.Fields.List {
		info := parseField(field, file, fset)
		if !info.IsNested {
			fields = append(fields, info)
		}
	}

	// 检查是否有子命令
	var subcommands []struct {
		Name   string
		Usage  string
		Fields []FieldInfo
	}

	for _, field := range structType.Fields.List {
		info := parseField(field, file, fset)
		if info.IsNested && info.StructType != nil {
			// 解析子命令的字段
			var subFields []FieldInfo
			for _, subField := range info.StructType.Fields.List {
				subInfo := parseField(subField, file, fset)
				subInfo.CmdName = info.Name
				subFields = append(subFields, subInfo)
			}

			// 添加子命令
			subcommands = append(subcommands, struct {
				Name   string
				Usage  string
				Fields []FieldInfo
			}{
				Name:   info.Name,
				Usage:  info.Usage,
				Fields: subFields,
			})
		}
	}

	// 准备模板数据
	data := struct {
		AppName        string
		StructName     string
		Fields         []FieldInfo
		HasSubcommands bool
		Subcommands    []struct {
			Name   string
			Usage  string
			Fields []FieldInfo
		}
	}{
		AppName:        strings.ToLower(structName),
		StructName:     structName,
		Fields:         fields,
		HasSubcommands: len(subcommands) > 0,
		Subcommands:    subcommands,
	}

	// 选择适当的模板
	var tmpl *template.Template
	var ext string
	switch shellType {
	case "bash":
		tmpl = bashCompletionTmpl
		ext = ".bash"
	case "zsh":
		tmpl = zshCompletionTmpl
		ext = ".zsh"
	case "fish":
		tmpl = fishCompletionTmpl
		ext = ".fish"
	default:
		return fmt.Errorf("不支持的 shell 类型: %s", shellType)
	}

	// 生成输出文件路径
	outputFile := filepath.Join(outputDir, strings.ToLower(structName)+"_completion"+ext)

	// 创建或清空输出文件
	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("创建补全脚本文件失败: %w", err)
	}
	defer f.Close()

	// 执行模板
	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("生成补全脚本失败: %w", err)
	}

	log.Printf("生成补全脚本: %s", outputFile)
	return nil
}

// 创建输出文件
func createOutputFile(outputFile string, packageName string) error {
	// 检查文件是否存在
	_, err := os.Stat(outputFile)
	if err == nil {
		err := os.Remove(outputFile)
		if err != nil {
			return fmt.Errorf("清空文件失败: %w", err)
		}
	} else if !os.IsNotExist(err) {
		// 其他错误
		return fmt.Errorf("检查文件状态失败: %w", err)
	}

	// 创建或打开文件
	f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}
	defer f.Close()

	return nil
}

// Generate 生成代码
func Generate(options *Options) error {
	// 解析输入文件
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, options.InputFile, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("解析文件失败: %w", err)
	}

	// 获取包名
	packageName := node.Name.Name

	// 查找指定结构体
	var structType *ast.StructType
	ast.Inspect(node, func(n ast.Node) bool {
		// 检查是否是类型声明
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok || typeSpec.Name.Name != options.StructName {
			return true
		}

		// 检查是否是结构体
		structType, ok = typeSpec.Type.(*ast.StructType)
		return false
	})

	if structType == nil {
		return fmt.Errorf("找不到结构体: %s", options.StructName)
	}

	// 如果只生成补全脚本
	if options.Completion {
		outputDir := "."
		if options.OutputFile != "" {
			outputDir = filepath.Dir(options.OutputFile)
		}

		if options.ShellType != "" {
			// 生成指定类型的补全脚本
			if err := generateCompletionScript(options.StructName, structType, outputDir, packageName, node, fset, options.ShellType); err != nil {
				return fmt.Errorf("生成补全脚本失败: %w", err)
			}
		} else {
			// 当没有指定 ShellType 时，提示用户需要指定 shell 类型
			return fmt.Errorf("请使用 --shell 参数指定要生成的补全脚本类型 (bash, zsh, fish)")
		}
		return nil
	}

	// 如果输出文件路径为空，生成默认输出文件路径
	if options.OutputFile == "" {
		options.OutputFile = generateOutputFilePath(options.InputFile)
		log.Printf("未指定输出文件路径，使用默认路径: %s", options.OutputFile)
	}

	// 创建输出文件
	if err := createOutputFile(options.OutputFile, packageName); err != nil {
		return fmt.Errorf("创建输出文件失败: %w", err)
	}

	// 生成代码
	if err := generateCode(options.StructName, structType, options.OutputFile, packageName, node, fset); err != nil {
		return fmt.Errorf("生成代码失败: %w", err)
	}

	log.Printf("代码生成完成: %s", options.OutputFile)
	return nil
}
