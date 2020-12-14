package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main()  {
	var validPasswords int
	var invalidPasswords int
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		counter := 0

		split := strings.Split(line, " ")

		positions := strings.Split(split[0], "-")
		letter := strings.Trim(string(split[1][0]), " ")
		password := strings.Trim(split[2], " ")

		pos1, _ := strconv.Atoi(positions[0])
		pos2, _ := strconv.Atoi(positions[1])

		for pos, char := range password {
			if string(char) == letter && (pos == pos1-1 || pos == pos2-1) {
				counter = counter + 1
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(validPasswords)
	fmt.Println(invalidPasswords)
}