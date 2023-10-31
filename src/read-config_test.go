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
		},
		"gen-elm-wrappers": {
			"generate": [
				{
					"underlying-type": "Dict",
					"wrapper-type": "Type.DictCabbage.DictCabbage",
					"public-key-type": "Type.Cabbage.Cabbage",
					"private-key-type": "String",
					"private-key-to-public-key": "Type.Cabbage.fromString",
					"public-key-to-private-key": "Type.Cabbage.toString"
				}
			]
		}
	}`

	expectedModule := dictModule{
		wrapperType: identifier{
			moduleName: "Type.DictCabbage",
			name:       "DictCabbage",
		},
		publicKeyType: identifier{
			moduleName: "Type.Cabbage",
			name:       "Cabbage",
		},
		privateKeyType: identifier{
			name: "String",
		},
		wrapKeyFn: identifier{
			moduleName: "Type.Cabbage",
			name:       "fromString",
		},
		unwrapKeyFn: identifier{
			moduleName: "Type.Cabbage",
			name:       "toString",
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
	if len(output.modules) != 1 || output.modules[0] != expectedModule {
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
		},
		"gen-elm-wrappers": {
			"generate": [
				{
					"underlying-type": "Dict",
					"wrapper-type": "Type.DictCabbage.DictCabbage",
					"public-key-type": "Type.Cabbage.Cabbage",
					"private-key-type": "String",
					"private-key-to-public-key": "Type.Cabbage.fromString",
					"public-key-to-private-key": "Type.Cabbage.toString"
				}
			]
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
		},
		"gen-elm-wrappers": {
			"generate": [
				{
					"underlying-type": "Dict",
					"wrapper-type": "Type.DictCabbage.DictCabbage",
					"public-key-type": "Type.Cabbage.Cabbage",
					"private-key-type": "String",
					"private-key-to-public-key": "Type.Cabbage.fromString",
					"public-key-to-private-key": "Type.Cabbage.toString"
				}
			]
		}
	}`

	// when
	output, err := decodeConfigFromBlob([]byte(input))

	// then
	if err == nil {
		t.Errorf("Expected error, instead got output: %v", output)
	}
}
