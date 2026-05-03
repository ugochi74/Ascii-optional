package main

import (
	"errors"
	"fmt"
)

func Validate(text string) (rune, error) {
	if len(text) == 0 {
		return 0, nil
	}
	for _, ch := range text {

		if ch < 32 || ch > 126 {
			err := fmt.Sprintf("unsupported character: %c", ch)
			return rune(ch), errors.New(err)

		}

	}
	return 0, nil

}
