package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	puzzles, last := loadPuzzles()
	var p *puzzle
	if os.Getenv("DONT_ASK") == "1" {
		p = findLastPuzzle(puzzles, last)
	} else {
		p = ask(puzzles, last)
	}
	fmt.Printf("Launching Day %d Part %d\n---\n", p.Day, p.Part)

	input, err := os.ReadFile(fmt.Sprintf("data/d%d.txt", p.Day))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	t := time.Now()
	result := p.Execute(string(input))
	total := time.Since(t)
	fmt.Println(result)
	fmt.Println("---\nFinished in", total)
}
