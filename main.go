package main

import (
	"coolGame/lib"
	"fmt"
	"time"
)

func main() {
	player := lib.NewPlayer("jack")
	fmt.Println(player)
	boss := lib.NewBoss()
	battle := lib.NewBattle(boss, player)
	fmt.Println(battle.Boss)
	// battleDone := make(chan bool)

	t := time.Tick(1 * time.Second)
	i := 0
	for now := range t {
		fmt.Println(i, now)
		i += 1
	}

}
