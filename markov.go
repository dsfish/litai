package main

import (
	"errors"
	"fmt"
	"github.com/montanaflynn/stats"
	"regexp"
	"strings"
)

// createChain returns a Markov chain in the following form:
//
// { the : { quick: 1, lazy: 1 }, lazy: { dog: 1 } }
func createChain(input string) map[string]map[string]int {
	chain := map[string]map[string]int{}

	// Remove special characters, excluding hyphens and apostrophes.
	reg := regexp.MustCompile("[^a-zA-Z0-9\\-\\'\\s]+")
	input = reg.ReplaceAllString(input, "")
	input = strings.ToLower(input)

	words := strings.Split(input, " ")

	var prevWord string
	for _, word := range words {
		if prevWord != "" {
			if chain[prevWord] == nil {
				chain[prevWord] = map[string]int{}
			}
			chain[prevWord][word] += 1
		}
		prevWord = word
	}
	return chain
}

// getChainSimilarityScore returns a decimal number in the range [0, 1]
// representing how similar the chains are, with 0 being no overlap and 1 being
// identical, according to the given strategy.
func getChainSimilarityScore(
	chainA, chainB map[string]map[string]int,
	strategy Strategy,
) (float64, error) {
	// For every word pair in chainA, record the number of times that the word pair
	// occurs in each chain, storing the smaller value first.
	scoringInput := map[string][]int{}
	for fromWord, toWords := range chainA {
		for toWord, aN := range toWords {
			key := fmt.Sprintf("%s~%s", fromWord, toWord)
			var bN int
			if chainB[fromWord] != nil {
				bN = chainB[fromWord][toWord]
			}
			if aN < bN {
				scoringInput[key] = []int{aN, bN}
			} else {
				scoringInput[key] = []int{bN, aN}
			}
		}
	}

	// Repeat the process for word pairs that only appear in chainB.
	for fromWord, toWords := range chainB {
		for toWord, bN := range toWords {
			// Ignore word pairs we saw when going through chainA.
			if chainA[fromWord] != nil && chainA[fromWord][toWord] != 0 {
				continue
			}

			key := fmt.Sprintf("%s~%s", fromWord, toWord)
			scoringInput[key] = []int{0, bN}
		}
	}

	var subScores []float64
	for _, input := range scoringInput {
		quotient := float64(input[0]) / float64(input[1])
		weight := input[0] + input[1]
		for i := 0; i < weight; i++ {
			subScores = append(subScores, quotient)
		}
	}

	switch strategy {
	case Mean:
		return stats.Mean(subScores)
	case Median:
		return stats.Median(subScores)
	case Mode:
		modes, err := stats.Mode(subScores)
		if err != nil {
			return 0, err
		}
		// Arbitrarily return the first mode.
		return modes[0], nil
	default:
		return 0, errors.New("unknown strategy")
	}
}
