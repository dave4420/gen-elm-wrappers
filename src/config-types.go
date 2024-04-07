package main

type elmConfig struct {
	elmCoreVersion   version
	dictExtraVersion *version
	setExtraVersion  *version
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
	if *x.dictExtraVersion != *y.dictExtraVersion {
		return false
	}
	if x.setExtraVersion == nil && y.setExtraVersion == nil {
		return true
	}
	if x.setExtraVersion == nil || y.setExtraVersion == nil {
		return false
	}
	if *x.setExtraVersion != *y.setExtraVersion {
		return false
	}
	return true
}

type config struct {
	path    string
	modules []module
}
