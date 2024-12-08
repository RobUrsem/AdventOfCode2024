package ordering

import "fmt"

func swap(update Update, i, j int) {
	if i < 0 || j < 0 || i >= len(update) || j >= len(update) {
		fmt.Println("Invalid indices")
		return
	}
	update[i], update[j] = update[j], update[i]
}

func fixUpdate(update Update, rulebook Rulebook) Update {
	fixed := update

	for {
		ruleIdx, indices := identifyViolatedRule(update, rulebook)
		if ruleIdx >= 0 {
			swap(fixed, indices[0], indices[1])
		} else {
			break
		}
	}

	return fixed
}

func FixUpdates(updates []Update, rulebook Rulebook) []Update {
	var fixed []Update

	for _, update := range updates {
		fixed = append(fixed, fixUpdate(update, rulebook))
	}

	return fixed
}
