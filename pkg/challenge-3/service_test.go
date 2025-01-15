package challenge3_test

import (
	challenge3 "backend-challenge/pkg/challenge-3"
	"backend-challenge/pkg/utils"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBeefSummary(t *testing.T) {
	tests := []struct {
		name       string
		typeMeat 	string
	}{
		{
			name: "Pie file die test 1",
			typeMeat: "meat-and-filler",
		},
	}

	// Loop ผ่าน test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			resp, err := challenge3.CountMeat(tc.typeMeat)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			utils.PrintStructJson(resp)
			assert.Equal(t, err, nil)
		})
	}
}
