package main

import "strings"

func (module dictModule) name() string {
	return module.typeId.moduleName
}

func dictDef(wrapperType identifier, privateKeyId identifier) definition {
	return definition{
		export: []string{wrapperType.name},
		source: []string{
			strings.Join(
				[]string{
					"type",
					wrapperType.name,
					"a",
					"=",
					wrapperType.name,
					"(Dict",
					privateKeyId.fullName(),
					"a)",
				},
				" ",
			),
		},
	}
}

func emptyDictDef(wrapperType identifier, publicKeyId identifier) definition {
	return definition{
		export: []string{"empty"},
		source: []string{
			"empty : " + wrapperType.name + " a",
			"empty = " + wrapperType.name + " Dict.empty",
		},
	}
}

func (module dictModule) source() []string {
	definitions := []definition{
		dictDef(module.typeId, module.privateKeyId),
		emptyDictDef(module.typeId, module.publicKeyId),
	}

	lines := []string{
		"module " + module.typeId.moduleName + " exposing (..)",
		module.typeId.importLine(),
		module.publicKeyId.importLine(),
		module.privateKeyId.importLine(),
	}

	for _, def := range definitions {
		lines = append(lines, def.source...)
	}

	return lines
}
