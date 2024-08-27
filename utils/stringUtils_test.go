package utils

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"123456", "654321"},
		{"", ""},
		{"123456789", "987654321"},
		{"asdasd", "dsadsa"},
	}
	for _, tt := range tests {
		if got := Reverse(tt.input); got != tt.want {
			t.Errorf("Reverse(%s) = %s, want %s", tt.input, got, tt.want)
		}
	}
}
