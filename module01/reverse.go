package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// func Reverse(word string) string {
// 	var reversed string
// 	for i := len(word) - 1; i >= 0; i-- {
// 		reversed += string(word[i])
// 	}

// 	return reversed
// }

// Reverse managing better the string
// func Reverse(word string) string {
// 	var sb strings.Builder
// 	for i := len(word) - 1; i >= 0; i-- {
// 		sb.WriteByte(word[i])
// 	}

// 	return sb.String()
// }

// Reverse starting with the first letter
// func Reverse(word string) string {
// 	var reversed string
// 	for i := 0; i < len(word); i++ {
// 		reversed = string(word[i]) + reversed
// 	}

// 	return reversed
// }

// Reverse including japanese words
func Reverse(word string) string {
	var res string
	for _, r := range word {
		res = string(r) + res
	}

	return res
}
