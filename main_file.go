package main

import "fmt"

func main() {

	fmt.Println("")
	fmt.Println("Test 1, most occurences:")
	fmt.Println("in a string 'abc bbac', the most occured letter is:", MostOccurences("abc bbac"))

	fmt.Println("--")
	fmt.Println("Test 2, is anagram:")
	fmt.Println("strings 'abc cba', is anagram:", IsAnagram("abc", "cba"))
	fmt.Println("strings 'abc xba', is anagram:", IsAnagram("abc", "xbc"))

	fmt.Println("--")
	longestFilename, _ := FindFirstLongestFileNameInDir("./")
	fmt.Println("Test 3, first longest filename in current folder:", longestFilename)

	fmt.Println("--")
	fmt.Println("Test 4, prosječan tečaj HRK u odnosu na USD u zadnjih 5 dana:", ProsjecanTecajHrkUsd())

	fmt.Println("--")
	fmt.Print("Test 5, permutations for string 'abc':")
	pt := Permutations{}
	for _, s := range pt.GetPermutations("abc") {
		fmt.Print(" ", s)
	}
	fmt.Println()
	fmt.Println("--end of test--")

}
