package main

import (
	"coolGame/lib"
	"fmt"
	"math"
)

func main() {
	player := lib.NewPlayer("jack")
	fmt.Println(player)
	fmt.Println(1 / 50)
	fmt.Println(math.Ceil(float64((1.0 / 50) * 100)))
}
