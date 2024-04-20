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

// DAVE: add actual wrappers
