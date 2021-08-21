package main

import (
	"Seven_pokers/handler"
	"fmt"
	"time"
)

func main () {

	data := handler.ReadDataToModel("./seven_cards_with_ghost.result.json")
	start := time.Now()
	for i := range data {
		turn := handler.CreateTurn(&data[i])
		if turn.Alice.Pokers[6].Face == 0{continue}
		if turn.Bob.Pokers[6].Face == 0{continue}
		handler.Analyse(&turn)
		handler.JudgeWinner(&turn)
		if turn.Winner != data[i].Result {
			fmt.Println("-----Error-----")
			fmt.Println("position:", i)
			fmt.Println(data[i].Alice)
			fmt.Println("Alice.Pokers:")
			fmt.Println(turn.Alice.Pokers)
			fmt.Println("Alice.Decks:")
			fmt.Println(turn.Alice.Deck)
			fmt.Println("Alice.Level:")
			fmt.Println(turn.Alice.Level,turn.Alice.Finish)
			fmt.Println("-------------------------------------------")
			fmt.Println(data[i].Bob)
			fmt.Println("Bob.Pokers:")
			fmt.Println(turn.Bob.Pokers)
			fmt.Println("Bob.Decks:")
			fmt.Println(turn.Bob.Deck)
			fmt.Println("Bob.Level:")
			fmt.Println(turn.Bob.Level,turn.Bob.Finish)
			fmt.Println("your winner is :",turn.Winner,"Should be :",data[i].Result)
			panic("result wrong")
		}
	}
	cost := time.Since(start)
	fmt.Println("成功！ 耗时：", cost)
}

