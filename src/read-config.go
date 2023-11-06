package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func decodeModule(moduleJson interface{}, path string) (module, error) {
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
		return module, errors.New(path + "['wrapperType'] should be a fully qualified type name, " +
			"but either the type name or the module name is missing")
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

func decodeConfig(root interface{}) (config, error) {
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
			fmt.Sprintf("gen-elm-wrappers.json['generate'][%d]", i),
		)
		if err != nil {
			return config, err
		}

		config.modules = append(config.modules, module)
	}

	return config, nil
}

func decodeConfigFromBlob(blob []byte) (config, error) {
	var root interface{}
	err := json.Unmarshal(blob, &root)
	if err != nil {
		return config{}, err
	}

	return decodeConfig(root)
}

func readConfig() (config, error) {
	configJson, err := os.ReadFile("gen-elm-wrappers.json")
	if err != nil {
		return config{}, err
	}

	return decodeConfigFromBlob(configJson)
}
