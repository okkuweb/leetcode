package main

import (
	"strings"
)

// @leet start
func letterCombinations(digits string) []string {
	var phoneArr = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	var finalArr []string
	digitArr := strings.Split(digits, "")
	finalLen := len(digitArr)

	var loop func(string, int, int) bool
	loop = func(final string, depthN int, charN int) bool {
		for _, v := range phoneArr[digitArr[depthN]] {
			finalfinal := final + v
			if len(finalfinal) == finalLen {
				finalArr = append(finalArr, finalfinal)
			}
			if len(phoneArr[digitArr[depthN]]) > charN {
				if len(digitArr) > depthN+1 {
					loop(finalfinal, depthN+1, 0)
				}
				charN++
			}
			finalfinal = ""
			if charN >= len(phoneArr[digitArr[depthN]]) {
				return false
			}
		}
		return true
	}

	if len(digitArr) > 0 {
		for loop("", 0, 0) {
		}
	}

	return finalArr
}

// @leet end
