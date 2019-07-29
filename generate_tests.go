package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("numbers") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for i := 0; i < 10000; i++ { // Generating...
		_, err = file.WriteString(fmt.Sprintf("%d\n", i)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}
