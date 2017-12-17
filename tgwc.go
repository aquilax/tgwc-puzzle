package tgwc

import (
	"sort"
	"strings"
)

// Row contains possible word combination
type Row struct {
	before string
	after  string
}

// Letter represents a letter of the puzle of possible combinations for words
type Letter struct {
	ltr  rune
	rows []Row
}

// Puzzle represents new puzzle
type Puzzle []Letter

// Generate creates new puzzle, given word and dictionary of words
func Generate(word string, dictionary []string) (*Puzzle, error) {
	puzzle := make(Puzzle, len(word))
	runes := []rune(word)
	sort.Sort(sort.StringSlice(dictionary))
	for i, r := range runes {
		puzzle[i] = Letter{
			r,
			findMatches(r, dictionary),
		}
	}
	return &puzzle, nil
}

func findMatches(r rune, dictionary []string) []Row {
	var result []Row
	var wordIndexes []int
	for i, w := range dictionary {
		if strings.IndexRune(w, r) > -1 {
			wordIndexes = append(wordIndexes, i)
		}
	}
	for _, i := range wordIndexes {
		for p, l := range []rune(dictionary[i]) {
			if l == r {
				left := dictionary[i][:p]
				right := dictionary[i][p+1:]
				l := sort.SearchStrings(dictionary, left)
				if dictionary[l] == left {
					r := sort.SearchStrings(dictionary, right)
					if dictionary[r] == right {
						result = append(result, Row{left, right})
					}
				}
			}
		}
	}

	return result
}
