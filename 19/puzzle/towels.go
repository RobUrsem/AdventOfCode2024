package puzzle

import (
	"strings"

	lru "github.com/hashicorp/golang-lru/v2"
)

type Towels struct {
	patterns map[string]int
	designs  []string
	cache    *lru.Cache[string, int]
}

func MakeTowels(input []string) Towels {
	towels := Towels{
		patterns: map[string]int{},
		designs:  input[2:],
	}

	for _, pattern := range strings.Split(input[0], ", ") {
		towels.patterns[pattern]++
	}

	towels.cache, _ = lru.New[string, int](128)

	return towels
}

func (t *Towels) is_possible(candidate string) int {
	if len(candidate) == 0 {
		// Nothing more to evaluate, return that
		// we have found a possible way.
		return 1
	}

	if value, ok := t.cache.Get(candidate); ok {
		// we have evaluated this pattern before
		// no need to redo it and just return the
		// previous evaluation
		return value
	}

	numWays := 0
	for i := range candidate {
		left := candidate[:i+1]
		if _, found := t.patterns[left]; !found {
			continue
		}

		// We know the left is a known pattern,
		// now test the right side
		right := candidate[i+1:]
		numWays += t.is_possible(right)
	}

	// This is a lot like the stones. We know
	// that the rest `candidate` can be created in
	// a certain number of ways. When we test other
	// patterns at the start of the string (i.e. backtracking)
	// we don't need to re-evaluate these patterns anymore
	t.cache.Add(candidate, numWays)
	return numWays
}

func (t *Towels) Part1() int {
	numPossible := 0

	for _, text := range t.designs {
		if t.is_possible(text) > 0 {
			numPossible++
		}
	}

	return numPossible
}

func (t *Towels) Part2() int {
	numPossible := 0

	for _, text := range t.designs {
		numPossible += t.is_possible(text)
	}

	return numPossible
}
