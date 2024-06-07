package leetcode

import (
	"testing"
)

type parameters struct {
	wordsToInsert []string
	searchWords   []string
	prefixes      []string
}

func TestTrie(t *testing.T) {
	tests := []struct {
		parameters
		expectedSearchResults []bool
		expectedPrefixResults []bool
	}{
		{
			parameters{
				wordsToInsert: []string{"apple", "app"},
				searchWords:   []string{"apple", "app", "ap", "banana"},
				prefixes:      []string{"app", "ap", "ban"},
			},
			[]bool{true, true, false, false},
			[]bool{true, true, false},
		},
		{
			parameters{
				wordsToInsert: []string{"banana", "band", "bee"},
				searchWords:   []string{"banana", "band", "bee", "ban"},
				prefixes:      []string{"ban", "ba", "b"},
			},
			[]bool{true, true, true, false},
			[]bool{true, true, true},
		},
		{
			parameters{
				wordsToInsert: []string{"cat", "car", "cart"},
				searchWords:   []string{"cat", "car", "cart", "ca"},
				prefixes:      []string{"car", "ca", "c"},
			},
			[]bool{true, true, true, false},
			[]bool{true, true, true},
		},
	}

	for _, test := range tests {
		t.Run("Test Trie", func(t *testing.T) {
			trie := Constructor()
			for _, word := range test.parameters.wordsToInsert {
				trie.Insert(word)
			}
			for i, word := range test.parameters.searchWords {
				result := trie.Search(word)
				if result != test.expectedSearchResults[i] {
					t.Errorf("Search(%s) got %v, want %v", word, result, test.expectedSearchResults[i])
				}
			}
			for i, prefix := range test.parameters.prefixes {
				result := trie.StartsWith(prefix)
				if result != test.expectedPrefixResults[i] {
					t.Errorf("StartsWith(%s) got %v, want %v", prefix, result, test.expectedPrefixResults[i])
				}
			}
		})
	}
}
