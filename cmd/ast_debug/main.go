package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "/Users/guonaihong/my-github/quickclop/mytest/basic/basic.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}

	// Print all comments in the file
	fmt.Println("All Comments:")
	for _, cg := range file.Comments {
		fmt.Printf("Comment Group at position %v: %q\n", fset.Position(cg.Pos()), cg.Text())
	}

	// Print all declarations and their associated comments
	fmt.Println("\nDeclarations:")
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			fmt.Printf("GenDecl at position %v\n", fset.Position(genDecl.Pos()))
			fmt.Printf("  Doc: %v\n", genDecl.Doc != nil)
			if genDecl.Doc != nil {
				fmt.Printf("  Doc Text: %q\n", genDecl.Doc.Text())
			}

			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					fmt.Printf("  TypeSpec %q at position %v\n", typeSpec.Name.Name, fset.Position(typeSpec.Pos()))
					fmt.Printf("    Doc: %v\n", typeSpec.Doc != nil)
					if typeSpec.Doc != nil {
						fmt.Printf("    Doc Text: %q\n", typeSpec.Doc.Text())
					}
					
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						fmt.Printf("    StructType at position %v\n", fset.Position(structType.Pos()))
					}
				}
			}
		}
	}

	// Print the comment map
	fmt.Println("\nComment Map:")
	cmap := ast.NewCommentMap(fset, file, file.Comments)
	for node, commentGroups := range cmap {
		fmt.Printf("Node at position %v:\n", fset.Position(node.Pos()))
		for _, cg := range commentGroups {
			fmt.Printf("  Comment: %q\n", cg.Text())
		}
	}
}
