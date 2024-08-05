package main

import (
	"fmt"
	"strings"
)

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=7965s

// Write a function 'canConstruct(target, workBank)' that accepts a target string and an array of strings
// The function should return a boolean indicating whether or not the "target" can be constructed by concatenating elements of the "wordBank" array
// You may reuse elements of "wordBank" as many times as needed

func canConstruct(target string, workBank []string) bool {
	if len(target) == 0 {
		return true
	}

	for _, word := range workBank {
		if strings.HasPrefix(target, word) {
			if canConstruct(target[len(word):], workBank) {
				return true
			}
		}
	}

	return false
}

func canConstructMemoized(target string, workBank []string, memo map[string]bool) bool {
	if value, ok := memo[target]; ok {
		return value
	}
	if len(target) == 0 {
		return true
	}

	for _, word := range workBank {
		if strings.HasPrefix(target, word) {
			if canConstructMemoized(target[len(word):], workBank, memo) {
				memo[target] = true
				return true
			}
		}
	}

	memo[target] = false
	return false
}

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                  //true
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))   //false
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) //true
	fmt.Println(canConstructMemoized("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
		"eeeeee",
		// "eeeeeef",
	}, make(map[string]bool))) //false
}
