package main

import (
	"bufio"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	none := []string{}
	one := []string{"one"}
	multiple := []string{"one", "two", "three"}

	testNone, err := validateArgs(none)
	if testNone != nil || err == nil {
		t.Error("No arguements should fail")
	}

	testOne, err := validateArgs(one)
	if testOne != nil || err == nil {
		t.Error("One argument should fail")
	}

	testMultiple, err := validateArgs(multiple)
	if testMultiple == nil || err != nil {
		t.Error("Multiple arguements should succeed")
	}
}

func TestOpenFile(t *testing.T) {
	valid := "./words.txt"
	nonValid := "./garbage.md"

	validFile, err := openFile(valid)

	if validFile == nil || err != nil {
		t.Error("Valid file path should succeed")
	}

	invalidFile, err := openFile(nonValid)

	if invalidFile != nil || err == nil {
		t.Error("Invalid file path should fail")
	}
}

func TestReadLines(t *testing.T) {

	args := []string{"won", "two"}
	file, _ := openFile("./words.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result, err := readLines(scanner, args)

	if result == nil || err != nil {
		t.Error("Valid file and args should succeed")
	}
}
