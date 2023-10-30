package main

import (
	"fmt"
)

func main() {
	//t1 := time.Now()
	inputString := "KAWE"
	filePath := "sample_prefixes.txt"
	trie := preProcessStrings(filePath, 20)
	//trie := preProcessStrings2(filePath)
	longestPrefix := trie.findLongestPrefix(inputString)
	fmt.Println(longestPrefix)
	//t2 := time.Now()
	//fmt.Println(float64(t2.Nanosecond() - t1.Nanosecond()))
}
