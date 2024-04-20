package main

import "errors"

func (module setModule) extraDefs(setExtraVersion *version) ([]definition, error) {
	if setExtraVersion == nil {
		return []definition{}, nil
	}
	supportedSetExtraVersion := version{major: 1, minor: 2}
	if setExtraVersion.isSameMajorAndNotEarlierMinorThan(supportedSetExtraVersion) {
		return []definition{
		}, nil
	}
	return []definition{}, errors.New("Versions " + setExtraVersion.toString() +
		" of stoeffel/set-extra are not supported, must be compatible with " +
		supportedSetExtraVersion.toString(),
	)
}

// DAVE: add actual wrappers
