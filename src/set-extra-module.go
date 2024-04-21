package main

import "errors"

func (module setModule) extraDefs(setExtraVersion *version) ([]definition, error) {
	if setExtraVersion == nil {
		return []definition{}, nil
	}
	supportedSetExtraVersion := version{major: 1, minor: 2}
	if setExtraVersion.isSameMajorAndNotEarlierMinorThan(supportedSetExtraVersion) {
		return []definition{
			module.subsetDef(),
			module.toggleDef(),
		}, nil
	}
	return []definition{}, errors.New("Versions " + setExtraVersion.toString() +
		" of stoeffel/set-extra are not supported, must be compatible with " +
		supportedSetExtraVersion.toString(),
	)
}

// DAVE: handle variations of concatMap
// DAVE: handle variations of filterMap

func (module setModule) subsetDef() definition {
	return definition{
		localName: "subset",
		source: []string{
			"subset : " + module.wrapperType.name + " -> " + module.wrapperType.name + " -> Bool",
			"subset (" + module.wrapperType.name + " set1) (" + module.wrapperType.name + " set2) =",
			"    Set.Extra.subset set1 set2",
		},
	}
}

func (module setModule) toggleDef() definition {
	return definition{
		localName: "toggle",
		source: []string{
			"toggle : " + module.privateKeyType.fullName() + " -> " + module.wrapperType.name + " -> " + module.wrapperType.name,
			"toggle key (" + module.wrapperType.name + " set) =",
			"    " + module.wrapperType.name + " <| Set.Extra.toggle key set",
		},
	}
}
