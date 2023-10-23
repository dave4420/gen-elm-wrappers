package main

import "strings"

func (module dictModule) name() string {
	return module.wrapperType.moduleName
}

func dictDef(wrapperType identifier, privateKeyId identifier) definition {
	return definition{
		localName: wrapperType.name,
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
		localName: "empty",
		source: []string{
			"empty : " + wrapperType.name + " a",
			"empty = " + wrapperType.name + " Dict.empty",
		},
	}
}

func singletonDictDef(wrapperType identifier, publicKeyId identifier, unwrapKeyFn identifier) definition {
	return definition{
		localName: "singleton",
		source: []string{
			"singleton : " + publicKeyId.fullName() + " -> v -> " + wrapperType.name + " v",
			"singleton k v = " + wrapperType.name + " (Dict.singleton (" + unwrapKeyFn.fullName() + " k) v)",
		},
	}
}

func (module dictModule) source() []string {
	definitions := []definition{
		dictDef(module.wrapperType, module.privateKeyType),
		emptyDictDef(module.wrapperType, module.publicKeyType),
		singletonDictDef(module.wrapperType, module.publicKeyType, module.unwrapKeyFn),
	}

	exports := []string{}
	for _, export := range definitions {
		exports = append(exports, export.localName)
	}

	lines := []string{
		"module " + module.wrapperType.moduleName + " exposing (" + strings.Join(exports, ", ") + ")",
		module.publicKeyType.importLine(),
		module.privateKeyType.importLine(),
		module.wrapKeyFn.importLine(),
		module.unwrapKeyFn.importLine(),
	}

	for _, def := range definitions {
		lines = append(lines, def.source...)
	}

	return lines
}
