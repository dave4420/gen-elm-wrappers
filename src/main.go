package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

func writeModule(basePath string, module module, elmConfig elmConfig) error {
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

	moduleSource, err := module.source(elmConfig)
	if err != nil {
		return err
	}
	lines := append(
		[]string{
			"-- Please do not edit this file.",
			"-- It was generated by https://github.com/dave4420/gen-elm-wrappers",
			"",
		},
		moduleSource...,
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

		elmFormatCmd.Err = elmFormatCmd.Run()

		err = elmFormatCmd.Stdout.(*os.File).Close()
		if err != nil {
			return err
		}

		return elmFormatCmd.Err
	} else {
		return os.WriteFile(path, buffer, 0666)
	}
}

func runMain() error {
	elmConfig, err := readElmConfig()
	if err != nil {
		return err
	}

	conf, err := readConfig()
	if err != nil {
		return err
	}

	for _, module := range conf.modules {
		err := writeModule(conf.path, module, elmConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

func runHelp() error {
	fmt.Println("Usage:")
	fmt.Println("    gen-elm-wrappers")
	fmt.Println("        Generate Elm source code based on the contents of elm.json and")
	fmt.Println("        gen-elm-wrappers.json.")
	fmt.Println("    gen-elm-wrappers help")
	fmt.Println("        Display this help text.")
	fmt.Println("    gen-elm-wrappers version")
	fmt.Println("        Display this program’s version (" + Version + ").")
	fmt.Println()
	fmt.Println("For more info, see https://github.com/dave4420/gen-elm-wrappers")
	fmt.Println()
	return nil
}

var Version string

func runVersion() error {
	fmt.Printf("gen-elm-wrappers version %s\n", Version)
	return nil
}

func main() {
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	params := []string{}
	var err error

	if len(os.Args) > 0 {
		params = os.Args[1:]
	}

	if slices.Equal(params, []string{}) {
		err = runMain()
	} else if slices.Equal(params, []string{"help"}) {
		err = runHelp()
	} else if slices.Equal(params, []string{"version"}) {
		err = runVersion()
	} else {
		fmt.Fprintf(os.Stderr, "gen-elm-wrappers: don’t understand command %v\n", params)
		err = runHelp()
		exitCode = 1
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "gen-elm-wrappers: %s\n", err.Error())
		exitCode = 1
	}
}
