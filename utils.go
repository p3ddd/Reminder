package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func abort(funcname string, err error) {
	panic(funcname + " failed: " + err.Error())
}

// 将 \\ 替换成 /
func ClearPath(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

// 获取程序执行目录
func GetExecDir() string {
	// 返回程序执行目录绝对路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	ExecDir := ClearPath(dir)

	return ExecDir
}

func GetIconPath() string {
	return path.Join(GetExecDir(), "assets/favicon.ico")
}
