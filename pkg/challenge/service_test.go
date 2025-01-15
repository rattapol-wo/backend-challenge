package challenge_test

import (
	"backend-challenge/pkg/challenge"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNumberSummary(t *testing.T) {
	tests := []struct {
		name         string
		filePath     string
		expected     int
	}{
		{
			name:         "Simple Data",
			filePath:     "assert/simple.json",
			expected:     237,
		},
		{
			name:         "Hard Data",
			filePath:     "assert/hard.json",
			expected:     7273,
		},
	}

	// Loop ผ่าน test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data, err := ioutil.ReadFile(tc.filePath)
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}

			var body [][]int
			err = json.Unmarshal(data, &body)
			if err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			resp := challenge.SumValue(body)
			fmt.Println("resp sum: ", resp)
			assert.Equal(t, tc.expected, resp)
		})
	}
}
