package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/shreyesh0610/ffai-interview/utils"
)

var MAX_WORKERS int = 10

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// time.Sleep(10 * time.Minute)

	for j := range jobs {
		fmt.Printf("Worker %d started job %s\n", id, j)
		utils.ProcessEssay(j)
		fmt.Printf("Worker %d completed job %s\n", id, j)
	}

	fmt.Println("Ending Worker")
}

func main() {
	fmt.Println("Start")
	essayLinks, err := utils.ReadLines("files/endg-urls")
	if err != nil {
		fmt.Printf("Error >> ReadLines: %+v", err)
	}

	bankWords, err := utils.ReadLines("files/words.txt")
	if err != nil {
		fmt.Printf("Error >> ReadLines: %+v", err)
	}
	utils.UpdateBankWordsMap(bankWords)

	fmt.Printf("Total Essay Links: %d\n", len(essayLinks))
	fmt.Printf("Total Bank Words: %d\n", len(bankWords))

	jobs := make(chan string, len(essayLinks))
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= MAX_WORKERS; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for _, essayLink := range essayLinks {
		jobs <- essayLink
	}

	close(jobs)
	wg.Wait()

	wf := utils.GetWordFrequency()

	// Convert map to slice of structs
	wordCounts := []WordCount{}
	for word, freq := range wf.FreqMap {
		wordCounts = append(wordCounts, WordCount{
			Word:  word,
			Count: freq,
		})
	}

	// Sort the slice by count in descending order
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	// Print the top words
	top10 := []WordCount{}
	fmt.Println("Top 10 Words:")
	for i, wordCount := range wordCounts {
		if i == 10 {
			break
		}
		top10 = append(top10, wordCount)
	}

	prettyJSON, err := json.MarshalIndent(top10, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
