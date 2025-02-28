all: build

bootstrap:
	# 不需要预先生成options_clop.go文件
	go build -o quickclop_bootstrap ./cmd/quickclop/quickclop.go || true
	rm -f quickclop_bootstrap

build: bootstrap
	go build -o quickclop ./cmd/quickclop/quickclop.go
	go build -o ast_debug ./cmd/ast_debug/main.go

examples:
	go build ./examples/...

clean-completion:
	find . \( -name "*_completion.bash" -o -name "*_completion.fish" -o -name "*_completion.zsh" \) -type f -exec rm -f {} \;

clean: clean-completion
	find . -name "*_clop.go" -type f -exec rm -f {} \;
	rm -f quickclop ast_debug
	rm -rf bin