package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	flagsplit := strings.Split(os.Args[1], "=")
	flag := flagsplit[1]
	//fmt.Println(flag)
	if len(os.Args) > 4 {
		fmt.Println("error")
	}

	userText := os.Args[2]

	bannerFile := "standard.txt"
	bannerFile = os.Args[3] + ".txt"

	bannermap, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Println("error", err)
		return
	}
		finalOutput := GenerateArt(userText, bannermap)
		err = os.WriteFile(flag, []byte(finalOutput), 0644)
		if err != nil {
			fmt.Println("error reading file", err)
			return
		} 
		
					//fmt.Println(finalOutput)

	
}