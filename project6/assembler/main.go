package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read from file
	data, err := os.Open("sample.asm")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	defer data.Close()
	// print

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}
}
