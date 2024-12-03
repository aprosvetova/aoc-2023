package main

import (
	"math"
	"strconv"
	"strings"
	"sync"
)

type mapping struct {
	srcMin int64
	srcMax int64
	delta  int64
}

func d0Convert(maps []mapping, val int64) int64 {
	for _, m := range maps {
		if val >= m.srcMin && val <= m.srcMax {
			return val + m.delta
		}
	}
	return val
}

func (*methods) D5P1(input string) string {
	lines := strings.Split(input, "\n")

	var allMappings [][]mapping

	var mappings []mapping
	for _, l := range lines[2:] {
		if strings.Contains(l, "map") {
			continue
		}
		if l == "" {
			allMappings = append(allMappings, mappings)
			mappings = nil
			continue
		}
		parts := strings.Split(l, " ")
		dest, _ := strconv.ParseInt(parts[0], 10, 64)
		src, _ := strconv.ParseInt(parts[1], 10, 64)
		length, _ := strconv.ParseInt(parts[2], 10, 64)
		mappings = append(mappings, mapping{
			delta:  dest - src,
			srcMin: src,
			srcMax: src + length - 1,
		})
	}
	allMappings = append(allMappings, mappings)

	var min int64 = math.MaxInt64
	for _, seed := range strings.Split(lines[0], " ")[1:] {
		val, _ := strconv.ParseInt(seed, 10, 64)
		for _, maps := range allMappings {
			val = d0Convert(maps, val)
		}
		if val < min {
			min = val
		}
	}

	return strconv.FormatInt(min, 10)
}

func (*methods) D5P2(input string) string {
	lines := strings.Split(input, "\n")

	var allMappings [][]mapping

	var mappings []mapping
	for _, l := range lines[2:] {
		if strings.Contains(l, "map") {
			continue
		}
		if l == "" {
			allMappings = append(allMappings, mappings)
			mappings = nil
			continue
		}
		parts := strings.Split(l, " ")
		dest, _ := strconv.ParseInt(parts[0], 10, 64)
		src, _ := strconv.ParseInt(parts[1], 10, 64)
		length, _ := strconv.ParseInt(parts[2], 10, 64)
		mappings = append(mappings, mapping{
			delta:  dest - src,
			srcMin: src,
			srcMax: src + length - 1,
		})
	}
	allMappings = append(allMappings, mappings)

	var results []int64
	wg := sync.WaitGroup{}

	values := strings.Split(lines[0], " ")[1:]
	for i := 0; i < len(values); i += 2 {
		min, _ := strconv.ParseInt(values[i], 10, 64)
		length, _ := strconv.ParseInt(values[i+1], 10, 64)
		max := min + length - 1

		wg.Add(1)
		go func(min, max int64) {
			var totalMin int64 = math.MaxInt64
			for val := min; val <= max; val++ {
				res := val
				for _, maps := range allMappings {
					res = d0Convert(maps, res)
				}
				if res < totalMin {
					totalMin = res
				}
			}
			results = append(results, totalMin)
			wg.Done()
		}(min, max)
	}

	wg.Wait()

	var totalMin int64 = math.MaxInt64
	for _, r := range results {
		if r < totalMin {
			totalMin = r
		}
	}

	return strconv.FormatInt(totalMin, 10)
}
