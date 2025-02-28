package quickclop

import (
	"text/template"
)

// Bash 补全脚本模板
var bashCompletionTmpl = template.Must(template.New("bash_completion").Parse(`
# bash completion for {{ .AppName }}

_{{ .AppName }}() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    
    # 基本选项
    opts="--help -h --config -c {{ range .Fields }}{{ if not .IsNested }}{{ if .Short }}-{{ .Short }} {{ end }}{{ if .Long }}--{{ .Long }} {{ end }}{{ end }}{{ end }}"
    
    # 子命令
    {{ if .HasSubcommands }}
    subcmds="{{ range .Subcommands }}{{ .Name }} {{ end }}"
    
    # 如果当前是第一个参数，可能是子命令
    if [[ ${COMP_CWORD} -eq 1 ]]; then
        COMPREPLY=( $(compgen -W "${opts} ${subcmds}" -- ${cur}) )
        return 0
    fi
    
    # 检查是否在处理子命令
    for subcmd in ${subcmds}; do
        if [[ "${COMP_WORDS[1]}" == "${subcmd}" ]]; then
            case "${subcmd}" in
                {{ range .Subcommands }}
                "{{ .Name }}")
                    # 子命令特定的补全
                    _{{ $.AppName }}_{{ .Name }} "${cur}" "${prev}"
                    return 0
                    ;;
                {{ end }}
            esac
        fi
    done
    {{ end }}
    
    # 选项补全
    case "${prev}" in
        {{ range .Fields }}{{ if not .IsNested }}{{ if or .Short .Long }}
        {{ if .Short }}"-{{ .Short }}"{{ end }}{{ if and .Short .Long }}|{{ end }}{{ if .Long }}"--{{ .Long }}"{{ end }})
            {{ if eq .Completion "file" }}
            # 文件补全
            COMPREPLY=( $(compgen -f "${cur}") )
            {{ else if eq .Completion "dir" }}
            # 目录补全
            COMPREPLY=( $(compgen -d "${cur}") )
            {{ else if eq .Completion "custom" }}
            # 自定义补全
            local values="{{ .CompletionValues }}"
            COMPREPLY=( $(compgen -W "${values}" -- ${cur}) )
            {{ else }}
            # 无特定补全
            COMPREPLY=()
            {{ end }}
            return 0
            ;;
        {{ end }}{{ end }}{{ end }}
        "--config"|"-c")
            # 配置文件补全
            COMPREPLY=( $(compgen -f "${cur}") )
            return 0
            ;;
        *)
            ;;
    esac
    
    # 默认补全所有选项
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
}

{{ range .Subcommands }}
# {{ .Name }} 子命令的补全函数
_{{ $.AppName }}_{{ .Name }}() {
    local cur="${1}"
    local prev="${2}"
    local opts="{{ range .Fields }}{{ if not .IsNested }}{{ if .Short }}-{{ .Short }} {{ end }}{{ if .Long }}--{{ .Long }} {{ end }}{{ end }}{{ end }}"
    
    case "${prev}" in
        {{ range .Fields }}{{ if not .IsNested }}{{ if or .Short .Long }}
        {{ if .Short }}"-{{ .Short }}"{{ end }}{{ if and .Short .Long }}|{{ end }}{{ if .Long }}"--{{ .Long }}"{{ end }})
            {{ if eq .Completion "file" }}
            # 文件补全
            COMPREPLY=( $(compgen -f "${cur}") )
            {{ else if eq .Completion "dir" }}
            # 目录补全
            COMPREPLY=( $(compgen -d "${cur}") )
            {{ else if eq .Completion "custom" }}
            # 自定义补全
            local values="{{ .CompletionValues }}"
            COMPREPLY=( $(compgen -W "${values}" -- ${cur}) )
            {{ else }}
            # 无特定补全
            COMPREPLY=()
            {{ end }}
            return 0
            ;;
        {{ end }}{{ end }}{{ end }}
        *)
            ;;
    esac
    
    # 默认补全所有选项
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
}
{{ end }}

complete -F _{{ .AppName }} {{ .AppName }}
`))

// Zsh 补全脚本模板
var zshCompletionTmpl = template.Must(template.New("zsh_completion").Parse(`
#compdef {{ .AppName }}

_{{ .AppName }}() {
    local line state
    
    _arguments -C \
        {{ range .Fields }}{{ if not .IsNested }}{{ if or .Short .Long }}
        {{ if and .Short .Long }}'({{ if .Short }}-{{ .Short }}{{ end }} {{ if .Long }}--{{ .Long }}{{ end }})'{{ end }}{{ if .Short }}'-{{ .Short }}[{{ .Usage }}]{{ if eq .Completion "file" }}:file:_files{{ else if eq .Completion "dir" }}:directory:_files -/{{ else if eq .Completion "custom" }}:custom:({{ .CompletionValues }}){{ end }}'{{ end }} \
        {{ if and .Short .Long }}'({{ if .Short }}-{{ .Short }}{{ end }} {{ if .Long }}--{{ .Long }}{{ end }})'{{ end }}{{ if .Long }}'--{{ .Long }}[{{ .Usage }}]{{ if eq .Completion "file" }}:file:_files{{ else if eq .Completion "dir" }}:directory:_files -/{{ else if eq .Completion "custom" }}:custom:({{ .CompletionValues }}){{ end }}'{{ end }} \
        {{ end }}{{ end }}{{ end }}
        '--config[指定配置文件路径]:file:_files' \
        '--help[显示帮助信息]' \
        {{ if .HasSubcommands }}
        '1: :->command' \
        '*::arg:->args' \
        {{ end }}
    
    {{ if .HasSubcommands }}
    case $state in
        command)
            _values 'command' \
                {{ range .Subcommands }}
                '{{ .Name }}[{{ .Usage }}]' \
                {{ end }}
            ;;
        args)
            case $line[1] in
                {{ range .Subcommands }}
                {{ .Name }})
                    _{{ $.AppName }}_{{ .Name }}
                    ;;
                {{ end }}
            esac
            ;;
    esac
    {{ end }}
}

{{ range .Subcommands }}
_{{ $.AppName }}_{{ .Name }}() {
    _arguments \
        {{ range .Fields }}{{ if not .IsNested }}{{ if or .Short .Long }}
        {{ if and .Short .Long }}'({{ if .Short }}-{{ .Short }}{{ end }} {{ if .Long }}--{{ .Long }}{{ end }})'{{ end }}{{ if .Short }}'-{{ .Short }}[{{ .Usage }}]{{ if eq .Completion "file" }}:file:_files{{ else if eq .Completion "dir" }}:directory:_files -/{{ else if eq .Completion "custom" }}:custom:({{ .CompletionValues }}){{ end }}'{{ end }} \
        {{ if and .Short .Long }}'({{ if .Short }}-{{ .Short }}{{ end }} {{ if .Long }}--{{ .Long }}{{ end }})'{{ end }}{{ if .Long }}'--{{ .Long }}[{{ .Usage }}]{{ if eq .Completion "file" }}:file:_files{{ else if eq .Completion "dir" }}:directory:_files -/{{ else if eq .Completion "custom" }}:custom:({{ .CompletionValues }}){{ end }}'{{ end }} \
        {{ end }}{{ end }}{{ end }}
        '--help[显示帮助信息]'
}
{{ end }}

_{{ .AppName }}
`))

// Fish 补全脚本模板
var fishCompletionTmpl = template.Must(template.New("fish_completion").Parse(`
# fish completion for {{ .AppName }}

# 全局选项
complete -c {{ .AppName }} -l help -s h -d "显示帮助信息"
complete -c {{ .AppName }} -l config -s c -d "指定配置文件路径" -r -F

{{ range .Fields }}{{ if not .IsNested }}{{ if or .Short .Long }}
# {{ .Usage }}
{{ if .Short }}complete -c {{ $.AppName }} -s {{ .Short }}{{ if not (eq .Type "bool") }} -r{{ end }}{{ if .Long }} -l {{ .Long }}{{ end }} -d "{{ .Usage }}"{{ if eq .Completion "file" }} -F{{ else if eq .Completion "dir" }} -F -a "(__fish_complete_directories)"{{ else if eq .Completion "custom" }} -a "{{ .CompletionValues }}"{{ end }}{{ end }}
{{ if and (not .Short) .Long }}complete -c {{ $.AppName }} -l {{ .Long }}{{ if not (eq .Type "bool") }} -r{{ end }} -d "{{ .Usage }}"{{ if eq .Completion "file" }} -F{{ else if eq .Completion "dir" }} -F -a "(__fish_complete_directories)"{{ else if eq .Completion "custom" }} -a "{{ .CompletionValues }}"{{ end }}{{ end }}
{{ end }}{{ end }}{{ end }}

{{ if .HasSubcommands }}
# 子命令
{{ range .Subcommands }}
complete -c {{ $.AppName }} -f -n "__fish_use_subcommand" -a {{ .Name }} -d "{{ .Usage }}"

# {{ .Name }} 子命令选项
{{ range .Fields }}{{ if not .IsNested }}{{ if or .Short .Long }}
{{ if .Short }}complete -c {{ $.AppName }} -n "__fish_seen_subcommand_from {{ .Name }}" -s {{ .Short }}{{ if not (eq .Type "bool") }} -r{{ end }}{{ if .Long }} -l {{ .Long }}{{ end }} -d "{{ .Usage }}"{{ if eq .Completion "file" }} -F{{ else if eq .Completion "dir" }} -F -a "(__fish_complete_directories)"{{ else if eq .Completion "custom" }} -a "{{ .CompletionValues }}"{{ end }}{{ end }}
{{ if and (not .Short) .Long }}complete -c {{ $.AppName }} -n "__fish_seen_subcommand_from {{ .Name }}" -l {{ .Long }}{{ if not (eq .Type "bool") }} -r{{ end }} -d "{{ .Usage }}"{{ if eq .Completion "file" }} -F{{ else if eq .Completion "dir" }} -F -a "(__fish_complete_directories)"{{ else if eq .Completion "custom" }} -a "{{ .CompletionValues }}"{{ end }}{{ end }}
{{ end }}{{ end }}{{ end }}
{{ end }}
{{ end }}
`))
