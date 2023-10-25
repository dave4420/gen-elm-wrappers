package main

import "strings"

func (module dictModule) name() string {
	return module.wrapperType.moduleName
}

func (module dictModule) dictDef() definition {
	return definition{
		localName: module.wrapperType.name,
		source: []string{
			strings.Join(
				[]string{
					"type",
					module.wrapperType.name,
					"a",
					"=",
					module.wrapperType.name,
					"(Dict",
					module.privateKeyType.fullName(),
					"a)",
				},
				" ",
			),
		},
	}
}

func (module dictModule) emptyDictDef() definition {
	return definition{
		localName: "empty",
		source: []string{
			"empty : " + module.wrapperType.name + " a",
			"empty = " + module.wrapperType.name + " Dict.empty",
		},
	}
}

func (module dictModule) singletonDictDef() definition {
	return definition{
		localName: "singleton",
		source: []string{
			"singleton : " + module.publicKeyType.fullName() + " -> v -> " + module.wrapperType.name + " v",
			"singleton k v = " + module.wrapperType.name + " (Dict.singleton (" + module.unwrapKeyFn.fullName() + " k) v)",
		},
	}
}

func (module dictModule) insertDictDef() definition {
	return definition{
		localName: "insert",
		source: []string{
			"insert : " + module.publicKeyType.fullName() + " -> v -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v",
			"insert k v (" + module.wrapperType.name + " d) = " + module.wrapperType.name + " (Dict.insert (" + module.unwrapKeyFn.fullName() + " k) v d)",
		},
	}
}

func (module dictModule) updateDictDef() definition {
	return definition{
		localName: "update",
		source: []string{
			"update : " + module.publicKeyType.fullName() + " -> (Maybe v -> Maybe v) -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v",
			"update k f (" + module.wrapperType.name + " d) = " + module.wrapperType.name + " (Dict.update (" + module.unwrapKeyFn.fullName() + " k) f d)",
		},
	}
}

func (module dictModule) removeDictDef() definition {
	return definition{
		localName: "remove",
		source: []string{
			"remove : " + module.publicKeyType.fullName() + " -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v",
			"remove k (" + module.wrapperType.name + " d) = " + module.wrapperType.name + " (Dict.remove (" + module.unwrapKeyFn.fullName() + " k) d)",
		},
	}
}

func (module dictModule) isEmptyDictDef() definition {
	return definition{
		localName: "isEmpty",
		source: []string{
			"isEmpty : " + module.wrapperType.name + " v -> Bool",
			"isEmpty (" + module.wrapperType.name + " d) = Dict.isEmpty d",
		},
	}
}

func (module dictModule) memberDictDef() definition {
	return definition{
		localName: "member",
		source: []string{
			"member : " + module.publicKeyType.fullName() + " -> " + module.wrapperType.name + " v -> Bool",
			"member k (" + module.wrapperType.name + " d) = Dict.member (" + module.unwrapKeyFn.fullName() + " k) d",
		},
	}
}

func (module dictModule) getDictDef() definition {
	return definition{
		localName: "get",
		source: []string{
			"get : " + module.publicKeyType.fullName() + " -> " + module.wrapperType.name + " v -> Maybe v",
			"get k (" + module.wrapperType.name + " d) = Dict.get (" + module.unwrapKeyFn.fullName() + " k) d",
		},
	}
}

func (module dictModule) sizeDictDef() definition {
	return definition{
		localName: "size",
		source: []string{
			"size : " + module.wrapperType.name + " v -> Int",
			"size (" + module.wrapperType.name + " d) = Dict.size d",
		},
	}
}

func (module dictModule) source() []string {
	definitions := []definition{
		module.dictDef(),
		module.emptyDictDef(),
		module.singletonDictDef(),
		module.insertDictDef(),
		module.updateDictDef(),
		module.removeDictDef(),
		module.isEmptyDictDef(),
		module.memberDictDef(),
		module.getDictDef(),
		module.sizeDictDef(),
	}

	exports := []string{}
	for _, export := range definitions {
		exports = append(exports, export.localName)
	}

	lines := []string{
		"module " + module.wrapperType.moduleName + " exposing (" + strings.Join(exports, ", ") + ")",
		"import Dict exposing (Dict)",
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
