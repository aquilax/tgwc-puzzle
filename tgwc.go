package tgwc

import (
	"sort"
	"strings"
)

// Row contains possible word combination
type Row struct {
	Before string
	After  string
}

// Letter represents a letter of the puzle of possible combinations for words
type Letter struct {
	Ltr  rune
	Rows []Row
}

// Puzzle represents new puzzle
type Puzzle []Letter

// Generate creates new puzzle, given word and dictionary of words
func Generate(word string, dictionary []string) Puzzle {
	puzzle := make(Puzzle, len(word))
	runes := []rune(word)
	sort.Sort(sort.StringSlice(dictionary))
	for i, r := range runes {
		puzzle[i] = Letter{
			r,
			findMatches(r, dictionary),
		}
	}
	return puzzle
}

func findMatches(r rune, dictionary []string) []Row {
	var result []Row
	var wordIndexes []int
	for i, w := range dictionary {
		if strings.IndexRune(w, r) > -1 {
			wordIndexes = append(wordIndexes, i)
		}
	}
	ld := len(dictionary)
	for _, i := range wordIndexes {
		word := []rune([]rune(dictionary[i]))
		for p, l := range word {
			if l == r {
				left := string(word[:p])
				right := string(word[p+1:])
				le := sort.SearchStrings(dictionary, left)
				if le < ld && dictionary[le] == left {
					ri := sort.SearchStrings(dictionary, right)
					if ri < ld && dictionary[ri] == right {
						result = append(result, Row{left, right})
					}
				}
			}
		}
	}

	return result
}
