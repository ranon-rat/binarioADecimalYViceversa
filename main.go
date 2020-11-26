package main

import (
	"fmt"
	"math"
)

func aDecimal(bit []int) float64 {
	var decimal float64 = 0
	for i := len(bit) - 1; i >= 0; i-- {
		decimal += math.Pow(float64(bit[i]*2), float64((len(bit)-1)-i))
	}
	return decimal
}
func aBinario(decimal int) []int {
	var bits []int
	resultado := decimal
	for resultado > 1 {
		
		bits = append(bits, resultado%2)
		resultado = int(resultado / 2)

	}
	return bits
}
func main() {
	var decimal []int = []int{1, 0, 1, 0, 1, 1, 0, 1}
	fmt.Println(aDecimal(decimal))
	fmt.Println(aBinario(173))

}
