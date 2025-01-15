package challenge2_test

import (
	challenge2 "backend-challenge/pkg/challenge-2"
	"backend-challenge/pkg/utils"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestKeyboardEncoded(t *testing.T) {
	tests := []struct {
		name     string
		encoded  string
		expected string
	}{
		{
			name:     "Case 1",
			encoded:  "LLRR=",
			expected: "210122",
		},
		{
			name:     "Case 2",
			encoded:  "==RLL",
			expected: "000210",
		},
		{
			name:     "Case 3",
			encoded:  "=LLRR",
			expected: "221012",
		},
		{
			name:     "Case 4",
			encoded:  "RRL=R",
			expected: "012001",
		},
	}

	// Loop ผ่าน test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			resp := challenge2.ProcessKeyboardDecode(tc.encoded)

			utils.PrintStructJson(resp)
			assert.Equal(t, resp.Decoded, tc.expected)
		})
	}
}
