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

// DAVE: add actual wrappers
