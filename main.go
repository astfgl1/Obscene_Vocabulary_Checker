package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fileName string
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tabooWords := make(map[string]bool)
	for scanner.Scan() {
		tabooWords[strings.ToLower(scanner.Text())] = true
	}

	for {
		var sentence string
		fmt.Scan(&sentence)

		if sentence == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		words := strings.Fields(sentence)
		for _, word := range words {
			if _, ok := tabooWords[strings.ToLower(word)]; ok {
				sentence = strings.ReplaceAll(sentence, word, strings.Repeat("*", len(word)))
			}
		}
		fmt.Println(sentence)

	}
}
