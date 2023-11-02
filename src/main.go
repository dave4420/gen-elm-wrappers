package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func writeModule(basePath string, module module) error {
	elmFormatCmd := exec.Command("elm-format", "--stdin")

	pathSegments := []string{basePath}
	pathSegments = append(pathSegments, strings.Split(module.name(), ".")...)
	path := strings.Join(pathSegments, string(filepath.Separator)) + ".elm"

	if elmFormatCmd.Err == nil {
		fmt.Println("Writing formatted module", module.name(), "to", path)
	} else {
		fmt.Println("Writing unformatted module", module.name(), "to", path)
	}

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
	buffer := []byte(text)

	if elmFormatCmd.Err == nil {
		elmFormatCmd.Stdin = bytes.NewReader(buffer)

		elmFormatCmd.Stdout, err = os.Create(path)
		if err != nil {
			return err
		}

		elmFormatCmd.Stderr = os.Stderr

		// elmFormatCmd.Err = elmFormatCmd.Run()

		err = elmFormatCmd.Stdout.(*os.File).Close()
		if err != nil {
			return err
		}

		return elmFormatCmd.Err
	} else {
		return os.WriteFile(path, buffer, 0666)
	}
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
