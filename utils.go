package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadPipeInput() []int {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: cat file | main")
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
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
