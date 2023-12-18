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


	distString := strings.Split(strings.Split(lines[1], ":")[1], " ") 
	timeString := strings.Split(strings.Split(lines[0], ":")[1], " ") 
	distS := ""
	timeS := ""

	for  i:= 0; i<len(distString); i++ {
		if distString[i] == "" {continue}
		distS += distString[i]
	}

	for  i:= 0; i<len(timeString); i++ {
		if timeString[i] == "" {continue}
		timeS += timeString[i]
	}

	dist, _:= strconv.Atoi(distS) 
	time, _ := strconv.Atoi(timeS) 


	search := func(l int, r int, dir int) int{

		for l <= r{
			mid := (l + r) / 2
			tot := mid * (time - mid)
			totDir := (mid + dir)  * (time - (mid + dir))

			if tot > dist{
				if totDir <= dist{
					return mid
				}
				if dir == - 1{
					r = mid - 1
				} else {
					l = mid + 1
				}

			} else {
				if dir == - 1{
					if tot > totDir{
						l = mid + 1
					} else {
						r = mid - 1
					}
				} else {
					if tot < totDir{
						l = mid + 1
					} else {
						r = mid - 1
					}
				}

			} 
		}

		return -1
	}



	l := 0
	r := time

	s := search(l, r, -1)
	e := search(l, r, 1)

	fmt.Println(e - s + 1)


}
