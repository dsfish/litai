package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const input string = "But [oratory] should be dignified and moving. " +
	"The trochaic meter is\nrather too much of a comic dance, " +
	"as is clear from trochaic tetrameters; for they are a tripping rhythm. " +
	"What remains is the paean; it\ncame into use beginning with Thrasymachus, " +
	"though at the time people did not recognize what it was. The paean is a third kind of rhythm," +
	"\nrelated to those under discussion; for it has the ratio of three to two\n[three short syllables and one long, " +
	"the latter equal in time to two\nbeats], whereas the others are one to one [the heroic, " +
	"with one long\nand two shorts] or two to one [iambic and trochaic, a long and a short\nor a short and a long, " +
	"respectively]."

func main() {
	texts := map[string]string{}

	files, err := listFiles(textDir)
	check(err)

	for _, file := range files {
		text, err := readFile(file)
		check(err)

		filename := strings.Replace(file, "texts/", "", 1)
		texts[filename] = text
	}

	results := map[string]map[string]string{}

	inputChain := createChain(input)
	var bestScore float64
	var bestStrategy Strategy
	var bestMatch string
	for name, text := range texts {
		results[name] = map[string]string{}
		textChain := createChain(text)

		var bestScoreForText float64
		var bestStrategyForText Strategy
		for _, strategy := range strategies {
			score, err := getChainSimilarityScore(inputChain, textChain, strategy)
			check(err)

			results[name][strategy.String()] = fmt.Sprintf("%.6f", score)

			if score >= bestScoreForText {
				bestScoreForText = score
				bestStrategyForText = strategy
			}
		}

		if bestScoreForText > bestScore {
			bestScore = bestScoreForText
			bestStrategy = bestStrategyForText
			bestMatch = name
		}
	}

	if bestMatch == "" {
		bestMatch = "none"
	}

	prettyPrint(map[string]interface{}{
		"best_match": bestMatch,
		"best_score": bestScore,
		"best_strategy": bestStrategy.String(),
		"detailed_results": results,
	})
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func prettyPrint(i interface{}) {
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}