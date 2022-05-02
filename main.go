package main

import (
	"encoding/json"
	"fmt"
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
		filename = strings.Replace(file, ".txt", "", 1)
		textChains[filename] = createChain(text)
	}

	results := map[passage][]result{}
	for _, passage := range passages {
		results[passage] = []result{}
		for filename, textChain := range textChains {
			for _, strategy := range strategies {
				passageChain := createChain(passage.text)
				score, err := getChainSimilarityScore(passageChain, textChain, strategy)
				check(err)

				results[passage] = append(results[passage], result{
					filename: filename,
					strategy: strategy,
					score: score,
				})
			}
		}
		break
	}

	prettyPrint(results)
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