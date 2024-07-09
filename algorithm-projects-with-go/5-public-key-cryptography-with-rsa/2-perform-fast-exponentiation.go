package main

import (
	"fmt"
	"math"
	"strconv"
)

func fastExp(num, pow int) int {
	result := 1
	current_pow := num
	exp := pow
	for exp >= 1 {
		if exp%2 == 1 {
			result *= current_pow
		}
		exp /= 2
		current_pow *= current_pow
	}
	return result
}

func fastExpMod(num, pow, mod int) int {
	var result int = 1
	current_pow := num
	exp := pow
	for exp >= 1 {
		if exp%2 == 1 {
			result = (result * current_pow) % mod
		}
		exp /= 2
		current_pow = (current_pow * current_pow) % mod
	}
	return result
}

func main() {
	for {
		num := ""
		pow := ""
		mod := ""

		fmt.Printf("Enter a number: ")
		fmt.Scanln(&num)
		if num == "" {
			break
		}
		intNum, err := strconv.ParseInt(num, 10, 64)
		if err != nil || intNum < 1 {
			fmt.Println(err)
			break
		}

		fmt.Printf("Enter the pow value: ")
		fmt.Scanln(&pow)
		if pow == "" {
			break
		}
		intPow, err := strconv.ParseInt(pow, 10, 64)
		if err != nil || intPow < 1 {
			fmt.Println(err)
			break
		}

		fmt.Printf("Enter modulo value: ")
		fmt.Scanln(&mod)
		if mod == "" {
			break
		}
		intMod, err := strconv.ParseInt(mod, 10, 64)
		if err != nil || intMod < 1 {
			fmt.Println(err)
			break
		}

		fastExpResult := fastExp(int(intNum), int(intPow))
		fastExpModResult := fastExpMod(int(intNum), int(intPow), int(intMod))
		rawExpResult := math.Pow(float64(intNum), float64(intPow))

		fmt.Println("*******************************************")
		fmt.Println("Fast Exp. results: ", fastExpResult)
		fmt.Println("Fast Mod. results: ", int(fastExpModResult))
		fmt.Println("Raw  Exp. results: ", int(rawExpResult))
		fmt.Println("Raw  Mod. results: ", int64(rawExpResult)%intMod)
		fmt.Println("*******************************************")
		fmt.Println("\n")
	}
}
