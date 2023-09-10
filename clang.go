package goclangast

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

type Options struct {
	ClangPath string
	Args      []string
}

func NewASTOptions(path string, opts Options) (*TranslationUnitDecl, error) {
	args := []string{"-Xclang", "-ast-dump=json", "-fsyntax-only"}
	args = append(args, opts.Args...)
	args = append(args, path)
	cmd := exec.Command(opts.ClangPath, args...)

	var b bytes.Buffer

	cmd.Stdout = &b
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("run: %v", err)
	}

	ast, err := ParseTU(&b)
	if err != nil {
		return nil, fmt.Errorf("parse: %v", err)
	}

	return ast, nil
}

func NewAST(path string) (*TranslationUnitDecl, error) {
	return NewASTOptions(path, Options{
		ClangPath: "clang",
	})
}
