package quickclop

// :quickclop
type Options struct {
	InputFile  string `clop:"-i" usage:"输入文件路径"`
	OutputFile string `clop:"-o" usage:"输出文件路径" default:""`
	StructName string `clop:"-s" usage:"结构体名称"`
	ShellType  string `clop:"--shell" usage:"生成补全脚本的 shell 类型 (bash, zsh, fish)"`
	Completion bool   `clop:"--completion" usage:"仅生成补全脚本"`
	Help       bool   `clop:"-h,--help" usage:"显示帮助信息"`
}
