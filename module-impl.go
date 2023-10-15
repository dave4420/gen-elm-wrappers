package main

func (id identifier) importLine() string {
	if id.moduleName == "" {
		return ""
	}
	return "import " + id.moduleName
}

func (id identifier) fullName() string {
	if id.moduleName == "" {
		return id.name
	}
	return id.moduleName + "." + id.name
}
