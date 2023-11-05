package main

import (
	"encoding/json"
	"errors"
	"os"
)

func decodeElmConfig(root interface{}) (elmConfig, error) {
	var ret elmConfig

	dependencies, err := getObjectProperty(root, "elm.json", "dependencies")
	if err != nil {
		return ret, err
	}

	direct, err := getObjectProperty(dependencies, "elm.json['dependencies']", "direct")
	if err != nil {
		return ret, err
	}

	directObject, ok := direct.(map[string]interface{})
	if !ok {
		return ret, errors.New("elm.json['dependencies']['direct'] is not an object")
	}

	elmCoreVersion, ok := directObject["elm/core"]
	if !ok {
		return ret, errors.New("elm.json['dependencies']['direct']['elm/core'] not found")
	}

	elmCoreVersionString, ok := elmCoreVersion.(string)
	if !ok {
		return ret, errors.New("elm.json['dependencies']['direct']['elm/core'] is not a string")
	}

	ret.elmCoreVersion, err = parseVersion(elmCoreVersionString)
	if err != nil {
		return ret, err
	}

	dictExtraVersion, ok := directObject["elm-community/dict-extra"]
	var dictExtraVersionString string
	if ok {
		dictExtraVersionString, ok = dictExtraVersion.(string)
	}
	if ok {
		dictExtraVersion, err := parseVersion(dictExtraVersionString)
		if err != nil {
			return ret, err
		}
		ret.dictExtraVersion = &dictExtraVersion
	}

	return ret, nil
}

func decodeElmConfigFromBlob(blob []byte) (elmConfig, error) {
	var root interface{}
	err := json.Unmarshal(blob, &root)
	if err != nil {
		return elmConfig{}, err
	}

	return decodeElmConfig(root)
}

func readElmConfig() (elmConfig, error) {
	elmJson, err := os.ReadFile("elm.json")
	if err != nil {
		return elmConfig{}, err
	}

	return decodeElmConfigFromBlob(elmJson)
}
