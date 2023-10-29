package main

import (
	"encoding/json"
	"errors"
	"os"
)

func decodeConfig(root interface{}) (config, error) {
	var ok bool

	var rootObject map[string]interface{}
	rootObject, ok = root.(map[string]interface{})
	if !ok {
		return config{}, errors.New("elm.json is not an object")
	}

	var dependencies interface{}
	dependencies, ok = rootObject["dependencies"]
	if !ok {
		return config{}, errors.New("elm.json['dependencies'] not found")
	}

	var dependenciesObject map[string]interface{}
	dependenciesObject, ok = dependencies.(map[string]interface{})
	if !ok {
		return config{}, errors.New("elm.json['dependencies'] is not an object")
	}

	var direct interface{}
	direct, ok = dependenciesObject["direct"]
	if !ok {
		return config{}, errors.New("elm.json['dependencies']['direct'] not found")
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
	dictExtraVersion, ok = directObject["elm-community/dict-extra"]
	if !ok {
		return config{}, errors.New("elm.json['dependencies']['direct']['elm-community/dict-extra'] not found")
	}

	var dictExtraVersionString string
	dictExtraVersionString, ok = dictExtraVersion.(string)
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
			},
		},
		elmCoreVersion:   elmCoreVersionString,
		dictExtraVersion: dictExtraVersionString,
	}, nil
}

func readConfig() (config, error) {
	blob, err := os.ReadFile("elm.json")
	if err != nil {
		return config{}, err
	}

	var root interface{}
	err = json.Unmarshal(blob, &root)
	if err != nil {
		return config{}, err
	}

	return decodeConfig(root)
}
