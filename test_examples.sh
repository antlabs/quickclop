#!/bin/bash

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

# 显示帮助信息
show_help() {
    echo "用法: $0 [选项] [示例名称]"
    echo ""
    echo "选项:"
    echo "  -h, --help    显示帮助信息"
    echo "  -v, --verbose 显示详细输出"
    echo ""
    echo "示例:"
    echo "  $0            测试所有示例"
    echo "  $0 slices     只测试 slices 示例"
    exit 0
}

# 解析命令行参数
VERBOSE=false
EXAMPLE_NAME=""

while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            ;;
        -v|--verbose)
            VERBOSE=true
            shift
            ;;
        *)
            EXAMPLE_NAME="$1"
            shift
            ;;
    esac
done

# 获取示例目录
EXAMPLES_DIR="./examples"
if [ -n "$EXAMPLE_NAME" ]; then
    if [ ! -d "$EXAMPLES_DIR/$EXAMPLE_NAME" ]; then
        echo -e "${RED}错误: 示例 '$EXAMPLE_NAME' 不存在${NC}"
        exit 1
    fi
    EXAMPLES="$EXAMPLES_DIR/$EXAMPLE_NAME"
else
    EXAMPLES=$(find "$EXAMPLES_DIR" -type d -depth 1)
fi

# 清理所有示例目录中的可执行文件
echo -e "${YELLOW}清理所有示例目录中的可执行文件...${NC}"
for example in $EXAMPLES; do
    if [ -d "$example" ]; then
        # 获取示例目录名
        EXAMPLE_NAME=$(basename "$example")
        # 删除可能存在的可执行文件（与目录名相同的文件）
        if [ -f "$example/$EXAMPLE_NAME" ]; then
            rm -f "$example/$EXAMPLE_NAME"
            echo "已删除: $example/$EXAMPLE_NAME"
        fi
    fi
done

# 编译 quickclop 工具
echo -e "${YELLOW}编译 quickclop 工具...${NC}"
cd ./cmd/quickclop
go build
if [ $? -ne 0 ]; then
    echo -e "${RED}编译 quickclop 工具失败${NC}"
    exit 1
fi
cd - > /dev/null

echo -e "${YELLOW}开始测试示例...${NC}"
echo ""

SUCCESS_COUNT=0
FAILURE_COUNT=0
TOTAL_COUNT=0

# 创建数组存储测试结果
declare -a EXAMPLE_NAMES=()
declare -a EXAMPLE_STATUSES=()
declare -a EXAMPLE_ERRORS=()

for example in $EXAMPLES; do
    EXAMPLE_NAME=$(basename "$example")
    TOTAL_COUNT=$((TOTAL_COUNT+1))
    
    echo -e "${YELLOW}测试示例: ${EXAMPLE_NAME}${NC}"
    
    # 切换到示例目录
    cd "$example" || continue
    
    # 清理之前生成的文件和可执行文件
    rm -f *_clop.go
    rm -f "$EXAMPLE_NAME"
    
    # 执行代码生成
    echo "执行代码生成..."
    GEN_CMD="../../quickclop"

    echo "执行命令: ${GEN_CMD}"
    
    if $VERBOSE; then
        eval $GEN_CMD
    else
        eval $GEN_CMD 2>/dev/null
    fi
    GEN_STATUS=$?
    
    if [ $GEN_STATUS -ne 0 ]; then
        echo -e "${RED}✘ 代码生成失败: ${EXAMPLE_NAME}${NC}"
        echo -e "${RED}失败命令: ${GEN_CMD}${NC}, pwd $(pwd)${NC}"
        
        if $VERBOSE; then
            echo "尝试使用 -s 参数指定结构体名称..."
            # 查找可能的结构体名称
            STRUCT_NAME=$(grep -E "type +[A-Z][a-zA-Z0-9]* +struct" *.go | head -1 | sed -E 's/.*type +([A-Z][a-zA-Z0-9]*) +struct.*/\1/')
            if [ -n "$STRUCT_NAME" ]; then
                echo "找到结构体: $STRUCT_NAME，尝试重新生成..."
                # 查找输入和输出文件名
                INPUT_FILE=$(grep -l "type $STRUCT_NAME struct" *.go | head -1)
                OUTPUT_FILE="${INPUT_FILE%.*}_clop.go"
                
                if [ -z "$INPUT_FILE" ]; then
                    INPUT_FILE="main.go"
                    OUTPUT_FILE="main_clop.go"
                fi
                
                GEN_CMD="../../cmd/quickclop/quickclop -i \"$INPUT_FILE\" -o \"$OUTPUT_FILE\" -s \"$STRUCT_NAME\""
                echo "执行命令: ${GEN_CMD}"
                eval $GEN_CMD
                GEN_STATUS=$?
                if [ $GEN_STATUS -ne 0 ]; then
                    echo -e "${RED}✘ 使用结构体名称重新生成仍然失败${NC}"
                    echo -e "${RED}失败命令: ${GEN_CMD}${NC}"
                    FAILURE_COUNT=$((FAILURE_COUNT+1))
                    EXAMPLE_NAMES+=("$EXAMPLE_NAME")
                    EXAMPLE_STATUSES+=("失败")
                    EXAMPLE_ERRORS+=("代码生成失败")
                    cd - > /dev/null
                    echo ""
                    continue
                fi
            else
                FAILURE_COUNT=$((FAILURE_COUNT+1))
                EXAMPLE_NAMES+=("$EXAMPLE_NAME")
                EXAMPLE_STATUSES+=("失败")
                EXAMPLE_ERRORS+=("代码生成失败：未找到结构体")
                cd - > /dev/null
                echo ""
                continue
            fi
        else
            FAILURE_COUNT=$((FAILURE_COUNT+1))
            EXAMPLE_NAMES+=("$EXAMPLE_NAME")
            EXAMPLE_STATUSES+=("失败")
            EXAMPLE_ERRORS+=("代码生成失败")
            cd - > /dev/null
            echo ""
            continue
        fi
    fi
    
    
    # 尝试编译
    echo "编译代码..."
    if $VERBOSE; then
        go build
    else
        go build 2>/dev/null
    fi
    BUILD_STATUS=$?
    
    if [ $BUILD_STATUS -ne 0 ]; then
        echo -e "${RED}✘ 编译失败: ${EXAMPLE_NAME}${NC}"
        FAILURE_COUNT=$((FAILURE_COUNT+1))
        EXAMPLE_NAMES+=("$EXAMPLE_NAME")
        EXAMPLE_STATUSES+=("失败")
        EXAMPLE_ERRORS+=("编译失败")
        cd - > /dev/null
        echo ""
        continue
    fi
    
    # 尝试运行 (不捕获输出，除非是详细模式)
    echo "运行代码..."
    
    # 检查是否需要取消注释 quickclop.MustRun
    UNCOMMENT_NEEDED=$(grep -l "// quickclop.MustRun" *.go)
    if [ -n "$UNCOMMENT_NEEDED" ]; then
        echo "检测到需要取消注释 quickclop.MustRun 行..."
        for file in $UNCOMMENT_NEEDED; do
            sed -i '' 's|// quickclop.MustRun|quickclop.MustRun|g' "$file"
        done
    fi
    
    if $VERBOSE; then
        go run .
    else
        go run . 2>/dev/null 1>/dev/null
    fi
    RUN_STATUS=$?
    
    # 恢复注释，以便下次测试
    if [ -n "$UNCOMMENT_NEEDED" ]; then
        for file in $UNCOMMENT_NEEDED; do
            sed -i '' 's|quickclop.MustRun|// quickclop.MustRun|g' "$file"
        done
    fi
    
    if [ $RUN_STATUS -ne 0 ]; then
        echo -e "${RED}✘ 运行失败: ${EXAMPLE_NAME}${NC}"
        FAILURE_COUNT=$((FAILURE_COUNT+1))
        EXAMPLE_NAMES+=("$EXAMPLE_NAME")
        EXAMPLE_STATUSES+=("失败")
        EXAMPLE_ERRORS+=("运行失败")
    else
        echo -e "${GREEN}✓ 测试成功: ${EXAMPLE_NAME}${NC}"
        SUCCESS_COUNT=$((SUCCESS_COUNT+1))
        EXAMPLE_NAMES+=("$EXAMPLE_NAME")
        EXAMPLE_STATUSES+=("成功")
        EXAMPLE_ERRORS+=("无")
    fi
    
    # 清理编译生成的可执行文件
    if [ -f "$EXAMPLE_NAME" ]; then
        rm -f "$EXAMPLE_NAME"
        echo "已删除编译生成的可执行文件: $EXAMPLE_NAME"
    fi
    
    # 返回上级目录
    cd - > /dev/null
    echo ""
done

# 输出总结
echo -e "${YELLOW}测试总结:${NC}"
echo -e "总计: $TOTAL_COUNT 个示例"
echo -e "${GREEN}成功: $SUCCESS_COUNT 个${NC}"
if [ $FAILURE_COUNT -gt 0 ]; then
    echo -e "${RED}失败: $FAILURE_COUNT 个${NC}"
else
    echo -e "失败: $FAILURE_COUNT 个"
fi

# 输出 Markdown 表格
echo ""
echo "## 测试结果表格"
echo ""
echo "| 示例名称 | 状态 | 错误信息 |"
echo "| --- | --- | --- |"

for i in "${!EXAMPLE_NAMES[@]}"; do
    STATUS_ICON="❌"
    if [ "${EXAMPLE_STATUSES[$i]}" == "成功" ]; then
        STATUS_ICON="✅"
    fi
    echo "| ${EXAMPLE_NAMES[$i]} | ${STATUS_ICON} ${EXAMPLE_STATUSES[$i]} | ${EXAMPLE_ERRORS[$i]} |"
done

# 设置退出状态
if [ $FAILURE_COUNT -gt 0 ]; then
    exit 1
else
    exit 0
fi
