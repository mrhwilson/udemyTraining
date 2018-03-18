package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}

	for i := 0; i < 200; i++ {
		// fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
		fmt.Fprintf(file, "%d \t %b \t %x \t %q \n", i, i, i, i)
	}
	defer file.Close()
}
