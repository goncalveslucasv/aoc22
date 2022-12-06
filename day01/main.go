package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input/input.txt")
	if err != nil {
		panic("invalid input!")
	}

	elvesCalories := sortDescending(countCalories(parseElves(parseInput(content))))

	fmt.Println(elvesCalories[0])
	fmt.Println(sum(elvesCalories[0:3]))
}

func parseElves(input []string) map[int][]string {
	elves := map[int][]string{}
	for key, elf := range input {
		elves[key] = strings.Split(string(elf), "\n")
	}

	return elves
}

func parseInput(content []byte) []string {
	return strings.Split(string(content), "\n\n")
}

func countCalories(content map[int][]string) []int {
	var calories []int
	for _, elfCalories := range content {
		var totalCalories = 0
		for _, itemCalories := range elfCalories {
			itemCaloriesToSum, _ := strconv.Atoi(itemCalories)
			totalCalories = totalCalories + itemCaloriesToSum
		}
		calories = append(calories, totalCalories)
	}

	return calories
}

func sortDescending(calories []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	return calories
}

func sum(calories []int) int {
	var sum = 0
	for _, v := range calories {
		sum = sum + v
	}

	return sum
}
