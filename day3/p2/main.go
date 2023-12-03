package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	minR := '0'
	maxR := '9'

	isDigit := func(r rune) bool {
		return r >= minR && r <= maxR
	}

	findDigit := func(x int, y int) (int, int) {

		if x >= len(lines) || x < 0 {
			return -1, 0
		} 

		if y >= len(lines[0]) || y < 0{
			return -1, 0
		} 

		if !isDigit(rune(lines[x][y])){
			return -1, 0
		}

		s := y  
		for (s - 1) >= 0  && isDigit(rune(lines[x][s - 1])){
			s -= 1
		}
		e := y + 1
		for  (e) < len(lines[0]) && isDigit(rune(lines[x][e])){
			e += 1
		}
		num, _ := strconv.Atoi(lines[x][s:e])

		return num, (e - y)
	}

	gearScore := func(x int, y int) int{

		nums := []int{0, 0}
		index := 0

		for i := -1; i < 2; i++{
			for j := -1; j < 2; j++ {
				num, offset := findDigit(x + i, y + j)

				if i != 0 {
					j += offset 
				}

				if num != -1{
					if index == 2{
						return 0
					}
					nums[index] = num
					index += 1
				}

			}
		} 
		return nums[0] * nums[1]
	}


	for i := 0; i < len(lines); i++{
		for j := 1; j < len(lines[0]); j++ {
			if lines[i][j] == '*'{
				currVal += gearScore(i, j)
			}
		}
	} 


	fmt.Println(currVal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
