package main

import "strings"

func RenderLine(text string, banner map[rune][]string) []string {
	var textslice []string

	for r := 0; r < 8; r++ {
		var write strings.Builder
		//word := ""
		for _, ch := range text {
			write.WriteString(banner[ch][r])
			//word += banner[ch][r]
		}
		textslice = append(textslice, write.String())
	}
	return textslice

}
