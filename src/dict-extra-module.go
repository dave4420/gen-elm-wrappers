package main

func (module dictModule) extraDefs() []definition {
	if module.dictExtraVersion == "" {
		return []definition{}
	}
	return []definition{
		module.groupByDef(),
		module.filterGroupByDef(),
		module.fromListByDef(),
		module.fromListDedupeDef(),
		module.fromListDedupeByDef(),
		module.frequenciesDef(),
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

func (module dictModule) fromListDedupeDef() definition {
	return definition{
		localName: "fromListDedupe",
		source: []string{
			"fromListDedupe : (a -> a -> a) -> List (" + module.publicKeyType.fullName() + ", a) -> " + module.wrapperType.name + " a",
			"fromListDedupe f xs = " + module.wrapperType.name + " (Dict.Extra.fromListDedupe f (List.map (Tuple.mapFirst " + module.unwrapKeyFn.fullName() + ") xs))",
		},
	}
}

func (module dictModule) fromListDedupeByDef() definition {
	return definition{
		localName: "fromListDedupeBy",
		source: []string{
			"fromListDedupeBy : (a -> a -> a) -> (a -> " + module.publicKeyType.fullName() + ") -> List a -> " + module.wrapperType.name + " a",
			"fromListDedupeBy f g xs = " + module.wrapperType.name + " (Dict.Extra.fromListDedupeBy f (g >> " + module.unwrapKeyFn.fullName() + ") xs)",
		},
	}
}

func (module dictModule) frequenciesDef() definition {
	return definition{
		localName: "frequencies",
		source: []string{
			"frequencies : List " + module.publicKeyType.fullName() + " -> " + module.wrapperType.name + " Int",
			"frequencies = List.map " + module.unwrapKeyFn.fullName() + " >> Dict.Extra.frequencies >> " + module.wrapperType.name,
		},
	}
}
