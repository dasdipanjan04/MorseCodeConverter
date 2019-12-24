package morsedictionary

import (
	internal "morseConverter/morseConvert/internal"
	"strconv"
	"strings"
)

func convertAny(text string) (morseCode string) {
	for eachLetter := range text {
		morseCode += internal.MorsedictionaryMap[string(text[eachLetter])] + internal.GlobalMorseSeparator
	}
	return
}

// SimpleTextToMorseConvert ... convert text to morse
func SimpleTextToMorseConvert(simpleText string) (morseCode string) {
	simpleText = strings.ToUpper(simpleText)
	morseCode = convertAny(simpleText)
	return
}

// SimpleMorseToTextConvert ... convert morse to text
func SimpleMorseToTextConvert(morseCode string) (simpleText string) {
	individualMorseCollections := strings.Split(morseCode, internal.GlobalMorseSeparator)
	for eachMorseCode := range individualMorseCollections {
		simpleText += internal.MorseToTextDictionaryMap[string(individualMorseCollections[eachMorseCode])]
	}
	return
}

// SimpleIntToMorseConvert ... convert int to morse
func SimpleIntToMorseConvert(number int) (morseCode string) {
	numStringyfy := strconv.Itoa(number)
	morseCode = convertAny(numStringyfy)
	return
}
