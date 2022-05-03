package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
)

type match struct {
	filename string
	score float64
}

type overallScore struct {
	correct float64
	tied float64
}

func main() {
	textChains := map[string]map[string]map[string]int{}
	files, err := listFiles(textDir)
	check(err)

	for _, file := range files {
		text, err := readFile(file)
		check(err)

		filename := strings.Replace(file, textDir + "/", "", 1)
		filename = strings.Replace(filename, ".txt", "", 1)
		textChains[filename] = createChain(text)
	}

	results := map[uint]overallScore{}
	for _, topWordPairs := range []uint{1, 10, 20, 30, 40, 50, 100} {
		passageMatches := map[passage][]match{}
		for _, passage := range passages {
			tempMatches := []match{}
			passageChain := createChain(passage.text)
			for filename, textChain := range textChains {
				scoreOptions := scoreOptions{
					strategy: Mean,
					topWordPairs: topWordPairs,
				}
				score, err := getChainSimilarityScore(passageChain, textChain, scoreOptions)
				check(err)

				tempMatches = append(tempMatches, match{
					filename: filename,
					score: score,
				})
			}
			sort.Slice(tempMatches, func(i, j int) bool {
				return tempMatches[i].score > tempMatches[j].score // using '>' to sort descending
			})
			passageMatches[passage] = tempMatches
		}

		var numCorrect uint
		var numTied uint
		for passage, matches := range passageMatches {
			firstGuessCorrect, err := isCorrect(matches[0].filename, passage)
			check(err)

			if firstGuessCorrect {
				numCorrect += 1

				secondGuessCorrect, err := isCorrect(matches[1].filename, passage)
				check(err)
				if secondGuessCorrect {
					numTied += 1
				}
			}
		}
		results[topWordPairs] = overallScore{
			correct: float64(numCorrect) / float64(len(passageMatches)),
			tied: float64(numTied) / float64(numCorrect),
		}
	}

	fmt.Printf("%+v", results)
}

// isCorrect returns true if the given filename corresponds to the given passage,
// false otherwise.
func isCorrect(filename string, passage passage) (bool, error) {
	parts := strings.Split(filename, " - ")
	if len(parts) != 2 {
		return false, errors.New("invalid filename: " + filename)
	}
	author := parts[0]
	title := parts[1]
	return passage.author == author && passage.title == title, nil
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
