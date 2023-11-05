package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func getObjectProperty(json interface{}, path string, propertyName string) (interface{}, error) {
	var ok bool

	var object map[string]interface{}
	object, ok = json.(map[string]interface{})
	if !ok {
		return nil, errors.New(path + " is not an object")
	}

	var property interface{}
	property, ok = object[propertyName]
	if !ok {
		return nil, errors.New(path + "['" + propertyName + "'] not found")
	}

	return property, nil
}

func getObjectPropertyIdentifier(json interface{}, path string, propertyName string) (identifier, error) {
	id := identifier{}

	property, err := getObjectProperty(json, path, propertyName)
	if err != nil {
		return id, err
	}

	string, ok := property.(string)
	if !ok {
		return id, errors.New(path + "['" + propertyName + "'] is not a string")
	}

	ixFinalDot := strings.LastIndex(string, ".")
	if ixFinalDot == -1 {
		id.name = string
	} else {
		id.moduleName = string[:ixFinalDot]
		id.name = string[ixFinalDot+1:]
	}

	return id, nil
}

func decodeModule(moduleJson interface{}, elmConfig elmConfig, path string) (module, error) {
	// DAVE: rm elmConfig param
	underlyingType, err := getObjectProperty(moduleJson, path, "underlying-type")
	if err != nil {
		return nil, err
	}

	underlyingTypeString, ok := underlyingType.(string)
	if !ok {
		return nil, errors.New(path + "['underlying-type'] is not a string")
	}
	if underlyingTypeString != "Dict" {
		return nil, errors.New(path + "is not wrapping a 'Dict'")
	}

	module := dictModule{}

	module.wrapperType, err = getObjectPropertyIdentifier(moduleJson, path, "wrapper-type")
	if err != nil {
		return module, err
	}
	if module.wrapperType.moduleName == "" {
		return module, errors.New(path + "['wrapperType'] is missing a module name")
	}

	module.publicKeyType, err = getObjectPropertyIdentifier(moduleJson, path, "public-key-type")
	if err != nil {
		return module, err
	}

	module.privateKeyType, err = getObjectPropertyIdentifier(moduleJson, path, "private-key-type")
	if err != nil {
		return module, err
	}

	module.wrapKeyFn, err = getObjectPropertyIdentifier(moduleJson, path, "private-key-to-public-key")
	if err != nil {
		return module, err
	}

	module.unwrapKeyFn, err = getObjectPropertyIdentifier(moduleJson, path, "public-key-to-private-key")
	return module, err
}

func decodeConfig(root interface{}, elmConfig elmConfig) (config, error) {
	config := config{
		path: "src",
	}

	generate, err := getObjectProperty(root, "gen-elm-wrappers.json", "generate")
	if err != nil {
		return config, err
	}

	generateArray, ok := generate.([]interface{})
	if !ok {
		return config, errors.New("gen-elm-wrappers.json['generate'] is not an array")
	}

	for i, moduleJson := range generateArray {
		module, err := decodeModule(
			moduleJson,
			elmConfig,
			fmt.Sprintf("gen-elm-wrappers.json['generate'][%d]", i),
		)
		if err != nil {
			return config, err
		}

		config.modules = append(config.modules, module)
	}

	return config, nil
}

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

func decodeConfigFromBlob(blob []byte, elmConfig elmConfig) (config, error) {
	var root interface{}
	err := json.Unmarshal(blob, &root)
	if err != nil {
		return config{}, err
	}

	return decodeConfig(root, elmConfig)
}

func decodeElmConfigFromBlob(blob []byte) (elmConfig, error) {
	var root interface{}
	err := json.Unmarshal(blob, &root)
	if err != nil {
		return elmConfig{}, err
	}

	return decodeElmConfig(root)
}

func readConfig() (config, elmConfig, error) {
	// DAVE: split into two files
	elmJson, err := os.ReadFile("elm.json")
	if err != nil {
		return config{}, elmConfig{}, err
	}

	elmConfig, err := decodeElmConfigFromBlob(elmJson)
	if err != nil {
		return config{}, elmConfig, err
	}

	configJson, err := os.ReadFile("gen-elm-wrappers.json")
	if err != nil {
		return config{}, elmConfig, err
	}

	config, err := decodeConfigFromBlob(configJson, elmConfig)
	return config, elmConfig, err
}
