package main

import "fmt"

type Trie struct {
	letterMap map[rune]*Trie
	isWord    bool
}

func Constructor() Trie {
	return Trie{
		letterMap: make(map[rune]*Trie, 0),
		isWord:    false,
	}
}

func (this *Trie) Insert(word string) {
	t := *this
	for _, letter := range word {
		curr, exist := t.letterMap[letter]
		if !exist {
			fmt.Println(t)
			newCons := Constructor()
			t.letterMap[letter] = &newCons
			t = newCons
		} else {
			t = *curr
		}
	}
	t.isWord = true
}

func (this *Trie) Search(word string) bool {
	t := *this
	var is bool
	for _, letter := range word {
		fmt.Println("searching %s", letter)
		is = t.isWord
		tier, exist := t.letterMap[letter]
		if !exist {
			return false
		}
		t = *tier
	}
	return is

}

// func (this *Trie) StartsWith(prefix string) bool {

// }

func main() {
	obj := Constructor()
	obj.Insert("cat")
	isExist := obj.Search("cat")
	fmt.Println(isExist)
}
