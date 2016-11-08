package controller

import (
	"main/com/codeconveyor/MorseDecoder/model"
	"main/com/codeconveyor/MorseDecoder/view"
)

func Decode() {
	morseCode := view.GetMorseCodeFromConsole()
	model.SeparateLeters(morseCode)
}
