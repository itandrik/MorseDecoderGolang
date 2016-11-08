package controller

import (
	"main/com/codeconveyor/MorseDecoder/model"
	//"main/com/codeconveyor/MorseDecoder/view"
)

func Decode() {
	morseCode := model.GetStringFromFile("src/main/com/codeconveyor/MorseDecoder/resources/morse.txt")//view.GetMorseCodeFromConsole()
	model.SeparateLeters(morseCode)
}
