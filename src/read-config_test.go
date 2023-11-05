package main

import "testing"

func TestDecodesElmConfig(t *testing.T) {
	// given
	input := `{
		"dependencies": {
			"direct": {
				"elm/core": "1.0.5",
				"elm-community/dict-extra": "2.4.0"
			}
		}
	}`

	expected := elmConfig{
		elmCoreVersion:   version{major: 1, minor: 0},
		dictExtraVersion: &version{major: 2, minor: 4},
	}

	// when
	output, err := decodeElmConfigFromBlob([]byte(input))

	// then
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if output != expected {
		t.Errorf("Unexpected output: %v", output)
	}
}

func TestDecodesElmConfigWithoutDictExtra(t *testing.T) {
	// given
	input := `{
		"dependencies": {
			"direct": {
				"elm/core": "1.0.5"
			}
		}
	}`

	expected := elmConfig{
		elmCoreVersion:   version{major: 1, minor: 0},
		dictExtraVersion: nil,
	}

	// when
	output, err := decodeElmConfigFromBlob([]byte(input))

	// then
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if output != expected {
		t.Errorf("Unexpected output: %v", output)
	}
}

func TestDoesntDecodeElmConfigWithoutElmCore(t *testing.T) {
	// given
	input := `{
		"dependencies": {
			"direct": {
				"elm-community/dict-extra": "2.4.0"
			}
		}
	}`

	// when
	output, err := decodeElmConfigFromBlob([]byte(input))

	// then
	if err == nil {
		t.Errorf("Expected error, instead got output: %v", output)
	}
}

func TestDecodesDictModuleConfig(t *testing.T) {
	// given
	input := `{
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
		elmCoreVersion:   version{major: 1, minor: 0},
		dictExtraVersion: &version{major: 2, minor: 4},
	}

	// when
	output, err := decodeConfigFromBlob(
		[]byte(input),
		elmConfig{
			elmCoreVersion:   version{major: 1, minor: 0},
			dictExtraVersion: &version{major: 2, minor: 4},
		},
	)

	// then
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(output.modules) != 1 || output.modules[0] != expectedModule {
		t.Errorf("Unexpected output: %v", output)
	}
}
