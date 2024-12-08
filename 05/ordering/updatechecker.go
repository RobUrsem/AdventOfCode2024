package ordering

func contains(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func updateIsValid(update Update, rulebook Rulebook) bool {
	for _, rule := range rulebook {
		indices := []int{
			contains(update, rule[0]),
			contains(update, rule[1]),
		}
		if indices[0] >= 0 && indices[1] >= 0 && indices[0] > indices[1] {
			// fmt.Printf("Violation of rule %v: %v", rule, indices)
			return false
		}
	}
	return true
}

func FilterUpdates(updates []Update, rulebook Rulebook) []Update {
	var validUpdates []Update

	for _, update := range updates {
		if updateIsValid(update, rulebook) {
			validUpdates = append(validUpdates, update)
		}
	}

	return validUpdates
}
