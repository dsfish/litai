package main

import (
	"fmt"
	"testing"
)

func TestGetChainSimilarityScore(t *testing.T) {
	texts := []string{
		"the quick brown fox jumped over the lazy dog but the cat jumped over nothing",
		"the big brown fox jumped over the questionable dog and the cat jumped over the moon",
		"there's always a moose around the bend",
		"there's never a good reason to give a moose a muffin",
		"underwater porcupines are a figment of your imagination",
		"the porcupines are the lazy friends we all need to respect",
	}

	chains := make([]map[string]map[string]int, len(texts))
	for i, s := range texts {
		chains[i] = createChain(s)
	}

	// Compare every texts to every other texts using every strategy.
	results := map[string]map[string]string{}
	for i := range texts {
		for j := i + 1; j < len(texts); j++ {
			key := fmt.Sprintf("%d,%d", i, j)
			results[key] = map[string]string{}
			for _, strategy := range strategies {
				score, err := getChainSimilarityScore(chains[i], chains[j], strategy)
				if err != nil {
					panic(err)
				}
				results[key][strategy.String()] = fmt.Sprintf("%.2f", score)
			}
		}
	}

	prettyPrint(results)
}
