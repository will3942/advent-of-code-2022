package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Move int

const (
	Rock     Move = 1
	Paper    Move = 2
	Scissors Move = 3
)

type DesiredOutcome int

const (
	Lose DesiredOutcome = 1
	Draw DesiredOutcome = 2
	Win  DesiredOutcome = 3
)

type Round struct {
	theirMove Move
	ourMove   Move
}

func (move Move) losingMove() Move {
	if move == Rock {
		return Scissors
	}

	if move == Paper {
		return Rock
	}

	return Paper
}

func (move Move) winningMove() Move {
	if move == Rock {
		return Paper
	}

	if move == Paper {
		return Scissors
	}

	return Rock
}

func (round Round) isDraw() bool {
	return (round.theirMove == round.ourMove)
}

func (round Round) isLoss() bool {
	if round.theirMove == Rock && round.ourMove == Scissors {
		return true
	}

	if round.theirMove == Paper && round.ourMove == Rock {
		return true
	}

	if round.theirMove == Scissors && round.ourMove == Paper {
		return true
	}

	return false
}

func (round Round) isWin() bool {
	return !round.isLoss() && !round.isDraw()
}

func (round Round) ourScore() int {
	if round.isDraw() {
		return int(round.ourMove) + 3
	}

	if round.isLoss() {
		return int(round.ourMove)
	}

	return int(round.ourMove) + 6
}

func parseMove(raw string) (m Move, err error) {
	switch raw {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	}

	return m, errors.New("invalid move provided")
}

func parseDesiredOutcome(raw string) (d DesiredOutcome, err error) {
	switch raw {
	case "X":
		return Lose, nil
	case "Y":
		return Draw, nil
	case "Z":
		return Win, nil
	}

	return d, errors.New("invalid outcome provided")
}

func calcOurMove(theirMove Move, outcome DesiredOutcome) Move {
	if outcome == Draw {
		return theirMove
	}

	if outcome == Win {
		return theirMove.winningMove()
	}

	return theirMove.losingMove()
}

func parseRoundPart1(line string) (round Round, err error) {
	moves := strings.Split(line, " ")

	maybeTheirMove, err := parseMove(moves[0])

	if err != nil {
		return round, err
	}

	maybeOurMove, err := parseMove(moves[1])

	if err != nil {
		return round, err
	}

	return Round{theirMove: maybeTheirMove, ourMove: maybeOurMove}, nil
}

func parseRoundPart2(line string) (round Round, err error) {
	input := strings.Split(line, " ")

	maybeTheirMove, err := parseMove(input[0])

	if err != nil {
		return round, err
	}

	maybeOurDesiredOutcome, err := parseDesiredOutcome(input[1])

	if err != nil {
		return round, err
	}

	ourMove := calcOurMove(maybeTheirMove, maybeOurDesiredOutcome)

	return Round{theirMove: maybeTheirMove, ourMove: ourMove}, nil
}

func parseInputFile(path string, isPart1 bool) []Round {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	rounds := []Round{}

	for scanner.Scan() {
		line := scanner.Text()

		var round Round

		if isPart1 {
			maybeRound, err := parseRoundPart1(line)

			if err != nil {
				log.Fatal(err)
			}

			round = maybeRound
		} else {
			maybeRound, err := parseRoundPart2(line)

			if err != nil {
				log.Fatal(err)
			}

			round = maybeRound
		}

		rounds = append(rounds, round)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rounds
}

func calcOurTotalScore(rounds []Round) int {
	total := 0

	for _, round := range rounds {
		total += round.ourScore()
	}

	return total
}

func main() {
	rounds := parseInputFile("./2-input.txt", false)

	fmt.Println("Num rounds: ", len(rounds))

	ourTotalScore := calcOurTotalScore(rounds)

	fmt.Println("Our Total Score: ", ourTotalScore)
}
