package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	inHeader := make(http.Header)
	inHeader.Set("Authorization", "ApiKey TestKey")
	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"correct": {input: inHeader, want: "TestKey"},
	}

	for name, tc := range tests {
		got1, got2 := GetAPIKey(tc.input)
		if tc.want != got1 && got2 == nil {
			t.Fatalf("%s: expected: %v - nil, got: %v - %v", name, tc.want, got1, got2)
		}
	}

}
