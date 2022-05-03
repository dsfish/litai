package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type match struct {
	filename string
	score float64
}

type overallScore struct {
	// The percentage of guesses that were correct.
	correct float64

	// A map of each passage to its incorrect match (only provided if matching
	// failed).
	incorrect map[passage]string
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

	results := map[uint64]overallScore{}
	var topWordPairs uint64
	for topWordPairs = 0; topWordPairs < 100; topWordPairs += 5 {
		passageMatches := map[passage][]match{}
		for _, passage := range passages {
			tempMatches := make([]match, len(textChains))
			passageChain := createChain(passage.text)
			i := 0
			for filename, textChain := range textChains {
				scoreOptions := scoreOptions{
					strategy: Mean,
					topWordPairs: topWordPairs,
				}
				score, err := getChainSimilarityScore(passageChain, textChain, scoreOptions)
				check(err)

				tempMatches[i] = match{
					filename: filename,
					score: score,
				}
				i++
			}
			sort.Slice(tempMatches, func(i, j int) bool {
				return tempMatches[i].score > tempMatches[j].score // using '>' to sort descending
			})
			passageMatches[passage] = tempMatches
		}

		var numCorrect uint
		incorrect := map[passage]string{}
		for passage, matches := range passageMatches {
			isCorrect, err := isCorrect(matches[0].filename, passage)
			check(err)

			if isCorrect {
				numCorrect += 1
			} else {
				incorrect[passage] = matches[0].filename
			}
		}
		results[topWordPairs] = overallScore{
			correct: float64(numCorrect) / float64(len(passageMatches)),
			incorrect: incorrect,
		}
	}

	var outcomes []map[string]interface{}
	for topWordPairs, score := range results {
		var incorrectGuesses []map[string]interface{}
		for passage, filename := range score.incorrect {
			guessedAuthor, guessedTitle, err := getAuthorAndTitle(filename)
			check(err)
			incorrectGuesses = append(incorrectGuesses, map[string]interface{}{
				"guess": map[string]interface{}{
					"author": guessedAuthor,
					"title": guessedTitle,
				},
				"actual": map[string]interface{}{
					"author": passage.author,
					"title": passage.title,
				},
			})
		}
		outcomes = append(outcomes, map[string]interface{}{
			"num_word_pairs": strconv.FormatUint(topWordPairs, 10),
			"score": fmt.Sprintf("%.2f", score.correct * 100),
			"incorrect_guesses": incorrectGuesses,
		})
	}

	sort.Slice(outcomes, func(i, j int) bool {
		iScore, err := strconv.ParseFloat(outcomes[i]["score"].(string), 64)
		check(err)
		jScore, err := strconv.ParseFloat(outcomes[j]["score"].(string), 64)
		check(err)
		return iScore > jScore // using '>' to sort descending
	})

	var scores []string
	for _, outcome := range outcomes {
		s := fmt.Sprintf("%s, %s", outcome["num_word_pairs"], outcome["score"])
		scores = append(scores, s)
	}

	var b []byte
	b = append(b, []byte("Scores:")...)
	bScores, err := json.MarshalIndent(scores, "", "  ")
	check(err)
	b = append(b, bScores...)

	b = append(b, []byte("\n\nOutcomes:")...)
	bOutcomes, err := json.MarshalIndent(outcomes, "", "  ")
	check(err)
	b = append(b, bOutcomes...)
	err = ioutil.WriteFile("output.txt", b, fs.ModeAppend)
	check(err)
}

// isCorrect returns true if the given filename corresponds to the given passage,
// false otherwise.
func isCorrect(filename string, passage passage) (bool, error) {
	author, title, err := getAuthorAndTitle(filename)
	if err != nil {
		return false, err
	}
	return passage.author == author && passage.title == title, nil
}

func getAuthorAndTitle(filename string) (string, string, error) {
	parts := strings.Split(filename, " - ")
	if len(parts) != 2 {
		return "", "", errors.New("invalid filename: " + filename)
	}
	author := parts[0]
	title := parts[1]
	return author, title, nil
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
