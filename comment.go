package quickclop

import (
	"go/ast"
	"strings"
)

const defaultCommentName = "quickclop"

// 检查是否包含目标注释
func hasQuickClopComment(doc *ast.CommentGroup) bool {
	if doc == nil {
		return false
	}
	for _, c := range doc.List {
		if strings.Contains(c.Text, ":"+defaultCommentName) {
			return true
		}
	}
	return false
}
