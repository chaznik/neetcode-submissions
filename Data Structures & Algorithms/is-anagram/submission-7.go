func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, char := range s {
		sMap[char]++
	}

	for _, char := range t {
		sMap[char]--

		if sMap[char] < 0 {
			return false
		}
	}
	return true
}
