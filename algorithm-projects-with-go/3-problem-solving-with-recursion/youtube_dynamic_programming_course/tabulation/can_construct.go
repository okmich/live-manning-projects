package main

import "fmt"

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=15650s

// Write a function 'canConstruct(target, workBank)' that accepts a target string and an array of strings
// The function should return a boolean indicating whether or not the "target" can be constructed by concatenating elements of the "wordBank" array
// You may reuse elements of "wordBank" as many times as needed

func canConstruct(target string, wordBank []string) bool {
	table := make([]bool, len(target)+1)
	table[0] = true //seed the first position

	targetLength := len(target)
	for i, _ := range table {
		if table[i] {
			for _, word := range wordBank {
				if targetLength >= i+len(word) && word == target[i:i+len(word)] {
					table[i+len(word)] = true
				}
			}
		}
	}

	return table[len(target)]
}

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                  //true
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))   //false
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) //true
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
		"eeeeee",
		// "eeeeeef",
	})) //false
}
