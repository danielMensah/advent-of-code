package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var year = 2020

func main() {
	result := twoEntries()

	fmt.Println(result)
}

func threeEntries() []int {
	data := readFile()

	var res []int

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for y := j + 1; y < len(data); y++ {
				if data[i] + data[j] + data[y] == 2020 {
					res = append(res, data[i])
					res = append(res, data[j])
					res = append(res, data[y])
					res = append(res, data[i] * data[j] * data[y])
				}
			}
		}
	}

	return res
}

func twoEntries() []int {
	data := readFile()

	var res []int

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] + data[j] == 2020 {
				res = append(res, data[i])
				res = append(res, data[j])
				res = append(res, data[i] * data[j])
			}
		}
	}

	return res
}

func readFile() []int  {
	file, err := os.Open("./input.txt")
	var arr []int

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}