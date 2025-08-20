package main

import (
	"fmt"
	"strings"
)

// @leet start
func letterCombinations(digits string) []string {
	phoneArr := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	digitArr := strings.Split(digits, "")
	var finalString []string
	var loopsArr []int
	if len(digitArr) > 0 {
		for _, v := range digitArr {
			append(loops, len(phoneArr[v]))
		}
		loops := 1
		for _, v := range loops {
			loopsVal *= v
		}

		incrementer := map[string]string{}
		for _, d2 := range digitArr {
			incrementer{d2} = 1
		}
	}

	return finalString
}

// @leet end
