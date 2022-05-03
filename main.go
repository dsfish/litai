package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type result struct {
	filename string
	strategy strategy
	score float64
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

	passageToResults := map[passage][]result{}
	for _, passage := range passages {
		passageChain := createChain(passage.text)
		var results []result
		for filename, textChain := range textChains {
			for _, strategy := range []strategy{Mean} { // TODO: Decide whether to include other strategies.
				score, err := getChainSimilarityScore(passageChain, textChain, scoreOptions{
					strategy: strategy,
					topWordPairs: 100,
				})
				check(err)

				if score == 0 {
					continue
				}

				results = append(results, result{
					filename: filename,
					strategy: strategy,
					score: score,
				})
			}
		}

		sort.Slice(results, func(i, j int) bool {
			return results[i].score > results[j].score // using '>' to sort descending
		})
		passageToResults[passage] = results
	}

	for passage, results := range passageToResults {
		prettyPrint(map[string]interface{}{
			"passage": map[string]interface{}{
				"author": passage.author,
				"title": passage.title,
				"year": passage.year,
			},
			"results": fmt.Sprintf("%+v", results),
		})
	}
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