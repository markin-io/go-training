package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var genType = ""

	if len(os.Args) > 1 {
		genType = string(os.Args[1])
	}

	rand.Seed(time.Now().UTC().UnixNano())
	file, err := os.Create("numbers") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	amount := 10000

	for i := 0; i < amount; i++ { // Generating...
		number := i
		if genType == "random" {
			number = rand.Intn(amount)
		}
		_, err = file.WriteString(fmt.Sprintf("%d\n", number)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}
