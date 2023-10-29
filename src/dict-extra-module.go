package main

func (module dictModule) extraDefs() []definition {
	if module.dictExtraVersion == "" {
		return []definition{}
	}
	return []definition{
		module.groupByDef(),
		module.filterGroupByDef(),
		module.fromListByDef(),
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

func (module dictModule) filterGroupByDef() definition {
	return definition{
		localName: "filterGroupBy",
		source: []string{
			"filterGroupBy : (a -> Maybe " + module.publicKeyType.fullName() + ") -> List a -> " + module.wrapperType.name + " (List a)",
			"filterGroupBy f xs = " + module.wrapperType.name + " (Dict.Extra.filterGroupBy (\\x -> Maybe.map " + module.unwrapKeyFn.fullName() + " (f x)) xs)",
		},
	}
}

func (module dictModule) fromListByDef() definition {
	return definition{
		localName: "fromListBy",
		source: []string{
			"fromListBy : (a -> " + module.publicKeyType.fullName() + ") -> List a -> " + module.wrapperType.name + " a",
			"fromListBy f xs = " + module.wrapperType.name + " (Dict.Extra.fromListBy (\\x -> " + module.unwrapKeyFn.fullName() + " (f x)) xs)",
		},
	}
}
