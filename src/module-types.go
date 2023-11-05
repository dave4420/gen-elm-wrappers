package main

type identifier struct {
	moduleName string
	name       string
}

type module interface {
	name() string
	source() ([]string, error)
}

type dictModule struct {
	wrapperType      identifier
	publicKeyType    identifier
	privateKeyType   identifier
	wrapKeyFn        identifier
	unwrapKeyFn      identifier
	elmCoreVersion   version
	dictExtraVersion *version
}
