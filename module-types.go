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
	typeId       identifier
	publicKeyId  identifier
	privateKeyId identifier
}
