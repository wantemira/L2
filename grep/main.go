package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	after      = flag.Int("A", 0, "print N lines after match")
	before     = flag.Int("B", 0, "print N lines before match")
	context    = flag.Int("C", 0, "print N lines before and after match")
	count      = flag.Bool("c", false, "print only count of matching lines")
	ignoreCase = flag.Bool("i", false, "ignore case")
	invert     = flag.Bool("v", false, "invert match")
	fixed      = flag.Bool("F", false, "fixed string match")
	lineNum    = flag.Bool("n", false, "print line numbers")
)

func main() {
	flag.Parse()

	if *context > 0 {
		*after = *context
		*before = *context
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "pattern required")
		os.Exit(1)
	}

	pattern := args[0]

	var input *os.File
	if len(args) > 1 {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	var re *regexp.Regexp
	if !*fixed {
		if *ignoreCase {
			pattern = "(?i)" + pattern
		}
		r, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		re = r
	} else if *ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	scanner := bufio.NewScanner(input)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	match := func(line string) bool {
		if *fixed {
			if *ignoreCase {
				line = strings.ToLower(line)
			}
			return strings.Contains(line, pattern)
		}
		return re.MatchString(line)
	}

	matches := make([]bool, len(lines))
	countMatches := 0

	for i, line := range lines {
		m := match(line)
		if *invert {
			m = !m
		}
		matches[i] = m
		if m {
			countMatches++
		}
	}

	if *count {
		fmt.Println(countMatches)
		return
	}

	printed := make([]bool, len(lines))

	for i, m := range matches {
		if !m {
			continue
		}

		start := i - *before
		if start < 0 {
			start = 0
		}

		end := i + *after
		if end >= len(lines) {
			end = len(lines) - 1
		}

		for j := start; j <= end; j++ {
			if printed[j] {
				continue
			}

			if *lineNum {
				fmt.Printf("%d:%s\n", j+1, lines[j])
			} else {
				fmt.Println(lines[j])
			}

			printed[j] = true
		}
	}
}
