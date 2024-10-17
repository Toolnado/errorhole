package main

import (
	"errors"

	"github.com/Toolnado/errorhole"
)

func main() {
	err := errors.New("error")
	defer errorhole.Catch() // Waiting for an error
	// do something...
	errorhole.Nil(err) // Check error
}
