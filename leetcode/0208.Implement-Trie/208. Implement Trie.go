package leetcode

type Trie struct {
	children  [26](*Trie)
	endOfWord bool
}

func Constructor() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	pointer := trie
	for _, char := range word {
		if pointer.children[char-'a'] == nil {
			pointer.children[char-'a'] = &Trie{}
		}
		pointer = pointer.children[char-'a']
	}
	pointer.endOfWord = true
}

func (trie *Trie) Search(word string) bool {
	pointer := trie
	for _, char := range word {
		if pointer.children[char-'a'] == nil {
			return false
		}
		pointer = pointer.children[char-'a']
	}
	return pointer.endOfWord
}

func (trie *Trie) StartsWith(prefix string) bool {
	pointer := trie
	for _, char := range prefix {
		if pointer.children[char-'a'] == nil {
			return false
		}
		pointer = pointer.children[char-'a']
	}
	return true
}
