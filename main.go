package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// A string flag.
	wordMap := flag.String("m", "", "Map that uses '.' for unknown letters")
	extraLetters := flag.String("e", "", "Extra letters that are in the word, but you don't know their exact location.")
	excludedLetters := flag.String("x", "", "Letters that should be excluded, because you know they aren't in the word.")
	version := flag.Bool("v", false, "Display version information")
	flag.Parse()

	if *version {
		fmt.Println("Version: 0.0.2")
		return
	}

	// Get the word map.
	rx := strings.ReplaceAll(*wordMap, "*", ".")

	// See if there are any extra letters
	ex := *extraLetters

	// See if there are any excluded letters
	xx := *excludedLetters

	words, err := readList("wordlist.txt")
	if err != nil {
		fmt.Println(err)
	}

	results := getMatches(words, rx, ex, xx)

	// for _, word := range results {
	// 	fmt.Println(word)
	// }

	colorizeOutput(results, rx, ex)
}

func colorizeOutput(results []string, pattern string, extras string) {
	colorReset := "\033[0m"

	// colorRed := "\033[31m"
	// colorGreen := "\033[32m"
	// colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	// colorCyan := "\033[36m"
	// colorWhite := "\033[37m"

	for _, word := range results {
		for i, v := range word {
			if strings.Contains(extras, string(v)) {
				fmt.Print(colorPurple)
			}

			if string(v) == string(pattern[i]) {
				fmt.Print(colorBlue)
			}
			fmt.Print(string(v), colorReset)
		}
		fmt.Print("\n")
	}
}

// getMatches returns a list of all
func getMatches(wordList []string, pattern string, extras string, excluded string) []string {

	var results []string

	for _, word := range wordList {
		if regexMatch(pattern, word) {

			// Check for the extra letters
			valid := true
			for _, v := range extras {
				if !regexMatch(string(v), word) {
					valid = false
				}
			}

			for _, v := range excluded {
				if regexMatch(string(v), word) {
					valid = false
				}
			}

			if valid {
				results = append(results, word)
			}
		}
	}

	return results
}

// readList reads the wordlist file and returns an array of strings.
func readList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// regexMatch returns true if the testString matches the regex pattern.
func regexMatch(pattern string, testString string) bool {
	m, err := regexp.MatchString(pattern, testString)
	if err != nil {
		fmt.Println("your regex is faulty")
	}

	if m {
		return true
	} else {
		return false
	}
}
