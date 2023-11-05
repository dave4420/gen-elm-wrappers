package main

import "testing"

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
