package auth

import (
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "simple string",
			input:    "test",
			expected: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		},
		{
			name:     "special characters",
			input:    "!@#$%^&*()",
			expected: "95ce789c5c9d18490972709838ca3a9719094bca3ac16332cfec0652b0236141",
		},
		{
			name:     "long string",
			input:    "ThisIsAVeryLongStringThatShouldStillBeHashedCorrectly",
			expected: "ab137642a8e691b75698bef9c19e3841738b834f330d8e11f0b8c6a8a225fee6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Hash(tt.input)
			if result != tt.expected {
				t.Errorf("Hash(%s) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
func TestCompareHash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		hashed   string
		expected bool
	}{
		{
			name:     "matching strings",
			input:    "password123",
			hashed:   "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
			expected: true,
		},
		{
			name:     "non-matching strings",
			input:    "password123",
			hashed:   "wronghash",
			expected: false,
		},
		{
			name:     "empty input vs non-empty hash",
			input:    "",
			hashed:   "somehashedvalue",
			expected: false,
		},
		{
			name:     "both empty strings",
			input:    "",
			hashed:   "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			expected: true,
		},
		{
			name:     "case sensitivity check",
			input:    "Password123",
			hashed:   "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
			expected: false,
		},
		{
			name:     "unicode characters",
			input:    "пароль123",
			hashed:   "74948ae38c64cef3291ac02c845fbe58dfa2e0fbed7ac384926502d6001806c2",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareHash(tt.input, tt.hashed)
			if result != tt.expected {
				t.Errorf("CompareHash(%s, %s) = %v, want %v", tt.input, tt.hashed, result, tt.expected)
			}
		})
	}
}