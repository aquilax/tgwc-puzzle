package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tgwc "github.com/aquilax/tgwc-puzzle"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide word and path to dictionary as parameters")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var dictionary []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w := scanner.Text()
		if len(w) > 5 && strings.IndexRune(w, '\'') == -1 {
			dictionary = append(dictionary, strings.ToLower(w))
		}
	}
	puzzle := tgwc.Generate(os.Args[1], dictionary)
	for _, l := range puzzle {
		fmt.Println(string(l.Ltr))
		for _, r := range l.Rows {
			fmt.Printf("\t(%s%s%s) %s | %s | %s\n", r.Before, string(l.Ltr), r.After, r.Before, string(l.Ltr), r.After)
		}
	}
}
