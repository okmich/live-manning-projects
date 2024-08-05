package main

import (
	"fmt"
	"strings"
)

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=7965s

// Write a function 'countConstruct(target, workBank)' that accepts a target string and an array of strings
// The function should return the number of ways that the "target" can be constrcuted by concatenating elements of the "wordBank" array
// You may reuse elements of "wordBank" as many times as needed

func countConstruct(target string, workBank []string) int {
	if len(target) == 0 {
		return 1
	}

	var result int = 0
	for _, word := range workBank {
		if strings.HasPrefix(target, word) {
			result += countConstruct(target[len(word):], workBank)
		}
	}

	return result
}

func countConstructMemoized(target string, workBank []string, memo map[string]int) int {
	if value, ok := memo[target]; ok {
		return value
	}
	if len(target) == 0 {
		return 1
	}

	var result int = 0
	for _, word := range workBank {
		if strings.HasPrefix(target, word) {
			result += countConstructMemoized(target[len(word):], workBank, memo)
		}
	}

	memo[target] = result
	return memo[target]
}

func main() {
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                  //1
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))                  //2
	fmt.Println(countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))   //0
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) //4
	fmt.Println(countConstructMemoized("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
		"eeeeee",
		// "eeeeeef",
	}, make(map[string]int))) //0 or 867844316
}
