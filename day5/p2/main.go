package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	seeds := []int{}


	stringSeeds := strings.Split(strings.Split(lines[0], ":")[1], " ") 

	for _, val := range stringSeeds {
		if val == ""{ continue }

		num, _ := strconv.Atoi(val)
		seeds = append(seeds, num)
	}

	fmt.Println(minVal)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
