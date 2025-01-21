package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
  db_manager()
  control_flow()
}
