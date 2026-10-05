func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[byte]int)
	nonZero := 0

	for i := 0; i < len(s); i++ {
		byteS := s[i]
		counts[byteS]++

		if counts[byteS] > 0 {
			nonZero++
		}
	}

	for i := 0; i < len(t); i++ {
		byteT := t[i]
		counts[byteT]--

		if (counts[byteT] < 0) {
			nonZero++;
		} else {
			nonZero--
		}
	}

	return nonZero == 0
}
