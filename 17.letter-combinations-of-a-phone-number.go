package main

import (
	"strings"
)

// @leet start
func letterCombinations(digits string) []string {
	buttons := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var finalSlice []string
	bases := strings.Split(digits, "")
	if len(bases) < 1 {
		return finalSlice
	}

	var loop func(string, int)
	loop = func(combination string, depth int) {
		chars := strings.Split(buttons[bases[depth]], "")
		for _, v := range chars {
			finalString := combination + string(v)
			if len(finalString) == len(digits) {
				finalSlice = append(finalSlice, finalString)
			} else {
				loop(finalString, depth+1)
			}
		}
	}

	loop("", 0)

	return finalSlice
}

// @leet end
