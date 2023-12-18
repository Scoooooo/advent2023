package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

)

type next struct{
	l string 
	r string 
}


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

	maze := map[string]next{}
	path := lines[0]

	for i := 2; i < len(lines); i++{
		maze[lines[i][0:3]] = next{lines[i][7:10], lines[i][12:15]}
	}


	var step func(curr string, index int) int 

	step = func(curr string, index int) int {
		if curr == "ZZZ" {
			return 0
		}
		if index == len(path){
			index = 0
		}

		nextStep  := ""
		if path[index] == 'L'{
			nextStep = maze[curr].l
		}  else {
			nextStep = maze[curr].r
		}
		return step(nextStep, index + 1) + 1
	}

	fmt.Println(step("AAA", 0))
	
}
