package main

import (
	"fmt"
)

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=16686s

// Write a function 'countConstruct(target, workBank)' that accepts a target string and an array of strings
// The function should return the number of ways that the "target" can be constrcuted by concatenating elements of the "wordBank" array
// You may reuse elements of "wordBank" as many times as needed

func countConstruct(target string, wordBank []string) int {
	table := make([]int, len(target)+1)
	//seed first result
	table[0] = 1

	targetLength := len(target)
	wordLength := 0
	for i, _ := range table {
		if table[i] > 0 {
			for _, word := range wordBank {
				wordLength = len(word)
				if i+wordLength <= targetLength && word == target[i:i+wordLength] {
					table[i+wordLength] += table[i]
				}
			}
		}
	}

	return table[targetLength]
}

func main() {
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                  //1
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))                  //2
	fmt.Println(countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))   //0
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) //4
	fmt.Println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
		"eeeeee",
		// "eeeeeef",
	})) //0 or 867844316
}
