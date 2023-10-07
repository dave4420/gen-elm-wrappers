package main

import "fmt"

type module interface {
	name() string
	source() []string
}

type dictModule struct {
	moduleName  string
	typeName    string
	keyTypeName string
}

func (module dictModule) name() string {
	return module.moduleName
}

func (module dictModule) source() []string {
	return []string{
		"module " + module.moduleName + " exposing (..)",
	}
}

func writeModule(path string, module module) {
	fmt.Println("Writing module", module.name(), "to", path)
	for _, line := range module.source() {
		fmt.Println(line)
	}
}

type config struct {
	path    string
	modules []module
}

func readConfig() config {
	conf := config{
		path: "src",
		modules: []module{
			dictModule{
				moduleName:  "Data.ByDate",
				typeName:    "ByDate",
				keyTypeName: "Date",
			},
		},
	}
	return conf
}

func main() {
	var conf = readConfig()

	for _, module := range conf.modules {
		writeModule(conf.path, module)
	}
}
