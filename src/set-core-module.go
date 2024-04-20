package main

import (
	"errors"
	"strings"
)

func (module setModule) coreDefs(elmCoreVersion version) ([]definition, error) {
	supportedElmCoreVersion := version{major: 1, minor: 0}
	if elmCoreVersion.isSameMajorAndNotEarlierMinorThan(supportedElmCoreVersion) {
		return []definition{
			module.setDef(),
			module.emptyDef(),
			module.singletonDef(),
			module.insertDef(),
			module.removeDef(),
			module.isEmptyDef(),
			module.memberDef(),
			module.sizeDef(),
			module.unionDef(),
			module.intersectDef(),
			module.diffDef(),
			module.toListDef(),
			module.fromListDef(),
		}, nil
	}
	return []definition{}, errors.New("Versions " + elmCoreVersion.toString() + " of elm/core " +
		"are not supported, must be compatible with " + supportedElmCoreVersion.toString(),
	)
}

func (module setModule) setDef() definition {
	return definition{
		localName: module.wrapperType.name,
		source: []string{
			strings.Join(
				[]string{
					"type",
					module.wrapperType.name,
					"=",
					module.wrapperType.name,
					"(Set",
					module.privateKeyType.fullName(),
					")",
				},
				" ",
			),
		},
	}
}

// Build

func (module setModule) emptyDef() definition {
	return definition{
		localName: "empty",
		source: []string{
			"empty : " + module.wrapperType.name,
			"empty = " + module.wrapperType.name + " Set.empty",
		},
	}
}

func (module setModule) singletonDef() definition {
	return definition{
		localName: "singleton",
		source: []string{
			"singleton : " + module.publicKeyType.fullName() + " -> " + module.wrapperType.name,
			"singleton k = " + module.wrapperType.name + " (Set.singleton (" + module.unwrapKeyFn.fullName() + " k))",
		},
	}
}

func (module setModule) insertDef() definition {
	return definition{
		localName: "insert",
		source: []string{
			"insert : " + module.publicKeyType.fullName() + " -> " + module.wrapperType.name + " -> " + module.wrapperType.name,
			"insert k (" + module.wrapperType.name + " d) = " + module.wrapperType.name + " (Set.insert (" + module.unwrapKeyFn.fullName() + " k) d)",
		},
	}
}

func (module setModule) removeDef() definition {
	return definition{
		localName: "remove",
		source: []string{
			"remove : " + module.publicKeyType.fullName() + " -> " + module.wrapperType.name + " -> " + module.wrapperType.name,
			"remove k (" + module.wrapperType.name + " d) = " + module.wrapperType.name + " (Set.remove (" + module.unwrapKeyFn.fullName() + " k) d)",
		},
	}
}

// Query

func (module setModule) isEmptyDef() definition {
	return definition{
		localName: "isEmpty",
		source: []string{
			"isEmpty : " + module.wrapperType.name + " -> Bool",
			"isEmpty (" + module.wrapperType.name + " d) = Set.isEmpty d",
		},
	}
}

func (module setModule) memberDef() definition {
	return definition{
		localName: "member",
		source: []string{
			"member : " + module.publicKeyType.fullName() + " -> " + module.wrapperType.name + " -> Bool",
			"member k (" + module.wrapperType.name + " d) = Set.member (" + module.unwrapKeyFn.fullName() + " k) d",
		},
	}
}

func (module setModule) sizeDef() definition {
	return definition{
		localName: "size",
		source: []string{
			"size : " + module.wrapperType.name + " -> Int",
			"size (" + module.wrapperType.name + " d) = Set.size d",
		},
	}
}

// Combine

func (module setModule) unionDef() definition {
	return definition{
		localName: "union",
		source: []string{
			"union : " + module.wrapperType.name + " -> " + module.wrapperType.name + " -> " + module.wrapperType.name,
			"union (" + module.wrapperType.name + " d1) (" + module.wrapperType.name + " d2) = " + module.wrapperType.name + " (Set.union d1 d2)",
		},
	}
}

func (module setModule) intersectDef() definition {
	return definition{
		localName: "intersect",
		source: []string{
			"intersect : " + module.wrapperType.name + " -> " + module.wrapperType.name + " -> " + module.wrapperType.name,
			"intersect (" + module.wrapperType.name + " d1) (" + module.wrapperType.name + " d2) = " + module.wrapperType.name + " (Set.intersect d1 d2)",
		},
	}
}

func (module setModule) diffDef() definition {
	return definition{
		localName: "diff",
		source: []string{
			"diff : " + module.wrapperType.name + " -> " + module.wrapperType.name + " -> " + module.wrapperType.name,
			"diff (" + module.wrapperType.name + " d1) (" + module.wrapperType.name + " d2) = " + module.wrapperType.name + " (Set.diff d1 d2)",
		},
	}
}

// Lists

func (module setModule) toListDef() definition {
	return definition{
		localName: "toList",
		source: []string{
			"toList : " + module.wrapperType.name + " -> List " + module.publicKeyType.fullName(),
			"toList (" + module.wrapperType.name + " d) = Set.toList d |> List.filterMap " + module.wrapKeyFn.fullName(),
		},
	}
}

func (module setModule) fromListDef() definition {
	return definition{
		localName: "fromList",
		source: []string{
			"fromList : List " + module.publicKeyType.fullName() + " -> " + module.wrapperType.name,
			"fromList l = " + module.wrapperType.name + " (Set.fromList (List.map " + module.unwrapKeyFn.fullName() + " l))",
		},
	}
}

// DAVE: add actual wrappers
