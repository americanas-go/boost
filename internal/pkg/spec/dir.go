package spec

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/americanas-go/boost/internal/pkg/model"
)

func ParseDir(path string) (model.Spec, error) {
	d, err := parseDir(path)
	if err != nil {
		return model.Spec{}, err
	}

	rangePackage(d)

	return model.Spec{}, nil
}

func parseDir(path string) (map[string]*ast.Package, error) {
	fset := token.NewFileSet() // positions are relative to fset
	d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func rangePackage(d map[string]*ast.Package) {
	for _, f := range d {
		fmt.Println("package", f.Name)
		rangeFiles(f)
	}
}

func rangeFiles(f *ast.Package) {
	for k, f := range f.Files {
		fmt.Printf("File name: %q\n", k)
		rangeCommentCroups(f)
	}
}

func rangeCommentCroups(f *ast.File) {
	for i, c := range f.Comments {
		fmt.Printf("Comment Group %d\n", i)
		rangeComments(c)
	}
}

func rangeComments(c *ast.CommentGroup) {
	for i2, c1 := range c.List {
		fmt.Printf("Comment %d: Position: %d, Text: %q\n", i2, c1.Slash, c1.Text)
	}
}
