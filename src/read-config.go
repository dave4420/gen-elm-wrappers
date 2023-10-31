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

func decodeConfig(root interface{}) (config, error) {
	var ok bool
	var err error

	var dependencies interface{}
	dependencies, err = getObjectProperty(root, "elm.json", "dependencies")
	if err != nil {
		return config{}, err
	}

	var direct interface{}
	direct, err = getObjectProperty(dependencies, "elm.json['dependencies']", "direct")
	if err != nil {
		return config{}, err
	}

	var directObject map[string]interface{}
	directObject, ok = direct.(map[string]interface{})
	if !ok {
		return config{}, errors.New("elm.json['dependencies']['direct'] is not an object")
	}

	var elmCoreVersion interface{}
	elmCoreVersion, ok = directObject["elm/core"]
	if !ok {
		return config{}, errors.New("elm.json['dependencies']['direct']['elm/core'] not found")
	}

	var elmCoreVersionString string
	elmCoreVersionString, ok = elmCoreVersion.(string)
	if !ok {
		return config{}, errors.New("elm.json['dependencies']['direct']['elm/core'] is not a string")
	}

	var dictExtraVersion interface{}
	var dictExtraVersionString string
	dictExtraVersion, ok = directObject["elm-community/dict-extra"]
	if ok {
		dictExtraVersionString, ok = dictExtraVersion.(string)
	}
	if !ok {
		dictExtraVersionString = ""
	}

	var genElmWrappers interface{}
	genElmWrappers, err = getObjectProperty(root, "elm.json", "gen-elm-wrappers")
	if err != nil {
		return config{}, err
	}

	var generate interface{}
	generate, err = getObjectProperty(genElmWrappers, "elm.json['gen-elm-wrappers']", "generate")
	if err != nil {
		return config{}, err
	}

	var generateArray []interface{}
	generateArray, ok = generate.([]interface{})
	if !ok {
		return config{}, errors.New("elm.json['gen-elm-wrappers']['generate'] is not an array")
	}

	modules := []module{}
	for i, moduleJson := range generateArray {
		module := dictModule{
			elmCoreVersion:   elmCoreVersionString,
			dictExtraVersion: dictExtraVersionString,
		}
		path := fmt.Sprintf("elm.json['gen-elm-wrappers']['generate'][%d]", i)

		module.wrapperType, err = getObjectPropertyIdentifier(moduleJson, path, "wrapper-type")
		if err != nil {
			return config{}, err
		}
		if module.wrapperType.moduleName == "" {
			return config{}, errors.New(path + "['wrapperType'] is missing a module name")
		}

		module.publicKeyType, err = getObjectPropertyIdentifier(moduleJson, path, "public-key-type")
		if err != nil {
			return config{}, err
		}

		module.privateKeyType, err = getObjectPropertyIdentifier(moduleJson, path, "private-key-type")
		if err != nil {
			return config{}, err
		}

		module.wrapKeyFn, err = getObjectPropertyIdentifier(moduleJson, path, "private-key-to-public-key")
		if err != nil {
			return config{}, err
		}

		module.unwrapKeyFn, err = getObjectPropertyIdentifier(moduleJson, path, "public-key-to-private-key")
		if err != nil {
			return config{}, err
		}

		modules = append(modules, module)
	}

	return config{
		path:    "src",
		modules: modules,
	}, nil
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
	blob, err := os.ReadFile("elm.json")
	if err != nil {
		return config{}, err
	}

	return decodeConfigFromBlob(blob)
}
