package main

import (
	"fmt"
	mrsConv "morseConverter/morseConvert"
)

func main() {
	var simpleText = "hello world, how are you doing?"
	var morseCode = ".----|..---|...--|....-|.....|-....|"
	var number = 123456
	fmt.Println(mrsConv.SimpleTextToMorseConvert(simpleText))
	fmt.Println(mrsConv.SimpleMorseToTextConvert(morseCode))
	fmt.Println(mrsConv.SimpleIntToMorseConvert(number))
}
