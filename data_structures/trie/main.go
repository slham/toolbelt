package main

import "log"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func Init() *Trie {
	out := &Trie{root: &Node{}}
	return out
}

// Insert
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	curr := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if curr.children[charIndex] == nil {
			curr.children[charIndex] = &Node{}
		}
		curr = curr.children[charIndex]
	}
	curr.isEnd = true
}

// Search
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	curr := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if curr.children[charIndex] == nil {
			return false
		}
		curr = curr.children[charIndex]
	}
	return curr.isEnd
}

func main() {
	trie := Init()
	log.Println(trie)
	log.Println(trie.root)
	words := []string{"hello", "world", "help", "helping", "word", "worldly"}
	for _, w := range words {
		trie.Insert(w)
	}
	log.Println("searching hello", trie.Search("hello"))
	log.Println("searching world", trie.Search("world"))
	log.Println("searching hola", trie.Search("hola"))
	log.Println("searching worldly", trie.Search("worldly"))
	log.Println("searching weird", trie.Search("weird"))
}
