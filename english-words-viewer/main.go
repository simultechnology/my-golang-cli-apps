package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	fmt.Println("start")

	file, err := os.Open("data/alc_2022-06-05.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wordList [][]string

	r := regexp.MustCompile(`\s{2,}`)

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		word := scanner.Text()
		splits := r.Split(word, -1)
		if len(splits) > 1 {
			// fmt.Printf("word: %s, date: %s\n", splits[0], splits[1])
			// strings.Split(word, "")
			wordList = append(wordList, splits)
			count += 1
		}
	}
	fmt.Println(count)
	for i, words := range wordList {
		if i < 10 {
			fmt.Printf("%d, word: %s, date: %s\n", i+1, words[0], words[1])
		}
	}
	fmt.Println(len(wordList))

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
