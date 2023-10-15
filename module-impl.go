package main

func (id identifier) importLine() string {
	if id.moduleName == "" {
		return ""
	}
	return "import " + id.moduleName
}
