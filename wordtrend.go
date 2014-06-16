package wordtrend

import (
	"sync"
	"strings"
	"regexp"
	"log"
	"github.com/agonopol/go-stem/stemmer"
)

var trending = make(map[string]int)

var mutex = &sync.Mutex{}

const alpha = "[^A-Za-z]+"

var reg = regexp.MustCompile(alpha)

func setStemAndCount(word string) {
	cleanedWord := reg.ReplaceAllString(word, "")

	w := string(stemmer.Stem([]byte(cleanedWord)))

	mutex.Lock()
	trending[w] = trending[w] + 1
	log.Println("saving ", w, trending[w])
	mutex.Unlock()
}

func SetWords(sentence string) []string {
	words := strings.Fields(sentence)
	wordLen := len(words)

	for i := 0; i <= wordLen - 1; i++ {
		setStemAndCount(words[i])
	}

	return words
}

func GetTrending() map[string]int {
	log.Println(trending)
	return trending
}
