package main

import (
	"bufio"
	"fmt"
	"log"
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

	dist:= []int{}
	time:= []int{}

	distString := strings.Split(strings.Split(lines[1], ":")[1], " ") 
	timeString := strings.Split(strings.Split(lines[0], ":")[1], " ") 

	for  i:= 0; i<len(distString); i++ {
		if distString[i] == "" {continue}
		num, _ := strconv.Atoi(distString[i])
		dist = append(dist, num)
	}

	for  i:= 0; i<len(timeString); i++ {
		if timeString[i] == "" {continue}
		num, _ := strconv.Atoi(timeString[i])
		time = append(time, num)
	}


	count := 1

	for i := 0; i < len(time); i++ {
		currVal := 0
		for j := 1; j < time[i]; j++ {
			tot := j * (time[i] - j) 
			if tot > dist[i] {
				currVal += 1
			}  
		}
		count *= currVal

	}
	fmt.Println(count)


}
