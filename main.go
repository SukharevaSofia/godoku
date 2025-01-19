package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
	log.Println("This is a test log entry")

	db_manager()
	field := generate()
	fmt.Println("")
	_ = getUserField(field, 27)
}
