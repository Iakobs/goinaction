package main

import (
	"log"
	"os"

	_ "github.com/iakobs/code/chapter2/sample/matchers"
	"github.com/iakobs/code/chapter2/sample/search"
)

// init is called prior to main
func init() {
	// Change the device for logging stdout
	log.SetOutput(os.Stdout)
}

// Main is the entry point for the program
func main() {
	// Perform the search for the specified term
	search.Run("president")
}