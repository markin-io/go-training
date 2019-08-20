package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func createReader() *bufio.Reader {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: 'cat file | main' or 'main < file'")
		return nil
	}

	return bufio.NewReader(os.Stdin)
}

func ReadPipeInputString() []string {
	reader := createReader()

	var output []string

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		byteToInt := string(input)
		output = append(output, byteToInt)
	}
	return output
}

func ReadPipeInputInt() []int {
	reader := createReader()

	var output []int

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		byteToInt, _ := strconv.Atoi(string(input))
		output = append(output, byteToInt)
	}
	return output
}
