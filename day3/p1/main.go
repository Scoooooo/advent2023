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

	currVal := 0

	prvLine := ""
	scanner.Scan()
	currLine := scanner.Text()
	nextLine := "" 

	minR := '0'
	maxR := '9'

	notDigitOrComma := func(b byte) bool{
		return (rune(b) < (minR) || rune(b)  > (maxR)) && b != '.'
	} 

	hasAdjacentElement := func(s int, e int) bool {

		arr := []int{} 

		for i := max(0, s-1); i <= min(len(currLine) - 1, e + 1); i++{
			arr = append(arr, i)
		}
		if prvLine != "" {
			for i := 0; i < len(arr); i++{
				if notDigitOrComma(prvLine[arr[i]]) {
					return true
				}
			} 
		}

		if nextLine != "" {
			for i := 0; i < len(arr); i++{
				if notDigitOrComma(nextLine[arr[i]]){
					return true
				}
			} 
		}

		if s - 1 >= 0{
			if notDigitOrComma(currLine[s - 1]){
				return true
			}
		}

		if e + 1 <  len(currLine){
			if notDigitOrComma(currLine[e + 1]){
				return true
			}
		}


		return false 
	}

	num  := func(str string, s int, e int) int {
		if str != "" {
			if hasAdjacentElement(s, e) {
				ret, _ := strconv.Atoi(str)
				return ret
			}
		}
		return 0

	}

	for scanner.Scan() {
		nextLine = scanner.Text()

		currNum := ""
		for i, val := range currLine{
			if val >= minR && val <= maxR{
				currNum += string(val) 

			} else {
				currVal += num(currNum, i - len(currNum), i - 1)
				currNum = ""
			}
		}

		currVal += num(currNum, len(currLine) - len(currNum), len(currLine) - 1)
		currNum = ""

		prvLine = currLine
		currLine = nextLine
	}

	nextLine = ""

	currNum := ""
	for i, val := range currLine{
		if val >= minR && val <= maxR{
			currNum += string(val) 

		} else {
			currVal += num(currNum, i - len(currNum), i - 1)
			currNum = ""
		}
	}

	currVal += num(currNum, len(currLine) - len(currNum), len(currLine) - 1)
	currNum = ""

	fmt.Println(currVal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
