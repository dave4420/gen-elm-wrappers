package main

import (
	"fmt"
	"os"
	"os/exec"
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

func run() error {
	conf, err := readConfig()
	if err != nil {
		return err
	}

	for _, module := range conf.modules {
		err := writeModule(conf.path, module)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gen-elm-wrappers: %s\n", err.Error())
		exitCode = 1
	}
}
