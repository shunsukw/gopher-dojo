package main

import (
	"time"

	"github.com/shunsukw/gopher-dojo/dojo1/kadai3/questioner"
)

func main() {
	words := []string{
		"enhance",
		"chocolate",
		"interruption",
		"permission",
		"volcanoconiosis",
		"pterygon",
		"masterpiece",
		"accumulate",
		"crossover",
		"sculpture",
		"elephant",
		"algorithms",
	}

	quiz := questioner.New(30*time.Second, words)
	quiz.Start()
}
