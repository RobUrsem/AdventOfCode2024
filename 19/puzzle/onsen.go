package puzzle

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Onsen struct {
	available        map[string]int
	longestAvailable int
	singles          map[string]int
	rx               *regexp.Regexp
	requested        []string
	possible         map[string]int
	availByLength    map[int][]string
}

func (o *Onsen) readAvailable(line string) {
	parts := strings.Split(line, ", ")

	o.available = map[string]int{}
	for _, part := range parts {
		o.available[part]++
	}
}

func (o *Onsen) reduceAvailable() {
	o.singles = map[string]int{"w": 0, "u": 0, "b": 0, "r": 0, "g": 0}
	for key := range o.singles {
		if _, exists := o.available[string(key)]; exists {
			o.singles[key]++
		}
	}
	fmt.Printf("Single patterns: %v\n", o.singles)

	notSingle := []string{}
	o.longestAvailable = 0
	for key := range o.available {
		allSingle := true
		for i := 0; i < len(key) && allSingle; i++ {
			if o.singles[string(key[i])] == 0 {
				allSingle = false
			}
		}

		if !allSingle {
			notSingle = append(notSingle, key)
		}

		o.longestAvailable = max(o.longestAvailable, len(key))
	}

	// Add the singles
	for key, count := range o.singles {
		if count > 0 {
			notSingle = append(notSingle, key)
		}
	}

	sort.Slice(notSingle, func(i, j int) bool {
		lenA := len(notSingle[i])
		lenB := len(notSingle[j])
		if lenA == lenB {
			return notSingle[i] > notSingle[j]
		}
		return lenA > lenB
	})

	o.availByLength = map[int][]string{}
	for i := range notSingle {
		length := len(notSingle[i])
		o.availByLength[length] = append(o.availByLength[length], notSingle[i])
	}

	expr := strings.Join(o.availByLength[2], "|")
	expr += "|" + strings.Join(o.availByLength[1], "|")
	for i := 3; i <= 8; i++ {
		keep := []string{}
		rx := regexp.MustCompile(expr)
		for _, pattern := range o.availByLength[i] {
			pat := rx.ReplaceAllString(pattern, "")
			if len(pat) > 0 {
				keep = append(keep, pattern)
				// } else {
				// 	fmt.Printf("Removing %v\n", pattern)
			}
		}
		o.availByLength[i] = keep
		if len(keep) > 0 {
			expr = strings.Join(o.availByLength[i], "|") + "|" + expr
		}
	}

	o.available = map[string]int{}
	for _, pattern := range strings.Split(expr, "|") {
		o.available[pattern]++
	}

	fmt.Println(expr)
	o.rx = regexp.MustCompile(expr)
}

func ReadPatterns(lines []string) Onsen {
	onsen := Onsen{}

	onsen.readAvailable(lines[0])
	onsen.reduceAvailable()

	onsen.requested = make([]string, len(lines)-2)
	for i := 2; i < len(lines); i++ {
		onsen.requested[i-2] = lines[i]
	}

	onsen.possible = make(map[string]int)
	return onsen
}

func (o *Onsen) RemovePattern(s string, index int) bool {
	if index == len(s) {
		o.possible[s]++
		// fmt.Printf(" is possible\n")
		return true
	} else if _, exists := o.possible[s]; exists {
		return true
	}

	found := false
	for i := min(index+8, len(s)); i > index && !found; i-- {
		prefix := s[index:i]
		if _, prefixExists := o.available[prefix]; prefixExists {
			if len(prefix) > 0 {
				// fmt.Printf("%v ", prefix)
				found = o.RemovePattern(s, i)
				// if !found {
				// fmt.Printf("\n      %v", s[:index])
				// }
			}
		}
	}
	return found
}

func (o *Onsen) FindPossible() int {
	fmt.Println("List of impossible patterns:")
	for _, req := range o.requested {
		if found := o.RemovePattern(req, 0); !found {
			fmt.Println(req)
		}
	}
	return len(o.possible)
}
