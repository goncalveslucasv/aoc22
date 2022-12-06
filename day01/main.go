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
		fmt.Printf("error %s", err)
	}

	counter := sortDescending(countCalories(parseElves(parseInput(content))))

	fmt.Println(counter[0])
	fmt.Println(sum(counter[0:3]))
}

func parseElves(input []string) map[int][]string {
	elves := map[int][]string{}
	for k, v := range input {
		elves[k] = strings.Split(string(v), "\n")
	}

	return elves
}

func parseInput(content []byte) []string {
	return strings.Split(string(content), "\n\n")
}

func countCalories(content map[int][]string) []int {
	var calories []int
	for _, v := range content {
		var value int = 0
		for _, t := range v {
			a, _ := strconv.Atoi(t)
			value = value + a
		}
		calories = append(calories, value)
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
