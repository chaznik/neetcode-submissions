func isValid(s string) bool {
    stack := []byte{}
    pairs := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for i := 0; i < len(s); i++ {
        bracket := s[i]

        if opening, ok := pairs[bracket]; ok {
            if len(stack) == 0 || stack[len(stack)-1] != opening {
                return false
            }

            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, bracket)
        }
    }

    return len(stack) == 0
}