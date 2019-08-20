package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const ORDERED_NUM = "orderednum"
const RAND_NUM = "randnum"
const STRING = "string"

func generateNumbers(genType string, amount int) {
	file, err := os.Create("numbers") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for i := 0; i < amount; i++ { // Generating...
		number := i
		if genType == RAND_NUM {
			number = rand.Intn(amount)
		}
		_, err = file.WriteString(fmt.Sprintf("%d\n", number)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateStrings(amount int) {
	file, err := os.Create("strings_test") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for i := 0; i < amount; i++ { // Generating...
		_, err = file.WriteString(fmt.Sprintf("%s\n", RandStringRunes(10))) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}

func main() {
	genType := ORDERED_NUM
	amount := 100

	if len(os.Args) > 1 {
		genType = string(os.Args[1])

		if len(os.Args) == 3 {
			number, _ := strconv.Atoi(string(os.Args[2]))
			amount = number
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())

	if genType == ORDERED_NUM || genType == RAND_NUM {
		generateNumbers(genType, amount)
	} else if genType == STRING {
		generateStrings(amount)
	}

}
