package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	column  = flag.Int("k", 0, "sort by column")
	numeric = flag.Bool("n", false, "numeric sort")
	reverse = flag.Bool("r", false, "reverse sort")
	unique  = flag.Bool("u", false, "unique lines")
	month   = flag.Bool("M", false, "sort by month")
	ignoreB = flag.Bool("b", false, "ignore leading blanks")
	check   = flag.Bool("c", false, "check if sorted")
	human   = flag.Bool("h", false, "human numeric sort")
)

var months = map[string]int{
	"Jan": 1, "Feb": 2, "Mar": 3,
	"Apr": 4, "May": 5, "Jun": 6,
	"Jul": 7, "Aug": 8, "Sep": 9,
	"Oct": 10, "Nov": 11, "Dec": 12,
}

func main() {
	flag.Parse()

	lines, err := readLines()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *unique {
		lines = uniqueLines(lines)
	}

	less := func(i, j int) bool {
		a := getKey(lines[i])
		b := getKey(lines[j])

		var result bool

		switch {
		case *numeric:
			ai, _ := strconv.ParseFloat(a, 64)
			bi, _ := strconv.ParseFloat(b, 64)
			result = ai < bi

		case *human:
			result = parseHuman(a) < parseHuman(b)

		case *month:
			result = months[a] < months[b]

		default:
			result = a < b
		}

		if *reverse {
			return !result
		}

		return result
	}

	if *check {
		if !sort.SliceIsSorted(lines, less) {
			fmt.Println("data is not sorted")
			os.Exit(1)
		}
		return
	}

	sort.Slice(lines, less)

	for _, l := range lines {
		fmt.Println(l)
	}
}

func readLines() ([]string, error) {
	var scanner *bufio.Scanner

	if len(flag.Args()) > 0 {
		file, err := os.Open(flag.Args()[0])
		if err != nil {
			return nil, err
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func uniqueLines(lines []string) []string {
	seen := make(map[string]struct{})
	var res []string

	for _, l := range lines {
		if _, ok := seen[l]; !ok {
			seen[l] = struct{}{}
			res = append(res, l)
		}
	}

	return res
}

func getKey(line string) string {
	if *ignoreB {
		line = strings.TrimLeft(line, " ")
	}

	if *column <= 0 {
		return line
	}

	fields := strings.Fields(line)

	if *column-1 < len(fields) {
		return fields[*column-1]
	}

	return ""
}

func parseHuman(s string) float64 {
	mult := 1.0

	switch {
	case strings.HasSuffix(s, "K"):
		mult = 1e3
		s = strings.TrimSuffix(s, "K")
	case strings.HasSuffix(s, "M"):
		mult = 1e6
		s = strings.TrimSuffix(s, "M")
	case strings.HasSuffix(s, "G"):
		mult = 1e9
		s = strings.TrimSuffix(s, "G")
	}

	v, _ := strconv.ParseFloat(s, 64)
	return v * mult
}
