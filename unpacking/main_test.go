package main

import "testing"

type Test struct {
	input    string
	expected string
	wantErr  bool
}

func TestUnpack(t *testing.T) {
	tests := GetMockTests()
	for _, tt := range tests {
		t.Run("test unpack string", func(t *testing.T) {
			result, err := UnpackString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error: %v, wanted error: %v", err, tt.wantErr)
			}
			if result != tt.expected {
				t.Errorf("error doesn't expected result. result: %v, expected: %v", result, tt.expected)
			}
		})
	}
}

func GetMockTests() []Test {
	return []Test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
			wantErr:  false,
		},
		{
			input:    "abcd",
			expected: "abcd",
			wantErr:  false,
		},
		{
			input:    "45",
			expected: "",
			wantErr:  true,
		},
		{
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			input:    "2abc",
			expected: "",
			wantErr:  true,
		},
		{
			input:    "а3б2в",
			expected: "аааббв",
			wantErr:  false,
		},
		{
			input:    `qwe\4\5`,
			expected: "qwe45",
			wantErr:  false,
		},
		{
			input:    `qwe\45`,
			expected: "qwe44444",
			wantErr:  false,
		},
		{
			input:    `\4\5`,
			expected: "45",
			wantErr:  false,
		},
	}
}