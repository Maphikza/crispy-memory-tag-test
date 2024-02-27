package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	startTime := time.Now()

	data, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var users map[string]string
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Preprocess users to lowercase keys
	lowercaseUsers := make(map[string]string)
	for username, tag := range users {
		lowercaseUsers[strings.ToLower(username)] = tag
	}

	input := "This is a test string with john, and TestUser mentioned."
	inputLower := strings.ToLower(input)

	foundResults := findUserTags(inputLower, lowercaseUsers)
	fmt.Println("Exact tags and corresponding usernames found in the string:")
	for tag, username := range foundResults {
		fmt.Printf("Tag: %s, Username: %s\n", tag, username)
	}

	fmt.Printf("Time taken for the entire process: %s\n", time.Since(startTime))
}

func findUserTags(input string, users map[string]string) map[string]string {
	results := make(map[string]string)
	words := strings.Fields(input)

	for _, word := range words {
		// Combine trimming and lowering case to minimize operations
		trimmedWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		}))

		if username, exists := users[trimmedWord]; exists {
			results[username] = trimmedWord
		}
	}

	return results
}
