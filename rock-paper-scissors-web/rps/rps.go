package rps

import(
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW 		 = 3
)

func PlayRound(PlayerValue int)(int, string, string){
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	ComputerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case ROCK:
		ComputerChoice = "Computer choice ROCKS"
		break
	case PAPER:
		ComputerChoice = "Computer choice PAPER"
		break
	case SCISSORS:
		ComputerChoice = "Computer choice SCISSORS"
		break
	default:
	}

	if PlayerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
	} else if PlayerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer wins!"
		winner = COMPUTERWINS
	}

	return winner, ComputerChoice, roundResult
}