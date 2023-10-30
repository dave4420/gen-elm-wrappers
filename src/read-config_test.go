package main

import "testing"

func TestDecodesDictModuleConfig(t *testing.T) {
	// given
	input := `{
		"dependencies": {
			"direct": {
				"elm/core": "1.0.5"
			}
		}
	}`

	expectedModule := dictModule{
		wrapperType: identifier{
			moduleName: "Type.DictInt",
			name:       "DictInt",
		},
		publicKeyType: identifier{
			name: "Int",
		},
		privateKeyType: identifier{
			name: "String",
		},
		wrapKeyFn: identifier{
			moduleName: "String",
			name:       "toInt",
		},
		unwrapKeyFn: identifier{
			moduleName: "String",
			name:       "fromInt",
		},
		elmCoreVersion:   "1.0.5",
		dictExtraVersion: "",
	}

	// when
	output, err := decodeConfigFromBlob([]byte(input))

	// then
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if output.modules[0] != expectedModule {
		t.Errorf("Unexpected output: %v", output)
	}
}
