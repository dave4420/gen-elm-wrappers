package main

import "testing"

func TestDecodesDictModuleConfig(t *testing.T) {
	// given
	input := `{
		"dependencies": {
			"direct": {
				"elm/core": "1.0.5",
				"elm-community/dict-extra": "2.4.0"
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
		dictExtraVersion: "2.4.0",
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

func TestDecodesDictModuleConfigWithoutDictExtra(t *testing.T) {
	// given
	input := `{
		"dependencies": {
			"direct": {
				"elm/core": "1.0.5"
			}
		}
	}`

	// when
	output, err := decodeConfigFromBlob([]byte(input))

	// then
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	module, ok := output.modules[0].(dictModule)
	if !ok {
		t.Errorf("Unexpected output: %v", output)
	}
	if module.dictExtraVersion != "" {
		t.Errorf("Unexpected dict-extra version: %v", module.dictExtraVersion)
	}
}

func TestDoesntDecodeDictModuleConfigWithoutElmCore(t *testing.T) {
	// given
	input := `{
		"dependencies": {
			"direct": {
				"elm-community/dict-extra": "2.4.0"
			}
		}
	}`

	// when
	output, err := decodeConfigFromBlob([]byte(input))

	// then
	if err == nil {
		t.Errorf("Expected error, instead got output: %v", output)
	}
}
