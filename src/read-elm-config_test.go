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
	if !output.equals(expected) {
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
	if !output.equals(expected) {
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
