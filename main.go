package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	timesFound := make(map[string]int, len(argsWithoutProg))

	for i := 0; i < len(argsWithoutProg); i++ {
		timesFound[argsWithoutProg[i]] = 0
	}

	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		count++

		for i := 0; i < len(argsWithoutProg); i++ {
			if strings.Contains(scanner.Text(), argsWithoutProg[i]) {
				timesFound[argsWithoutProg[i]]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(timesFound)

}
