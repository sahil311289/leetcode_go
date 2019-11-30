package main

import "fmt"

func main() {
	list := []string{"hello", "helo 123", "hello 12 3"}

	// method 1
	fmt.Println("Method 1: ")
	for _, str := range list {
		countMap := make(map[rune]bool)
		found := false
		for _, ch := range str {
			_, ok := countMap[ch]
			if !ok {
				countMap[ch] = true
			} else {
				fmt.Println("Duplicate char(s) found in this string: " + str)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Unque chars in this string: " + str)
		}
	}

	// method 2
	fmt.Println("Method 2: ")
	for _, str := range list {
		found := false
		var chList [128]bool
		for _, ch := range str {
			if chList[int(ch)] {
				fmt.Println("duplicate char found in string: " + str)
				found = true
				break
			} else {
				chList[int(ch)] = true
			}
		}
		if !found {
			fmt.Println("Unique chars found in str: " + str)
		}
	}
}
