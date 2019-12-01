package main

import (
	"fmt"
	"sort"
)

func getMapOfCounts(str string) map[rune]uint16 {
	charCount := make(map[rune]uint16)
	for _, ch := range str {
		val, ok := charCount[ch]
		if !ok {
			charCount[ch] = 1
		} else {
			charCount[ch] = val + 1
		}
	}
	return charCount
}

func main() {
	str1 := "ABRAKADABRA"
	map1 := getMapOfCounts(str1)
	strList := []string{"KADABRAABRA", "KADABRA", "ABRAKBDABRA"}

	// method 1
	fmt.Println("Method 1")
	for _, str := range strList {
		if len(str) != len(str1) {
			fmt.Println(str + " is not a permutation of " + str1)
		} else {
			mapReturned := getMapOfCounts(str)
			isPermutation := true
			for ch, count := range map1 {
				if mapReturned[ch] != count {
					fmt.Println(str + " is not a permutation of " + str1)
					isPermutation = false
					break
				}
			}
			if isPermutation {
				fmt.Println(str + " is a permutation of " + str1)
			}
		}
	}

	// method 2
	fmt.Println("Method 2")
	ar1 := []byte(str1)
	sort.Slice(ar1, func(i, j int) bool {
		return int(ar1[i]) < int(ar1[j])
	})
	fmt.Println(string(ar1))
	// Check if both of these are the same

}
