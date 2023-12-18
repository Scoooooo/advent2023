package main

import (
	"bufio"
	"fmt"
	"hash/maphash"
	"log"
	"os"
	"testing/quick"
)



type cord struct{
	x int
	y int
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

	x, y := 0, 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == 'S'{
				x = i
				y = j
				break
			}

		}

	}

	up := map[byte]cord{'|' : {-1, 0}, '7' : {0, -1}, 'F' : {0, 1}}
	down := map[byte]cord{'|' : {1, 0}, 'L' : {0, 1}, 'J' : {0, -1}}
	left := map[byte]cord{'-' : {0, -1}, 'F' : {1, 0}, 'L' : {-1, 0}}
	right := map[byte]cord{'-' : {0, 1}, '7' : {1, 0}, 'J' : {-1, 0}}


	nextStep := func(curr cord , step cord ) cord{

		if step.x == -1{
			return up[lines[curr.x][curr.y]] 
		}
		if step.x == 1{
			return down[lines[curr.x][curr.y]] 
		}
		if step.y == 1{
			return right[lines[curr.x][curr.y]] 
		}
		return left[lines[curr.x][curr.y]] 
	}

	curr := cord{x - 1, y }
	step := cord{0, -1}

	loop := map[cord]struct{}{}
	inLoop := map[cord]struct{}{}
	notInLoop := map[cord]struct{}{}

	loop[cord{x, y}] = struct{}{}

	for lines[curr.x][curr.y] != 'S' {

		loop[curr] = struct{}{}
		curr.x += step.x
		curr.y += step.y
		step = nextStep(curr, step)

	}  

	inLoop := func(curr cord) bool{

		if _, ok := notInLoop[newCord] ; ok {
			return false
		}  
		if _, ok := inLoop[newCord] ; ok {
			return true
		}  

		visted := []cord{curr}
		q := []cord{curr}

		for len(q) != 0{
			currPos, q := q[0], q[1:]

			for i := -1; i < 2; i++ {
				newCord := cord{currPos.x, currPos.y + i}
				if _, ok := visted[newCord] ; !ok {
					if _, ok := inLoop[newCord] ; ok {

					}  

				}  
			}

			for i := -1; i < 2; i++ {
			}

		}
	}  


	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			curr := cord{i, j}
			if _, ok := notInLoop[curr]; ok {
				continue
			} 

			if _, ok := inLoop[curr]; ok {
				lines[i][j] ==

			} 
				
			

		}

	}


	
}

