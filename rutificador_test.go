package rutificador

import "testing"

func TestObtenerDV(t *testing.T) {
	testCases := []struct {
		input    string
		expected rune
	}{
		{"18622178", '8'},
		{"11111111", '1'},
		{"70360100", '6'},
		{"22222222", '2'},
		{"15304340", 'K'},
		{"8367720", '1'},
	}

	for _, tc := range testCases {
		if result := obtenerDV(tc.input); result != tc.expected {
			t.Errorf("obtenerDV(%s) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}
