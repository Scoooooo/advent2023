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
	var histories [][]int

	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		history := []int{}
		for i := 0; i < len(vals); i ++{
			num, _ := strconv.Atoi(vals[i])
			history = append(history, num)
		}
		histories = append(histories, history)
	}




	getNext := func(index int) int{

		val := 0
		zero := true 
		for i := 0; i < len(histories[index]); i ++{
			zero = true
			for j := 0; j < len(histories[index]) - i - 1; j ++{
				histories[index][j] = histories[index][j + 1] - histories[index][j]
				if histories[index][j] != 0{
					zero = false
				}
			}
			val +=  histories[index][len(histories[index]) - i - 1] 
			if zero{
				return val
			}
		}
		return val
	}

	tot := 0
	for i := 0; i < len(histories); i ++{
		tot += getNext(i)
	}

	fmt.Println(tot)
	
}
