package main

import (
	"log"

	"github.com/efevans/trivia"
)

func main() {
	q := trivia.GetQuestion()
	log.Println(q.ID)
	log.Println(q.Type)
}
