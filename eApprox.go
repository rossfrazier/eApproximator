package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(eDigits(60, 1000))

}

func eDigits(precisionFactor int, stringLength int) string {
	array := make([]int64, precisionFactor)
	for i := 0; i < precisionFactor; i++ {
		var t int64 = int64(i)
		array[i] = t + 1
	}

	var digitsToMultiply []int64
	var e *big.Rat = big.NewRat(1, 1)

	for i := 1; i <= precisionFactor; i++ {
		digitsToMultiply = array[0:i]
		var digitsProduct int64 = 1
		for x := 0; x < len(digitsToMultiply); x++ {
			digitsProduct = digitsProduct * digitsToMultiply[x]
		}
		var finalProduct *big.Rat = big.NewRat(1, digitsProduct)
		e.Add(e, finalProduct)
	}

	return e.FloatString(stringLength)
}
