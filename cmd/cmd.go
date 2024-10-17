package main

import (
	"errors"

	"github.com/Toolnado/errorhole"
)

func main() {
	defer errorhole.Catch() // Waiting for an error
	// do something...
	errorhole.Nil(errors.New("error")) // Check error
}
