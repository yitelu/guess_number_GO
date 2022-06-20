package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	secret := getRandomNum()

	guess := getUserInput()

	//println(secret, guess)

	for guess != secret {
		if guess < secret {
			fmt.Print("it's too small, enter again! ")
			guess = getUserInput()
		} else {
			fmt.Print("it's too big, enter again! ")
			guess = getUserInput()
		}

	}
	fmt.Print("you got it right")

}

func getUserInput() int {
	fmt.Println("Please input your guess: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("you guessed: ", input)
	}
	return input
}

func getRandomNum() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}
