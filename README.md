# litai

## Overview
This is my final project for Professor Dennis Tenen's [Literature in the Age of Artificial Intelligence](https://github.com/denten-courses/LITAI/tree/master/2022) class, which I took in the spring of 2022.

For the midterm, students received a list of passages from texts we'd read throughout the semester, and we had to identify their sources. For this project, I've created a program that uses Markov chains to take the midterm.

Of course, a better approach would be to merely search for the passage in each source text. The goal of this project is not to produce the best possible test-taker but to learn about Markov chains and their applications and--if we're lucky--uncover some hidden Markov chain behavior.

The program works by constructing Markov chains out of the given [passages](./passages.go) and source texts](./texts) and matching them based on their similarity. The most interesting part of this program is the scoring function, `getChainSimilarityScore`, that determines how similar two Markov chains are.

## How scoring works

`getChainSimilarityScore` takes two Markov chains in the form of maps and returns a "similarity score" in the range [0, 1]. Here are two example sentences and their corresponding chains:

"That _fish_ ate that fish."
```json
{
  "ate": {
    "that": 1
  },
  "fish": {
    "ate": 1
  },
  "that": {
    "fish": 2
  }
}
```

"That _cat_ ate that fish."
```json
{
  "ate": {
    "that": 1
  },
  "cat": {
    "ate": 1
  },
  "that": {
    "cat": 1,
    "fish": 1
  }
}
```

Next, it creates a sub-score for every pair of consecutive words (e.g., "that fish") in each chain. For example, in the first chain, the word pair "that fish" appears twice, but in the second chain, it only appears once. The sub-score is calculated by diving the smaller number by the larger number. Thus, the score for the word pair "that fish" is 1/2, or 0.5.

I believe that a word pair that appears _more_ frequently should have a greater impact on the overall score than a word pair that appears _less_ frequently. Thus, the function weights these sub-scores according to the number of times they appear. For example, "that fish" appears three times across both chains, so its weight is three. That is, the sub-score of 0.5 will be factored in the overall score three times.

The function supports different mathematical "strategies" for aggregating these sub-scores into the overall score: mean, median, and mode. In my testing, the mean strategy has proven most effective.

The function also allows the user to specify the number of top word pairs to include in the calculation. For example, the top word pair in the first chain above is "that fish" because it appears twice. In my testing, I have found that setting this parameter to a relatively low number improves scoring accuracy.

## Running the program
The program can be run by executing `go run .` in the root directory. When run, the program performs a Monte Carlo simulation in which it takes the midterm repeatedly using different values for `options.topWordPairs`. It records the results in [output.txt](./output.txt).

## Analysis
