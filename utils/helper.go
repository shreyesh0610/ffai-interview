package utils

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

var bankWordsMap = make(map[string]struct{})

func UpdateBankWordsMap(bankWords []string) {
	for _, bankWord := range bankWords {
		bankWordsMap[strings.ToLower(bankWord)] = struct{}{}
	}
}

func ReadLines(filename string) (lines []string, err error) {

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		err = errors.Wrapf(err, "Error opening file")
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cleanLink := strings.TrimSpace(scanner.Text())
		if len(cleanLink) > 0 {
			lines = append(lines, cleanLink)
		}
	}

	// Check for errors during scanning
	if err = scanner.Err(); err != nil {
		err = errors.Wrapf(err, "Error reading file")
		return
	}
	return
}

func ScrapeURL(url string) (scrape_text string, err error) {

	// cmd := exec.Command("python", append([]string{pythonScript}, scriptArgs...)...)
	// cmd := exec.Command("python", append([]string{"scrape.py"}, url)...)
	// cmd := exec.Command("python", []string{"scrape.py", url}...)
	// cmd := exec.Command("/opt/anaconda3/envs/firefly/bin/python", []string{"scrape.py", url}...)
	cmd := exec.Command("python3", "scrape.py", url)

	output, err := cmd.CombinedOutput()
	if err != nil {
		err = errors.Wrapf(err, "Error Processing URL: %v", url)
		return
	}
	scrape_text = string(output)
	return
}

func isOnlyAlphabets(word string) bool {
	for _, char := range word {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func ValidateWord(word string, bankWordsMap map[string]struct{}) (is_valid bool) {

	// Check in cache
	cache = GetCache()
	is_valid, exists := cache.Get(word)
	if exists && is_valid {
		return
	}

	// Set Cache befire exiting so next time we check cache directly
	defer func() {
		cache.Set(word, is_valid)
	}()

	// Alphabetic and 3 letter check
	if !(isOnlyAlphabets(word) && len(word) >= 3) {
		return
	}

	// Check if exists in wordbank map
	_, is_valid = bankWordsMap[word]
	return
}
