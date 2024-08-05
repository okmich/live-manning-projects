package main

import (
	"fmt"
)

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=17423s

// Write a function 'allConstruct(target, workBank)' that accepts a target string and an array of strings
// The function should return a 2D array containing all of the ways that the "target" can be constructed by concatenating elements of the "wordBank" array
// Each element of the 2D array should represent one combination that constructs the "target".
// You may reuse elements of "wordBank" as many times as needed

func allConstruct(target string, wordBank []string) [][]string {
	targetLength := len(target)
	table := make([][][]string, targetLength+1)
	//all is nul
	//now seed first
	table[0] = [][]string{}

	for i, _ := range table {
		if table[i] != nil {
			for _, word := range wordBank {
				wordLength := len(word)
				if i+wordLength <= targetLength && word == target[i:i+wordLength] {
					if table[i+wordLength] == nil && len(table[i]) == 0 {
						table[i+wordLength] = [][]string{[]string{word}}
					} else {
						for k, _ := range table[i] {
							table[i+wordLength] = append(table[i+wordLength], append(table[i][k], word))
						}
					}
				}
			}
		}
	}

	return table[targetLength]
}

func main() {
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))
	fmt.Println(allConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	fmt.Println(allConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(allConstruct("eeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
		"eeeeee",
		// "eeeeeef",
	}))
}
