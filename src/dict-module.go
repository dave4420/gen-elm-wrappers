package main

import "strings"

func (module dictModule) name() string {
	return module.wrapperType.moduleName
}

func (module dictModule) source() ([]string, error) {
	definitions := []definition{}

	coreDefs, err := module.coreDefs()
	if err != nil {
		return []string{}, err
	}
	definitions = append(definitions, coreDefs...)

	extraDefs, err := module.extraDefs()
	if err != nil {
		return []string{}, err
	}
	definitions = append(definitions, extraDefs...)

	exports := []string{}
	for _, export := range definitions {
		exports = append(exports, export.localName)
	}

	var dictExtraImportLine string
	if module.dictExtraVersion != nil {
		dictExtraImportLine = "import Dict.Extra"
	}

	lines := []string{
		"module " + module.wrapperType.moduleName + " exposing (" + strings.Join(exports, ", ") + ")",
		"import Dict exposing (Dict)",
		dictExtraImportLine,
		module.publicKeyType.importLine(),
		module.privateKeyType.importLine(),
		module.wrapKeyFn.importLine(),
		module.unwrapKeyFn.importLine(),
	}

	for _, def := range definitions {
		lines = append(lines, def.source...)
	}

	return lines, nil
}
