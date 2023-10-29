package main

func (module dictModule) extraDefs() []definition {
	if module.dictExtraVersion == "" {
		return []definition{}
	}
	return []definition{
		module.groupByDef(),
	}
}

// List operations

func (module dictModule) groupByDef() definition {
	return definition{
		localName: "groupBy",
		source: []string{
			"groupBy : (a -> " + module.publicKeyType.fullName() + ") -> List a -> " + module.wrapperType.name + " (List a)",
			"groupBy f xs = " + module.wrapperType.name + " (Dict.Extra.groupBy (\\x -> " + module.unwrapKeyFn.fullName() + " (f x)) xs)",
		},
	}
}
