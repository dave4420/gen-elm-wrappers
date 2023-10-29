package main

type identifier struct {
	moduleName string
	name       string
}

type module interface {
	name() string
	source() []string
}

type dictModule struct {
	wrapperType      identifier
	publicKeyType    identifier
	privateKeyType   identifier
	wrapKeyFn        identifier
	unwrapKeyFn      identifier
	elmCoreVersion   string
	dictExtraVersion string
}
