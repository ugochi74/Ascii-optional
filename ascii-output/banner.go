package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {

	// READING FILE CONTENT
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty file")
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[1:]
	if len(lines) != 855 {
		return nil, errors.New("invalid file content")
	}
	result := make(map[rune][]string)

	for i := ' '; i <= '~'; i++ {
		start := int(i-32) * 9
		end := start + 8
		result[i] = lines[start:end]
	}
	return result, nil

	
}
