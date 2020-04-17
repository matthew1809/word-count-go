package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// get any user added arguments - user gives us the words to look for
	args, err := validateArgs(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	file, err := openFile("./words.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	timesFound, err := readLines(scanner, args)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(timesFound)
}

func validateArgs(args []string) ([]string, error) {

	if len(args) > 1 {
		return args[1:], nil
	}

	return nil, errors.New("not enough arguements")

}

func openFile(name string) (*os.File, error) {
	// open the file and check for error
	file, err := os.Open(name)

	if err != nil {
		return nil, errors.New("cannot open file")
	}

	return file, nil
}

func readLines(scanner *bufio.Scanner, args []string) (map[string]int, error) {
	timesFound := make(map[string]int, len(args))

	count := 0

	for scanner.Scan() {
		count++

		for _, searchWord := range args {
			words := strings.Fields(scanner.Text())
			// if strings.Contains(scanner.Text(), argsWithoutProg[i]) {
			for _, word := range words {
				if word == searchWord {
					//match!

					if _, found := timesFound[searchWord]; !found {
						timesFound[searchWord] = 0
					}

					timesFound[searchWord]++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("cannot read full file")
	}

	return timesFound, nil
}
