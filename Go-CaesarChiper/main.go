package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var length int
	var original string
	var delta int32
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &original)
	fmt.Scanf("%d\n", &delta)
	fmt.Printf("length: %d, text: %s, delta: %d", length, original, delta)
	result := caesarCipher(original, delta)
	fmt.Printf("result: %v\n", result)
}

func caesarCipher(s string, k int32) string {
	var isUpperCase bool
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	var result []string
	var targetChar rune
	var newIdx int
	for _, str := range s {
		isUpperCase = unicode.IsUpper(str)
		strLower := unicode.ToLower(str)
		found, idx := Find(alphabet, string(strLower))
		if found {
			newIdx = (idx + int(k)) % 26
			if isUpperCase {
				targetChar = unicode.ToUpper(rune(alphabet[newIdx]))
			} else {
				targetChar = rune(alphabet[newIdx])
			}
			result = append(result, string(targetChar))
		} else {
			result = append(result, string(str))
		}
	}
	return strings.Join(result, "")
}

func Find(alphabet string, str string) (bool, int) {
	for i, char := range alphabet {
		if string(char) == str {
			return true, i
		}
	}
	return false, 0
}
