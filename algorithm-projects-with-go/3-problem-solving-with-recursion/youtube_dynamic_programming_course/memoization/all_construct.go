package main

import (
	"fmt"
	"strings"
)

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=10050s

// Write a function 'allConstruct(target, workBank)' that accepts a target string and an array of strings
// The function should return a 2D array containing all of the ways that the "target" can be constructed by concatenating elements of the "wordBank" array
// Each element of the 2D array should represent one combination that constructs the "target".
// You may reuse elements of "wordBank" as many times as needed

func allConstruct(target string, workBank []string) [][]string {
	if len(target) == 0 {
		return [][]string{[]string{}}
	}

	results := [][]string{}
	for _, word := range workBank {
		if strings.HasPrefix(target, word) {
			moreWays := allConstruct(target[len(word):], workBank)
			for i, _ := range moreWays {
				moreWays[i] = append([]string{word}, moreWays[i]...)
			}
			results = append(results, moreWays...)
		}
	}

	return results
}

func allConstructMemoized(target string, workBank []string, memo map[string][][]string) [][]string {
	if value, ok := memo[target]; ok {
		return value
	}
	if len(target) == 0 {
		return [][]string{[]string{}}
	}

	results := [][]string{}
	for _, word := range workBank {
		if strings.HasPrefix(target, word) {
			moreWays := allConstructMemoized(target[len(word):], workBank, memo)
			for i, _ := range moreWays {
				moreWays[i] = append([]string{word}, moreWays[i]...)
			}
			results = append(results, moreWays...)
		}
	}

	memo[target] = results
	return results
}

func main() {
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))
	fmt.Println(allConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	fmt.Println(allConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(allConstructMemoized("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
		"eeeeee",
		// "eeeeeef",
	}, make(map[string][][]string)))
}
