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
- `env:"ENV_VAR"` - 定义绑定的环境变量名
- `default:"value"` - 定义选项的默认值

## 支持的类型

- 基本类型：`string`, `int`, `float64`, `bool`
- 指针类型：`*string`, `*int`, `*float64`, `*bool`
- 切片类型：`[]string` 等
- 更多类型支持正在开发中

## 环境变量支持

QuickClop 支持从环境变量中读取选项值。使用 `env` 标签可以将选项绑定到环境变量：

```go
// :quickclop
type Config struct {
    Debug bool   `clop:"-d; --debug" usage:"启用调试模式" env:"APP_DEBUG"`
    Port  int    `clop:"-p; --port" usage:"服务端口" env:"APP_PORT"`
    Host  string `clop:"--host" usage:"服务主机" env:"APP_HOST"`
}
```

环境变量的优先级低于命令行参数，即如果同时提供了环境变量和命令行参数，将使用命令行参数的值。

## 默认值支持

QuickClop 支持为选项设置默认值。使用 `default` 标签可以指定选项的默认值：

```go
// :quickclop
type Config struct {
    Debug    bool   `clop:"-d; --debug" usage:"启用调试模式" default:"false"`
    Port     int    `clop:"-p; --port" usage:"服务端口" default:"8080"`
    Host     string `clop:"--host" usage:"服务主机" default:"localhost"`
    LogLevel string `clop:"--log-level" usage:"日志级别" default:"info"`
}
```

默认值的优先级最低，当命令行参数和环境变量都未提供时，将使用默认值。在帮助信息中，默认值会显示在选项描述后面。

## 配置文件支持

QuickClop 支持从配置文件中读取选项值。使用 `config` 标签可以将选项绑定到配置文件中的键：

```go
// :quickclop
type Config struct {
    Debug bool   `clop:"-d; --debug" usage:"启用调试模式" config:"debug"`
    Port  int    `clop:"-p; --port" usage:"服务端口" config:"port"`
    Host  string `clop:"--host" usage:"服务主机" config:"host"`
}
```

支持的配置文件格式包括：
- JSON (.json)
- YAML (.yaml, .yml)
- TOML (.toml)

使用 `-c` 或 `--config` 选项指定配置文件路径。

## Shell 补全脚本

QuickClop 支持生成 shell 补全脚本，目前支持 Bash、Zsh 和 Fish。

### 生成补全脚本

可以使用 `completion` 标志生成补全脚本：

```bash
# 生成所有类型的补全脚本
quickclop -i your_file.go -s YourStructName -completion

# 生成特定类型的补全脚本
quickclop -i your_file.go -s YourStructName -completion -shell bash
quickclop -i your_file.go -s YourStructName -completion -shell zsh
quickclop -i your_file.go -s YourStructName -completion -shell fish
```

### 使用补全标签

可以使用 `completion` 标签为选项指定补全类型：

```go
// :quickclop
type Options struct {
    // 文件补全
    Input  string `clop:"-i; --input" usage:"输入文件" completion:"file"`
    
    // 目录补全
    Dir    string `clop:"--dir" usage:"工作目录" completion:"dir"`
    
    // 自定义值补全
    Format string `clop:"-f; --format" usage:"输出格式" completion:"custom" completion_values:"json yaml toml xml"`
}
```

支持的补全类型：
- `file` - 文件补全
- `dir` - 目录补全
- `custom` - 自定义值补全，需要配合 `completion_values` 标签使用

### 安装补全脚本

生成的补全脚本可以按照以下方式安装：

**Bash**:
```bash
# 将脚本复制到 bash_completion.d 目录
sudo cp your_app_completion.bash /etc/bash_completion.d/
# 或添加到 .bashrc
echo "source /path/to/your_app_completion.bash" >> ~/.bashrc
```

**Zsh**:
```bash
# 复制到 zsh 补全目录
cp your_app_completion.zsh ~/.zsh/completions/_your_app
# 确保 fpath 包含补全目录
echo "fpath=(~/.zsh/completions $fpath)" >> ~/.zshrc
```

**Fish**:
```bash
# 复制到 fish 补全目录
cp your_app_completion.fish ~/.config/fish/completions/
```

## 许可证

apache 2.0
