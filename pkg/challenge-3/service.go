package challenge3

import (
	"backend-challenge/pkg/client"
	"fmt"
	"strings"
)

func CountMeat(meatCategories string) (map[string]int, error) {


    textBacon, err := client.CallBaconipsumClient(meatCategories)
	if err != nil {
		fmt.Println("Call baconipsum client failed: ", err)
		return nil, err
	}
	fmt.Println("textBacon : ", len(textBacon))


	// Normalize the text: convert to lower case and remove punctuation
	normalized := strings.ToLower(textBacon)
	normalized = strings.NewReplacer(",", "", ".", "", "  ", " ", "..", "").Replace(normalized)

	// Split the text into words
	words := strings.Fields(normalized)

	// Create a map to count occurrences
	meatCount := make(map[string]int)

	// Count each word
	for _, word := range words {
		meatCount[word]++
	}

	return meatCount, nil
}
