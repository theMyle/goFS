package internal

import (
	"testing"
)

func TestGetExtension(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Good file name", "sample.txt", "txt"},
		{"Bad file name", "ZigWag", ""},
		{"No file name", "", ""},
		{"File with dot in front", ".gitignore", ""},
		{"File with invalid extension", "noype.", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getExtension(tc.input)
			if tc.expected == "" {
				if result != "" {
					t.Errorf("geExtension(%s) = %v, expected = %v", tc.input, result, tc.expected)
				}
			} else {
				if result == "" {
					t.Errorf("geExtension(%s) = %v, expected = %v", tc.input, result, tc.expected)
				}
			}
		})
	}
}

func TestGetParentDir(t *testing.T) {
}
