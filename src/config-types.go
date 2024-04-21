package main

type elmConfig struct {
	elmCoreVersion   version
	dictExtraVersion *version
	setExtraVersion  *version
}

func (x elmConfig) equals(y elmConfig) bool {
	return x.elmCoreVersion == y.elmCoreVersion &&
		versionPtrEquals(x.dictExtraVersion, y.dictExtraVersion) &&
		versionPtrEquals(x.setExtraVersion, y.setExtraVersion)
}

type config struct {
	path    string
	modules []module
}
