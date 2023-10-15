package main

func (module dictModule) name() string {
	return module.typeId.moduleName
}

func (module dictModule) source() []string {
	return []string{
		"module " + module.typeId.moduleName + " exposing (..)",
		module.typeId.importLine(),
		module.publicKeyId.importLine(),
		module.privateKeyId.importLine(),
	}
}
