package main

type elmConfig struct {
	elmCoreVersion   version
	dictExtraVersion *version
}

func (x elmConfig) equals(y elmConfig) bool {
	if x.elmCoreVersion != y.elmCoreVersion {
		return false
	}
	if x.dictExtraVersion == nil && y.dictExtraVersion == nil {
		return true
	}
	if x.dictExtraVersion == nil || y.dictExtraVersion == nil {
		return false
	}
	return *x.dictExtraVersion == *y.dictExtraVersion
}

type config struct {
	path    string
	modules []module
}
