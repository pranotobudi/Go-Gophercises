package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Scanf("%s \n", &input)
	counter := camelcase(input)
	fmt.Printf("%v", counter)
}
func camelcase(s string) int32 {
	var result int32 = 0
	if s == "" {
		return result
	}
	for _, str := range s {
		if unicode.IsUpper(rune(str)) {
			result++
		}
	}
	return result + 1
}
