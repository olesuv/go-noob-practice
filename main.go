package main

import "fmt"

func main() {
	words := []string{"go", "java", "go", "python", "go"}

	result := countWords(words)

	fmt.Println(result["go"])     // 3
	fmt.Println(result["java"])   // 1
	fmt.Println(result["python"]) // 1
	fmt.Println(result["rust"])   // 0
}

func countWords(words []string) map[string]int {
	wordsCounts := make(map[string]int)

	for _, word := range words {
		count, exists := wordsCounts[word]
		if exists {
			wordsCounts[word] = count + 1
		} else {
			wordsCounts[word] = 1
		}
	}

	return wordsCounts
}
