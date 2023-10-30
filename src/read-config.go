package main

import (
	"encoding/json"
	"errors"
	"os"
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

	return config{
		path: "src",
		modules: []module{
			dictModule{
				wrapperType: identifier{
					moduleName: "Type.DictInt",
					name:       "DictInt",
				},
				publicKeyType: identifier{
					name: "Int",
				},
				privateKeyType: identifier{
					name: "String",
				},
				wrapKeyFn: identifier{
					moduleName: "String",
					name:       "toInt",
				},
				unwrapKeyFn: identifier{
					moduleName: "String",
					name:       "fromInt",
				},
				elmCoreVersion:   elmCoreVersionString,
				dictExtraVersion: dictExtraVersionString,
			},
		},
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
