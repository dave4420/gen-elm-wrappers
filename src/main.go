package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func writeModule(basePath string, module module) error {
	pathSegments := []string{basePath}
	pathSegments = append(pathSegments, strings.Split(module.name(), ".")...)
	path := strings.Join(pathSegments, string(filepath.Separator)) + ".elm"
	fmt.Println("Writing module", module.name(), "to", path)
	dir, _ := filepath.Split(path)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	lines := append(
		[]string{
			"-- Please do not edit this file.",
			"-- It was generated by https://github.com/dave4420/gen-elm-wrappers",
			"",
		},
		module.source()...,
	)
	text := strings.Join(lines, "\n") + "\n"
	return os.WriteFile(path, []byte(text), 0666)
}

func main() {
	conf, _ := readConfig()
	// DAVE: handle error from readConfig()

	for _, module := range conf.modules {
		// DAVE: handle error from writeModule()
		writeModule(conf.path, module)
	}
}
