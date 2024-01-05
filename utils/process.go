package utils

import (
	"fmt"
	"strings"
	"sync"
)

type WordFrequency struct {
	FreqMap map[string]int
	mutex   sync.RWMutex
}

var wf *WordFrequency

func init() {
	wf = &WordFrequency{
		FreqMap: make(map[string]int),
	}
}

func GetWordFrequency() *WordFrequency {
	return wf
}

func (wf *WordFrequency) Set(word string, value int) {
	wf.mutex.Lock()
	defer wf.mutex.Unlock()
	wf.FreqMap[word] = value
}

func (wf *WordFrequency) Inc(word string, inc_value int) {
	wf.mutex.Lock()
	defer wf.mutex.Unlock()
	if _, ok := wf.FreqMap[word]; ok {
		wf.FreqMap[word] += inc_value
	} else {
		wf.FreqMap[word] = inc_value
	}
}

func (wf *WordFrequency) Get(key string) int {
	wf.mutex.RLock()
	defer wf.mutex.RUnlock()
	return wf.FreqMap[key]
}

func PreProcessWord(word string) string {
	punctuation := ",.!?:"
	word = strings.Trim(word, punctuation)
	return strings.ToLower(word)
}
func ProcessEssay(essayLink string) {
	essayText, err := ScrapeURL(essayLink)
	if err != nil {
		fmt.Printf("Error >> ScrapeURL: %+v\n", err)
		return
	}

	words := strings.Fields(essayText)

	for _, word := range words {
		word = PreProcessWord(word)
		if ValidateWord(word, bankWordsMap) {
			wf.Inc(word, 1)
		}
	}
}
