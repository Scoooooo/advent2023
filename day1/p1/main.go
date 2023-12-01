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

	digits := map[rune]string{'0': "0", '1': "1", '2': "2", '3': "3", '4': "4", '5': "5", '6': "6", '7': "7", '8': "8", '9': "9"}

	currVal := 0
	for scanner.Scan() {

		num := ""
		line := scanner.Text()
		for _, c := range line {
			if val, ok := digits[c]; ok {
				num += val
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if val, ok := digits[rune(line[i])]; ok {
				num += val
				break
			}
		}
		val, _ := strconv.Atoi(num)
		currVal += val
	}
	fmt.Println(currVal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
