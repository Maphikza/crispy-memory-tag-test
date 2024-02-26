package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
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

	input := "This is a test string with john🔥 and TestUser mentioned."
	inputLower := strings.ToLower(input) // Convert input to lowercase once

	foundTags := findUserTags(inputLower, users)
	fmt.Println("Tags found in the string:")
	for _, tag := range foundTags {
		fmt.Println(tag)
	}

	username, tag := findUsernameInInput(inputLower, users)
	if username != "" {
		post := fmt.Sprintf("Just posted a tweet! Follow me %s %s", tag, username)
		fmt.Println("Twitter-like Post:")
		fmt.Println(post)
	} else {
		fmt.Println("No username found in the input string.")
	}

	fmt.Printf("Time taken for the entire process: %s\n", time.Since(startTime))
}

func findUserTags(input string, users map[string]string) []string {
	var tags []string
	for username, tag := range users {
		if strings.Contains(input, strings.ToLower(username)) {
			tags = append(tags, tag)
		}
	}
	return tags
}

func findUsernameInInput(input string, users map[string]string) (string, string) {
	for username, tag := range users {
		if strings.Contains(input, strings.ToLower(username)) {
			return username, tag
		}
	}
	return "", ""
}
