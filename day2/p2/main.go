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

	max := map[string]int{"green": 0, "red": 0, "blue": 0}

	curr := map[string]int{"green": 0, "red": 0, "blue": 0}
	currVal := 0

	for scanner.Scan() {
		line := scanner.Text()
		split1 := strings.Split(line, ":")

		sets := strings.Split(split1[1], ";")

		max = map[string]int{"green": 0, "red": 0, "blue": 0}

		for _, set := range sets {
			curr = map[string]int{"green": 0, "red": 0, "blue": 0}
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				vals := strings.Split(cube, " ")
				num, _ := strconv.Atoi(vals[1])
				curr[vals[2]] += num
			}
			for key, val := range curr {
				if val > max[key] {
					max[key] = val
				}
			}
		}
		currVal += max["green"] * max["red"] * max["blue"]
	}
	fmt.Println(currVal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
