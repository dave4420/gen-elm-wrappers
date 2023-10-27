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

func (module dictModule) keysDictDef() definition {
	return definition{
		localName: "keys",
		source: []string{
			"keys : " + module.wrapperType.name + " v -> List " + module.publicKeyType.fullName(),
			"keys (" + module.wrapperType.name + " d) = Dict.keys d |> List.filterMap " + module.wrapKeyFn.fullName(),
		},
	}
}

func (module dictModule) valuesDictDef() definition {
	return definition{
		localName: "values",
		source: []string{
			"values : " + module.wrapperType.name + " v -> List v",
			"values (" + module.wrapperType.name + " d) = Dict.values d",
		},
	}
}

func (module dictModule) toListDictDef() definition {
	return definition{
		localName: "toList",
		source: []string{
			"toList : " + module.wrapperType.name + " v -> List (" + module.publicKeyType.fullName() + ", v)",
			"toList (" + module.wrapperType.name + " d) = Dict.toList d |> List.filterMap (\\(k, v) ->",
			"    case " + module.wrapKeyFn.fullName() + " k of",
			"      Just kk -> Just (kk, v)",
			"      Nothing -> Nothing",
			"  )",
		},
	}
}

func (module dictModule) fromListDictDef() definition {
	return definition{
		localName: "fromList",
		source: []string{
			"fromList : List (" + module.publicKeyType.fullName() + ", v) -> " + module.wrapperType.name + " v",
			"fromList l = " + module.wrapperType.name + " (Dict.fromList (List.map (Tuple.mapFirst " + module.unwrapKeyFn.fullName() + ") l))",
		},
	}
}

func (module dictModule) mapDictDef() definition {
	return definition{
		localName: "map",
		source: []string{
			"map : (" + module.publicKeyType.name + " -> v -> w) -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " w",
			"map f d = ",
			"  let",
			"    g k v dd = insert k (f k v) dd",
			"  in foldl g empty d",
		},
	}
}

func (module dictModule) foldlDictDef() definition {
	return definition{
		localName: "foldl",
		source: []string{
			"foldl : (" + module.publicKeyType.name + " -> v -> a -> a) -> a -> " + module.wrapperType.name + " v -> a",
			"foldl f z (" + module.wrapperType.name + " d) = ",
			"  let",
			"    g k v acc = case " + module.wrapKeyFn.fullName() + " k of",
			"      Nothing -> acc",
			"      Just kk -> f kk v acc",
			"  in Dict.foldl g z d",
		},
	}
}

func (module dictModule) foldrDictDef() definition {
	return definition{
		localName: "foldr",
		source: []string{
			"foldr : (" + module.publicKeyType.name + " -> v -> a -> a) -> a -> " + module.wrapperType.name + " v -> a",
			"foldr f z (" + module.wrapperType.name + " d) = ",
			"  let",
			"    g k v acc = case " + module.wrapKeyFn.fullName() + " k of",
			"      Nothing -> acc",
			"      Just kk -> f kk v acc",
			"  in Dict.foldr g z d",
		},
	}
}

func (module dictModule) filterDictDef() definition {
	return definition{
		localName: "filter",
		source: []string{
			"filter : (" + module.publicKeyType.name + " -> v -> Bool) -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v",
			"filter f (" + module.wrapperType.name + " d) = ",
			"  let",
			"    g k v =",
			"      " + module.wrapKeyFn.fullName() + " k",
			"        |> Maybe.map (\\kk -> f kk v)",
			"        |> Maybe.withDefault False",
			"  in",
			"    " + module.wrapperType.name + " (Dict.filter g d)",
		},
	}
}

func (module dictModule) partitionDictDef() definition {
	return definition{
		localName: "partition",
		source: []string{
			"partition : (" + module.publicKeyType.name + " -> v -> Bool) -> " + module.wrapperType.name + " v -> (" + module.wrapperType.name + " v, " + module.wrapperType.name + " v)",
			"partition f (" + module.wrapperType.name + " d) = ",
			"  let",
			"    g k v =",
			"      " + module.wrapKeyFn.fullName() + " k",
			"        |> Maybe.map (\\kk -> f kk v)",
			"        |> Maybe.withDefault False",
			"  in",
			"    Tuple.mapBoth " + module.wrapperType.name + " " + module.wrapperType.name + " (Dict.partition g d)",
		},
	}
}

func (module dictModule) unionDictDef() definition {
	return definition{
		localName: "union",
		source: []string{
			"union : " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v",
			"union (" + module.wrapperType.name + " d1) (" + module.wrapperType.name + " d2) = " + module.wrapperType.name + " (Dict.union d1 d2)",
		},
	}
}

func (module dictModule) intersectDictDef() definition {
	return definition{
		localName: "intersect",
		source: []string{
			"intersect : " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v",
			"intersect (" + module.wrapperType.name + " d1) (" + module.wrapperType.name + " d2) = " + module.wrapperType.name + " (Dict.intersect d1 d2)",
		},
	}
}

func (module dictModule) diffDictDef() definition {
	return definition{
		localName: "diff",
		source: []string{
			"diff : " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v -> " + module.wrapperType.name + " v",
			"diff (" + module.wrapperType.name + " d1) (" + module.wrapperType.name + " d2) = " + module.wrapperType.name + " (Dict.diff d1 d2)",
		},
	}
}

func (module dictModule) mergeDictDef() definition {
	return definition{
		localName: "merge",
		source: []string{
			"merge :",
			"  (" + module.publicKeyType.fullName() + " -> a -> result -> result)",
			"  -> (" + module.publicKeyType.fullName() + " -> a -> b -> result -> result)",
			"  -> (" + module.publicKeyType.fullName() + " -> b -> result -> result)",
			"  -> " + module.wrapperType.name + " a",
			"  -> " + module.wrapperType.name + " b",
			"  -> result",
			"  -> result",
			"merge f1 f2 f3 (" + module.wrapperType.name + " d1) (" + module.wrapperType.name + " d2) z =",
			"  let",
			"    wrap f k v = case " + module.wrapKeyFn.fullName() + " k of",
			"      Just kk -> f kk v",
			"      Nothing -> identity",
			"    wrap2 f k v w = case " + module.wrapKeyFn.fullName() + " k of",
			"      Just kk -> f kk v w",
			"      Nothing -> identity",
			"  in",
			"  Dict.merge (wrap f1) (wrap2 f2) (wrap f3) d1 d2 z",
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
		module.keysDictDef(),
		module.valuesDictDef(),
		module.toListDictDef(),
		module.fromListDictDef(),
		module.mapDictDef(),
		module.foldlDictDef(),
		module.foldrDictDef(),
		module.filterDictDef(),
		module.partitionDictDef(),
		module.unionDictDef(),
		module.intersectDictDef(),
		module.diffDictDef(),
		module.mergeDictDef(),
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
