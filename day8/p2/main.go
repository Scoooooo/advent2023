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
	q := []string{}

	for i := 2; i < len(lines); i++{
		maze[lines[i][0:3]] = next{lines[i][7:10], lines[i][12:15]}
		if lines[i][2] == 'A' {
			q = append(q, lines[i][0:3])
		}
	}


	steps := 0
	index := 0
	loopL := []int{}
	for len(q) != 0 {

		nextQ := []string{}
		for i:= 0; i < len(q); i++{

			nextStep  := ""
			if path[index] == 'L'{
				nextStep = maze[q[i]].l
			}  else {
				nextStep = maze[q[i]].r
			}
			if nextStep[2] == 'Z' { 
				loopL = append(loopL, steps + 1)
			} else {

				nextQ = append(nextQ, nextStep)
			}

		}
		q = nextQ
		index += 1
		steps += 1
		if index == len(path){
			index = 0
		}


	} 

	allEqual := func() bool {
		for i:= 0; i < len(loopL); i++{
			if loopL[0] != loopL[i]{
				return false
			}
		}

		return true
	}

	fmt.Println(loopL)
	// find lcm of the loop numbers

}
