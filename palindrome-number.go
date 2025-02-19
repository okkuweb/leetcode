func isPalindrome(x int) bool {
	xstr := strconv.Itoa(x)
	xrunes := []rune(xstr)

	for i, j := 0, len(xrunes)-1; i < j; i, j = i+1, j-1 {
		xrunes[i], xrunes[j] = xrunes[j], xrunes[i]
	}

	if string(xrunes) == xstr {
		return true
	}
	return false
}

