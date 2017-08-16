package tokenizer

import (
	"regexp"
	"strings"
)

// Tokenizer model
type Tokenizer struct {
	Patern string
	Tokens []string
}

// NewTokenizer create Tokenizer
func NewTokenizer(patern ...string) Tokenizer {
	var regexpPatern string
	if len(patern) != 0 {
		regexpPatern = patern[0]
	} else {
		regexpPatern = `([A-z]){2,}\w+`
	}
	t := Tokenizer{Patern: regexpPatern}
	//fmt.Println("t", t)
	return t
}

// Tokenize all words in the text without punctuation marks. Min length of word it is 3 letters. This method use stopwords
func (t *Tokenizer) Tokenize(text string) []string {

	r := regexp.MustCompile(t.Patern)
	words := r.FindAllString(text, -1)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	t.Tokens = words
	return words
}

// HandleStopWords func
func (t *Tokenizer) HandleStopWords(customeStopWords ...string) {
	// for wordIndex := 0; wordIndex < len(words); wordIndex++ {
	// 	if strsEquals := strings.Compare(pair1, pair2); strsEquals == 0 {
	// 		words = append(words[:j], words[j+1:]...)
	// 	}
	// }
}
