package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	currVal := 0
	cards := map[int]int{}
	
	addCards := func(s1 []string, s2 []string, pos int, curr int) {
		for i := 0; i < len(s2); i++{
			if s2[i] != "" && slices.Contains(s1, s2[i]){
				cards[pos] += curr	
				pos += 1
			}

		}
	}


	for i := 0; i < len(lines); i++{
		split := strings.Split(lines[i], ":")
		arrs := strings.Split(split[1], "|")
		arr1 := strings.Split(arrs[0], " ")
		arr2 := strings.Split(arrs[1], " ")
		addCards(arr1, arr2, i + 1, cards[i] + 1)
		currVal += cards[i] + 1
	} 

	fmt.Println(currVal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
