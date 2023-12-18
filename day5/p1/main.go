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

type mapping struct{
	num1 int
	num2 int 
	r int
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

	seeds := []int{}


	stringSeeds := strings.Split(strings.Split(lines[0], ":")[1], " ") 

	for _, val := range stringSeeds {
		if val == ""{ continue }

		num, _ := strconv.Atoi(val)
		seeds = append(seeds, num)
	}


	minVal := math.MaxInt 

	mappings := [][]mapping{}
	curMap := -1

	for i := 1; i < len(lines); i++{

		if lines[i] == ""{ 
			curMap += 1  
			mappings = append(mappings, []mapping{})
			continue
		} 
		nums := strings.Split(lines[i], " ") 
		if len(nums) < 3{
			continue
		}

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		num3, _ := strconv.Atoi(nums[2])

		mappings[curMap] = append(mappings[curMap], mapping{num1, num2, num3})

	} 


	nextVal := func(val int, mapIndex int) int {
		for _,  currMapping := range mappings[mapIndex]{
			if val >= currMapping.num2 && val <= (currMapping.num2 + currMapping.r){
				return currMapping.num1 + (val - currMapping.num2) 
			}

		}
		return val
	}  

	for _, seed := range seeds {
		val := seed
		for i := 0; i < len(mappings); i++{
			val = nextVal(val, i)		
		}
		minVal = min(val, minVal)

	}

	fmt.Println(minVal)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
