package main

import (
	"main/com/codeconveyor/MorseDecoder/controller"

	"time"
	"math/big"
	"fmt"
	"log"
)
func main() {
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	controller.Decode()

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}