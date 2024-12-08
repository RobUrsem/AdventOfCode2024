package ordering

func contains(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func identifyViolatedRule(update Update, rulebook Rulebook) (int, []int) {
	for i, rule := range rulebook {
		indices := []int{
			contains(update, rule[0]),
			contains(update, rule[1]),
		}
		if indices[0] >= 0 && indices[1] >= 0 && indices[0] > indices[1] {
			// fmt.Printf("Violation of rule %v: %v", rule, indices)
			return i, indices
		}
	}
	return -1, nil
}

func updateIsValid(update Update, rulebook Rulebook) bool {
	ruleIdx, _ := identifyViolatedRule(update, rulebook)
	return ruleIdx == -1
}

func FilterUpdates(updates []Update, rulebook Rulebook) ([]Update, []Update) {
	var validUpdates, invalidUpdates []Update

	for _, update := range updates {
		if updateIsValid(update, rulebook) {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	return validUpdates, invalidUpdates
}
