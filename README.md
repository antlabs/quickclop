# QuickClop

QuickClop 是 [clop](https://github.com/antlabs/clop) 的静态代码生成版本，适用于对性能要求比较高的场景。通过静态代码生成，避免了运行时反射带来的性能开销。

## 功能特点

- 基于结构体标签自动生成命令行参数解析代码
- 支持短选项和长选项（如 `-X; --request`）
- 支持各种数据类型，包括基本类型、指针类型、切片类型等
- 自动生成帮助信息
- 按需导入必要的包，优化生成的代码
- 详细的错误处理

更多功能详情请参考 [FEATURES.md](FEATURES.md)。

## 安装

```bash
go get github.com/antlabs/quickclop
```

## 使用方法

1. 定义带有 `clop` 标签的结构体，并添加 `:quickclop` 注释：

```go
package main

// MyApp represents a command-line application with various options.
// :quickclop
type MyApp struct {
    Name    string   `clop:"-n; --name" usage:"Application name"`
    Verbose bool     `clop:"-v; --verbose" usage:"Enable verbose output"`
    Files   []string `clop:"-f; --file" usage:"Input files"`
    Output  string   `clop:"args=output" usage:"Output file"`
}
```

2. 运行 quickclop 工具生成代码：

```bash
quickclop path/to/your/package
```

3. 在你的代码中使用生成的 `Parse` 和 `Usage` 方法：

```go
func main() {
    app := &MyApp{}
    if err := app.Parse(os.Args[1:]); err != nil {
        fmt.Println(err)
        app.Usage()
        os.Exit(1)
    }
    
    // 使用解析后的参数...
}
```

## 标签格式

- `clop:"-s; --long"` - 定义短选项和长选项
- `clop:"args=name"` - 定义命令行位置参数
- `usage:"description"` - 定义选项的描述文本

## 支持的类型

- 基本类型：`string`, `int`, `float64`, `bool`
- 指针类型：`*string`, `*int`, `*float64`, `*bool`
- 切片类型：`[]string` 等
- 更多类型支持正在开发中

## 许可证

MIT
