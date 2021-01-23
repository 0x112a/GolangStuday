package main

import (
	_ "GolangStuday/sample/matchers"
	"GolangStuday/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
