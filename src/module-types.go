package main

type identifier struct {
	moduleName string
	name       string
}

type module interface {
	name() string
	source(elmConfig elmConfig) ([]string, error)
}

type dictModule struct {
	wrapperType    identifier
	publicKeyType  identifier
	privateKeyType identifier
	wrapKeyFn      identifier
	unwrapKeyFn    identifier
}

type setModule struct {
	wrapperType    identifier
	publicKeyType  identifier
	privateKeyType identifier
	wrapKeyFn      identifier
	unwrapKeyFn    identifier
}
