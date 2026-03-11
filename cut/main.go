package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFields(spec string) map[int]bool {
	fields := make(map[int]bool)

	parts := strings.Split(spec, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			r := strings.Split(part, "-")

			start, err1 := strconv.Atoi(r[0])
			end, err2 := strconv.Atoi(r[1])

			if err1 != nil || err2 != nil || start > end {
				continue
			}

			for i := start; i <= end; i++ {
				fields[i] = true
			}

		} else {
			val, err := strconv.Atoi(part)
			if err != nil {
				continue
			}
			fields[val] = true
		}
	}

	return fields
}

func main() {

	fieldsFlag := flag.String("f", "", "fields to extract")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", false, "only lines with delimiter")

	flag.Parse()

	if *fieldsFlag == "" {
		fmt.Fprintln(os.Stderr, "flag -f is required")
		os.Exit(1)
	}

	fields := parseFields(*fieldsFlag)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := scanner.Text()

		if !strings.Contains(line, *delimiter) {
			if !*separated {
				fmt.Println(line)
			}
			continue
		}

		cols := strings.Split(line, *delimiter)

		var result []string

		for i := 1; i <= len(cols); i++ {
			if fields[i] {
				result = append(result, cols[i-1])
			}
		}

		if len(result) > 0 {
			fmt.Println(strings.Join(result, *delimiter))
		}
	}
}
