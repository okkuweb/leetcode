func scoreOfString(s string) int {
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		value := int(s[i]) - int(s[i+1])
		if value < 0 {
			sum += -value
		} else {
			sum += value
		}
	}
	return sum
}
