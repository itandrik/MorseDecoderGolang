package model

import (
	"strings"
	"fmt"
)

var MORSE_TABLE = map[string]string{
	""	: " ",
	".-"	: "A",
	"-..."	: "B",
	"-.-."	: "C",
	"-.."	: "D",
	"."	: "E",
	"..-."	: "F",
	"--."	: "G",
	"...."	: "H",
	".."	: "I",
	".---"	: "J",
	"-.-"	: "K",
	".-.."	: "L",
	"--"	: "M",
	"-."	: "N",
	"---"	: "O",
	".--."	: "P",
	"--.-"	: "Q",
	".-."	: "R",
	"..."	: "S",
	"-"	: "T",
	"..-"	: "U",
	"...-"	: "V",
	".--"	: "W",
	"-..-"	: "X",
	"-.--"	: "Y",
	"--.."	: "Z",
	"-----"	: "0",
	".----"	: "1",
	"..---"	: "2",
	"...--"	: "3",
	"....-"	: "4",
	"....."	: "5",
	"-...."	: "6",
	"--..."	: "7",
	"---.."	: "8",
	"----."	: "9",
}

func SeparateLeters(morseCode string){
	morseCode = strings.Replace(morseCode,"\n"," \n",-1)
	var result []string = strings.Split(morseCode," ")
	for _, i:= range result{
		fmt.Print(MORSE_TABLE[i])
	}

}

