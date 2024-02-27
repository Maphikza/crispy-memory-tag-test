package main

import (
	"strings"
	"testing"
)

func BenchmarkFindUserTags(b *testing.B) {
	// Setup: Initialize any required variables or state.
	users := map[string]string{
		"john":     "@johnny",
		"testuser": "@testtag",
		// Add more users as needed to simulate a realistic scenario
	}
	input := "This is a test string with john, and TestUser mentioned."
	inputLower := strings.ToLower(input)

	// Reset the timer to exclude setup time.
	b.ResetTimer()

	// Run the function b.N times
	for i := 0; i < b.N; i++ {
		findUserTags(inputLower, users)
	}
}
