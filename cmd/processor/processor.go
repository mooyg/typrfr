package processor

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Word struct {
	ExpectedText string
	ProvidedText []byte
	Letters      []string
}

func ConstructWords(s string) []Word {
	words := strings.Split(s, " ")
	constructedWords := make([]Word, 0, len(words))
	for _, v := range words {
		constructedWords = append(constructedWords, Word{
			ExpectedText: v,
			ProvidedText: []byte(""),
			Letters:      strings.Split(v, ""),
		})
	}
	return constructedWords
}

func ProcessTyping(event *tcell.EventKey, words []Word, currIndex int) {
	words[currIndex].ProvidedText = append(words[currIndex].ProvidedText, []byte(event.Name())...)
}
