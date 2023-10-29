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
		module.removeWhenDef(),
		module.insertDedupeDef(),
		module.filterMapDef(),
		module.anyDef(),
		module.findDef(),
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

// Manipulation

func (module dictModule) removeWhenDef() definition {
	return definition{
		localName: "removeWhen",
		source: []string{
			"removeWhen : (" + module.publicKeyType.fullName() + " -> v -> Bool) -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v",
			"removeWhen f (" + module.wrapperType.name + " d) =",
			"  let",
			"    g k v = case " + module.wrapKeyFn.fullName() + " k of",
			"      Nothing -> True",
			"      Just kk -> f kk v",
			"  in " + module.wrapperType.name + " (Dict.Extra.removeWhen g d)",
		},
	}
}

// DAVE: add removeMany when Set module generated
// DAVE: add keepOnly when Set module generated

func (module dictModule) insertDedupeDef() definition {
	return definition{
		localName: "insertDedupe",
		source: []string{
			"insertDedupe : (v -> v -> v) -> " + module.publicKeyType.fullName() + " -> v -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v",
			"insertDedupe f k v (" + module.wrapperType.name + " d) =",
			"  " + module.wrapperType.name + " <| Dict.Extra.insertDedupe f (" + module.unwrapKeyFn.fullName() + " k) v d",
		},
	}
}

// DAVE: handle variations of mapKeys

func (module dictModule) filterMapDef() definition {
	return definition{
		localName: "filterMap",
		source: []string{
			"filterMap : (" + module.publicKeyType.fullName() + " -> a -> Maybe b) -> " + module.wrapperType.name + " a -> " + module.wrapperType.name + " b",
			"filterMap f (" + module.wrapperType.name + " d) =",
			"  let",
			"    g k v = case " + module.wrapKeyFn.fullName() + " k of",
			"      Nothing -> Nothing",
			"      Just kk -> f kk v",
			"  in",
			"  " + module.wrapperType.name + " <| Dict.Extra.filterMap g d",
		},
	}
}

// DAVE: handle variations of invert

// Utilities

func (module dictModule) anyDef() definition {
	return definition{
		localName: "any",
		source: []string{
			"any : (" + module.publicKeyType.fullName() + " -> v -> Bool) -> " + module.wrapperType.name + " v -> Bool",
			"any f (" + module.wrapperType.name + " d) =",
			"  let",
			"    g k v = case " + module.wrapKeyFn.fullName() + " k of",
			"      Nothing -> False",
			"      Just kk -> f kk v",
			"  in",
			"  Dict.Extra.any g d",
		},
	}
}

func (module dictModule) findDef() definition {
	return definition{
		localName: "find",
		source: []string{
			"find : (" + module.publicKeyType.fullName() + " -> v -> Bool) -> " + module.wrapperType.name + " v -> Maybe (" + module.publicKeyType.fullName() + ", v)",
			"find f (" + module.wrapperType.name + " d) =",
			"  let",
			"    g kk v = case " + module.wrapKeyFn.fullName() + " kk of",
			"      Nothing -> False",
			"      Just k -> f k v",
			"    h (kk, v) = case " + module.wrapKeyFn.fullName() + " kk of",
			"      Nothing -> Nothing",
			"      Just k -> Just (k, v)",
			"  in",
			"  Dict.Extra.find g d |> Maybe.andThen h",
		},
	}
}
