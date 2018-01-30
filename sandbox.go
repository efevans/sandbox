package main

import (
	"fmt"

	"github.com/efevans/trivia"
)

func main() {
	q := trivia.GetQuestion()
	fmt.Println(q.ID)
	fmt.Println(q.Question)
	fmt.Println(q.Answer)
}
