package main

import (
	"fmt"	"strings"
)


func main() {
	fmt.Println("bismillah")
	filename := "birthday_001.txt"
	// => Birthday
	newName, err := match(filename)
	if err != nil {
		fmt.Println("no match")
	}
	fmt.Println(newName)
}
func match(filename string) (string, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	//Join other parts which get splitted but not extention
	strings.Join(pieces[:])
	return "", nil
}
