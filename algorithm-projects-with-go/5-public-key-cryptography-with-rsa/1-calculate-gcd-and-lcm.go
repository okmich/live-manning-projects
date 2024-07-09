package main

import (
	"fmt"
	"strconv"
)

func gcd(a, b int64) int64 {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else {
		max := max(a, b)
		min := min(a, b)

		// q := max / min
		r := max % min
		return gcd(min, r)
	}
}

func lcm(a, b int64) int64 {
	gcd := gcd(a, b)
	return (a * b) / gcd
}

func main() {
	for {
		value1 := ""
		value2 := ""

		fmt.Printf("Enter first value: ")
		fmt.Scanln(&value1)
		if value1 == "" {
			break
		}
		intValue1, err := strconv.ParseInt(value1, 10, 64)
		if err != nil || intValue1 < 1 {
			fmt.Println(err)
			break
		}

		fmt.Printf("Enter second value: ")
		fmt.Scanln(&value2)
		if value2 == "" {
			break
		}
		intValue2, err := strconv.ParseInt(value2, 10, 64)
		if err != nil || intValue2 < 1 {
			fmt.Println(err)
			break
		}
		gcd := gcd(intValue1, intValue2)
		lcm := lcm(intValue1, intValue2)

		fmt.Printf("\n\n%-8s \t %-8s \t %-12s \t %-12s \n", "A", "B", "gcd", "lcm")
		fmt.Printf("%-8d \t %-8d \t %-12d \t %-12d \n", intValue1, intValue2, gcd, lcm)
		fmt.Println()
	}
}
