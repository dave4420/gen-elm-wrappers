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

// DAVE: verify that the version number is not already in use
// DAVE: find out how to download a binary from an npm package
// DAVE: find out what architectures supported by go and node
//       node:
//        -  os.type()
//            -  Returns the operating system name as returned by uname(3).
//				 For example, it returns 'Linux' on Linux, 'Darwin' on macOS,
//				 and 'Windows_NT' on Windows.
// 				 https://linux.die.net/man/3/uname
//  			 https://en.wikipedia.org/wiki/Uname#Examples
//        -  os.arch()
//   		 - Possible values are 'arm', 'arm64', 'ia32', 'loong64', 'mips',
//			   'mipsel', 'ppc', 'ppc64', 'riscv64', 's390', 's390x', and 'x64'.
//       go:
//        -  $GOOS and $GOARCH
//        -  valid combinations listed at https://go.dev/doc/install/source#environment
// DAVE: get and save npm creds to 1password
// DAVE: install 1password in github actions and demo it fetching text from 1password
// DAVE: munge package.json to set version number and make it publishable
// DAVE: actually publish the package to npm
