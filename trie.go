package main

import (
	"sync"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
	mu   sync.RWMutex
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) findLongestPrefix(input string) string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	node := t.root
	longestPrefix := ""
	for _, char := range input {
		if child, ok := node.children[char]; ok {
			node = child
			longestPrefix += string(char)
		} else {
			break
		}
	}
	return longestPrefix
}
