package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func validateString(s string) bool {
	return s != ""
}

//Not so much efficient approach for using goroutines -

//func preProcessStrings2(filePath string) *Trie {
//
//	trie := NewTrie()
//
//	file, err := os.Open(filePath)
//	if err != nil {
//		fmt.Printf("Error opening the file: %v\n", err)
//		return nil
//	}
//	defer file.Close()
//
//	var wg sync.WaitGroup
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		line := scanner.Text()
//		wg.Add(1)
//		go func(s string) {
//			defer wg.Done()
//			s = strings.TrimSpace(s)
//			if validateString(s) {
//				trie.Insert(s)
//			}
//		}(line)
//	}
//	wg.Wait()
//	return trie
//}

// Better approach by using optimal no. of goroutines
func worker(jobs <-chan string, wg *sync.WaitGroup, trie *Trie) {
	defer wg.Done()
	for s := range jobs {
		s = strings.TrimSpace(s)
		if validateString(s) {
			trie.Insert(s)
		}
	}
}

func preProcessStrings(filePath string, numWorkers int) *Trie {
	trie := NewTrie()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return nil
	}
	defer file.Close()

	var wg sync.WaitGroup
	jobs := make(chan string, numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, &wg, trie)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		jobs <- line
	}
	close(jobs)

	wg.Wait()

	return trie
}
