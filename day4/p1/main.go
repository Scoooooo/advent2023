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

	lineScore := func(s1 []string, s2 []string) int {

		curr := 0

		for i := 0; i < len(s2); i++{
			if s2[i] != "" && slices.Contains(s1, s2[i]){
				if curr == 0{
					curr = 1
				} else {
					curr *= 2
				}
			}

		}
		return curr
	}

	for i := 0; i < len(lines); i++{
		split := strings.Split(lines[i], ":")
		arrs := strings.Split(split[1], "|")
		arr1 := strings.Split(arrs[0], " ")
		arr2 := strings.Split(arrs[1], " ")
		currVal += lineScore(arr1, arr2)
	} 

	fmt.Println(currVal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
