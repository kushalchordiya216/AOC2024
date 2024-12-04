package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kushalchordiya216/AOC2024/common"
	"github.com/kushalchordiya216/AOC2024/day1"
	"github.com/kushalchordiya216/AOC2024/day2"
	"github.com/kushalchordiya216/AOC2024/day3"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the day and part to be solved")
		return
	}
	day, err1 := strconv.Atoi(os.Args[1])
	part, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || day <= 0 || day > 25 || part < 1 || part > 2 {
		fmt.Println("Please provide a valid input")
		return
	}
	var solver common.Solver

	switch day {
	case 1:
		if part == 1 {
			solver = &day1.Part1Solver{}
		} else {
			solver = &day1.Part2Solver{}
		}
	case 2:
		if part == 1 {
			solver = &day2.Part1Solver{}
		} else {
			solver = &day2.Part2Solver{}
		}
	case 3:
		if part == 1 {
			solver = &day3.Part1Solver{}
		} else {
			solver = &day3.Part2Solver{}
		}
	}
	if err := solver.Read(fmt.Sprintf("day%d/input.txt", day)); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	result := solver.Solve()
	fmt.Printf("Day %d, Part %d solution: %d\n", day, part, result)

}
