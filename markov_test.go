package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetChainSimilarityScore(t *testing.T) {
	texts := []string{
		"That fish ate that fish.",
		"That cat ate that fish.",
	}

	chains := make([]map[string]map[string]int, len(texts))
	for i, s := range texts {
		chains[i] = createChain(s)
	}

	// Compare every text to every other text.
	results := map[string]string{}
	for i := range texts {
		for j := i + 1; j < len(texts); j++ {
			key := fmt.Sprintf("%d,%d", i, j)
			score, err := getChainSimilarityScore(chains[i], chains[j], scoreOptions{
				strategy:     Mean,
				topWordPairs: 0,
			})
			if err != nil {
				panic(err)
			}
			results[key] = fmt.Sprintf("%.2f", score)
		}
	}
	assert.NotEmpty(t, results)
}
