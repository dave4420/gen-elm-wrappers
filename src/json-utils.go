package main

import (
	"errors"
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
