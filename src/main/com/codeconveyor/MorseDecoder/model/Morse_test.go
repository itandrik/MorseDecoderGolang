package model
import (
	"testing"
	"strings"
)


func TestDecode(t *testing.T){
	strFromFile := GetStringFromFile(
		"src/main/com/codeconveyor/MorseDecoder/resources/test.txt")

	strDecoded :=
		SeparateLeters(
			GetStringFromFile(
				"src/main/com/codeconveyor/MorseDecoder/resources/morse.txt"));

	if(strings.EqualFold(strDecoded, strFromFile)){
		t.Error("Error comparing")
	}
}