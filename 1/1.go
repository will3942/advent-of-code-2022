package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Food struct {
	calories int64
}

type Elf struct {
	foods        []Food
	currCalories int64
}

func (elf Elf) sumFoodCalories() int64 {
	var calories int64

	calories = 0

	for _, food := range elf.foods {
		calories += food.calories
	}

	return calories
}

func (elf *Elf) storeFood(food Food) {
	elf.foods = append(elf.foods, food)
	elf.currCalories += food.calories
}

func parseInputFile(path string) []Elf {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	elves := []Elf{{foods: []Food{}, currCalories: 0}}
	currentElfIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elves = append(elves, Elf{foods: []Food{}, currCalories: 0})
			currentElfIndex += 1
			continue
		}

		calories, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		food := Food{calories: calories}

		elves[currentElfIndex].storeFood(food)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elves
}

func sortElvesByCalories(elves []Elf) []Elf {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].currCalories > elves[j].currCalories
	})

	return elves
}

func main() {
	elves := parseInputFile("./1-input.txt")

	fmt.Println("Num elves: ", len(elves))

	elves = sortElvesByCalories(elves)

	fmt.Println("Largest 3 elves: ", elves[0].currCalories, elves[1].currCalories, elves[2].currCalories)
}
