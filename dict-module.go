package main

import "strings"

func (module dictModule) name() string {
	return module.typeId.moduleName
}

func (module dictModule) typeDefLine() string {
	return strings.Join(
		[]string{
			"type",
			module.typeId.name,
			"a",
			"=",
			module.typeId.name,
			"(Dict",
			module.privateKeyId.fullName(),
			"a)",
		},
		" ",
	)
}

func (module dictModule) source() []string {
	return []string{
		"module " + module.typeId.moduleName + " exposing (..)",
		module.typeId.importLine(),
		module.publicKeyId.importLine(),
		module.privateKeyId.importLine(),
		module.typeDefLine(),
	}
}
