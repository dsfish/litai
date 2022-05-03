package main

import (
	"errors"
	"github.com/montanaflynn/stats"
	"regexp"
	"sort"
	"strings"
)

const defaultTopWordPairs = 10000

type scoreOptions struct{
	// The mathematical operation to use (mean, median, mode) when calculating a
	// weighted score.
	strategy strategy

	// The number of word pairs to include when calculating a weighted score. If 0
	// is given, up to defaultTopWordPairs word pairs will be included. If an
	// integer n is given, only the top n word pairs will be included, which allows
	// the caller to reduce noise. This can be helpful when one chain is
	// significantly larger than the other.
	topWordPairs uint64
}

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
// identical.
func getChainSimilarityScore(
	chainA, chainB map[string]map[string]int,
	options scoreOptions,
) (float64, error) {
	if options.topWordPairs == 0 {
		options.topWordPairs = defaultTopWordPairs
	}

	// For every word pair in chainA, record the number of times that the word pair
	// occurs in each chain, storing the smaller value first.
	var scoringInput [][]int
	for fromWord, toWords := range chainA {
		for toWord, aN := range toWords {
			var bN int
			if chainB[fromWord] != nil {
				bN = chainB[fromWord][toWord]
			}
			if aN < bN {
				scoringInput = append(scoringInput, []int{aN, bN})
			} else {
				scoringInput = append(scoringInput, []int{bN, aN})
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
			scoringInput = append(scoringInput, []int{0, bN})
		}
	}

	// Sort the map in descending order.
	sort.Slice(scoringInput, func(i, j int) bool {
		iQuotient := float64(scoringInput[i][0]) / float64(scoringInput[i][1])
		jQuotient := float64(scoringInput[j][0]) / float64(scoringInput[j][1])
		return iQuotient > jQuotient // using '>' to sort descending
	})

	var subScores []float64
	wordPairsRemaining := options.topWordPairs
	for _, input := range scoringInput {
		quotient := float64(input[0]) / float64(input[1])
		weight := input[0] + input[1]
		for i := 0; i < weight; i++ {
			subScores = append(subScores, quotient)
		}

		wordPairsRemaining -= 1
		if wordPairsRemaining <= 0 {
			break
		}
	}

	switch options.strategy {
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
