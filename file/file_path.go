package file

import (
	"os"
	"path/filepath"
)

//获得当前程序的根目录（与程序所在的同级目录）
func ProGramRoot() string {
	mname := os.Args[0]
	path, _ := filepath.Abs(mname)
	rootpath := filepath.Dir(path)
	return rootpath
}
