package main

import "strings"

func (module setModule) name() string {
	return module.wrapperType.moduleName
}

func (module setModule) source(elmConfig elmConfig) ([]string, error) {
	definitions := []definition{}

	coreDefs, err := module.coreDefs(elmConfig.elmCoreVersion)
	if err != nil {
		return []string{}, err
	}
	definitions = append(definitions, coreDefs...)

	// DAVE: uncomment
	// extraDefs, err := module.extraDefs(elmConfig.dictExtraVersion)
	// if err != nil {
	// 	return []string{}, err
	// }
	// definitions = append(definitions, extraDefs...)

	exports := []string{}
	for _, export := range definitions {
		exports = append(exports, export.localName)
	}

	var setExtraImportLine string
	if elmConfig.setExtraVersion != nil {
		setExtraImportLine = "import Set.Extra"
	}

	lines := []string{
		"module " + module.wrapperType.moduleName + " exposing (" + strings.Join(exports, ", ") + ")",
		"import Set exposing (Set)",
		setExtraImportLine,
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
