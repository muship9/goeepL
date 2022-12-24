package main

import (
	"log"
	"os"
)

func main() {
	key := os.Getenv("DEEPL_API_KEY")
	log.Println(key)
}
