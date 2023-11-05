package main

import (
	"fmt"
	"strconv"
	"strings"
)

type version struct {
	major int
	minor int
}

func (v version) toString() string {
	return fmt.Sprintf("%d.%d.x", v.major, v.minor)
}

func parseVersion(s string) (version, error) {
	ret := version{}
	components := strings.Split(s, ".")
	if len(components) == 0 {
		return ret, fmt.Errorf("empty version string")
	}
	for i, c := range components {
		n, err := strconv.ParseInt(c, 10, 0)
		if err != nil {
			return ret, err
		}
		if i == 0 {
			ret.major = int(n)
		}
		if i == 1 {
			ret.minor = int(n)
		}
	}
	return ret, nil
}

func (x version) isSameMajorAndNotEarlierMinorThan(y version) bool {
	return x.major == y.major && x.minor >= y.minor
}
