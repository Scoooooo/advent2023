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
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	limit := map[string]int{"green": 13, "red": 12, "blue": 14}

	curr := map[string]int{"green": 0, "red": 0, "blue": 0}

	currVal := 0
	for scanner.Scan() {
		line := scanner.Text()
		split1 := strings.Split(line, ":")

		sets := strings.Split(split1[1], ";")

		pos := true
		for _, set := range sets {
			curr = map[string]int{"green": 0, "red": 0, "blue": 0}
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				vals := strings.Split(cube, " ")
				num, _ := strconv.Atoi(vals[1])
				curr[vals[2]] += num
			}
			for key, val := range curr {
				limitVal, exisits := limit[key]
				if exisits {
					if val > limitVal {
						pos = false
						break
					}
				}
			}
		}
		if pos {
			gameNumber := strings.Split(split1[0], " ")
			val, _ := strconv.Atoi(gameNumber[1])
			currVal += val
		}
	}
	fmt.Println(currVal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
