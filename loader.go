package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type methods struct {
}

type puzzle struct {
	Day     int
	Part    int
	Execute func(input string) string
}

var re = regexp.MustCompile(`^D(\d{1,2})P(\d)$`)

func extractNumbers(input string) (int, int, error) {
	matches := re.FindStringSubmatch(input)
	if matches == nil {
		return 0, 0, fmt.Errorf("invalid format")
	}

	d, err1 := strconv.Atoi(matches[1])
	p, err2 := strconv.Atoi(matches[2])

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("error parsing numbers")
	}

	return d, p, nil
}

func loadPuzzles() ([]puzzle, int) {
	var puzzles []puzzle
	pst := &methods{}
	pt := reflect.TypeOf(pst)

	var last int

	for i := 0; i < pt.NumMethod(); i++ {
		method := pt.Method(i)
		d, p, err := extractNumbers(method.Name)
		if err != nil {
			continue
		}
		puzzles = append(puzzles, puzzle{
			Day:  d,
			Part: p,
			Execute: func(input string) string {
				args := []reflect.Value{
					reflect.ValueOf(pst),
					reflect.ValueOf(input),
				}
				res := method.Func.Call(args)[0]
				return res.String()
			},
		})
		if d > last {
			last = d
		}
	}

	return puzzles, last
}

func findPuzzles(puzzles []puzzle, day int) (p1, p2 *puzzle) {
	for _, p := range puzzles {
		if p.Day == day {
			if p.Part == 1 {
				p1 = &p
			} else if p.Part == 2 {
				p2 = &p
			}
		}
	}
	return
}

func findLastPuzzle(puzzles []puzzle, day int) *puzzle {
	p1, p2 := findPuzzles(puzzles, day)
	if p2 != nil {
		return p2
	}
	return p1
}

func askForDay(puzzles []puzzle, last int) (p1, p2 *puzzle) {
	var day int
	var input string
	fmt.Printf("Enter the day [%d]: ", last)
	_, err := fmt.Scanln(&input)
	if err != nil {
		day = last
	} else {
		day, _ = strconv.Atoi(input)
	}
	return findPuzzles(puzzles, day)
}

func ask(puzzles []puzzle, last int) *puzzle {
	var p1, p2 *puzzle
	for {
		p1, p2 = askForDay(puzzles, last)
		if p1 == nil {
			fmt.Println("Day not found")
		} else {
			break
		}
	}

	if p2 == nil {
		return p1
	}

	for {
		var input string
		fmt.Print("Enter the part [2]: ")
		_, err := fmt.Scanln(&input)
		if err != nil || input == "2" {
			return p2
		}
		if input == "1" {
			return p1
		}
	}
}
