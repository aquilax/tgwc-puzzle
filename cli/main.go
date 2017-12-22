package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	tgwc "github.com/aquilax/tgwc-puzzle"
)

type solution struct {
	left   string
	middle string
	right  string
}

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
		if len(w) > 2 && strings.IndexRune(w, '\'') == -1 {
			dictionary = append(dictionary, strings.ToLower(w))
		}
	}
	puzzle := tgwc.Generate(strings.TrimSpace(os.Args[1]), dictionary)
	if os.Args[3] == "random" {
		// Generate single random soluton
		rand.Seed(time.Now().Unix())
		var sol []solution
		leftLen := 0
		for _, l := range puzzle {
			if len(l.Rows) > 0 {
				r := l.Rows[rand.Intn(len(l.Rows))]
				if len([]rune(r.Before)) > leftLen {
					leftLen = len([]rune(r.Before))
				}
				sol = append(sol, solution{
					r.Before,
					string(l.Ltr),
					r.After,
				})
			}
		}
		for _, s := range sol {
			fmt.Printf("%s%s %s %s\n",
				strings.Repeat(" ", leftLen-len([]rune(s.left))),
				s.left,
				s.middle,
				s.right)
		}
		return
	}
	for _, l := range puzzle {
		fmt.Println(string(l.Ltr))
		for _, r := range l.Rows {
			fmt.Printf("\t(%s%s%s) %s | %s | %s\n", r.Before, string(l.Ltr), r.After, r.Before, string(l.Ltr), r.After)
		}
	}
}
