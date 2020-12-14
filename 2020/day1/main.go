package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var res []int
	var data []int
	year := 2020

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		data = append(data, num)

		if contains(year - num, data) {
			res = append(res, year -num)
			res = append(res, num)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if res != nil && len(res) == 2 {
		fmt.Println(res)
		fmt.Println(res[0] * res[1])
	} else {
		fmt.Println("Nothing found")
	}

}

func contains(item int, arr []int) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}
