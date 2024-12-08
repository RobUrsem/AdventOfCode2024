package ordering

import "testing"

func createTestInput() []string {
	return []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}
}

func TestConstructRulebook(t *testing.T) {
	input := createTestInput()
	rulebook, err := ConstructRulebook(input)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	numRules := len(rulebook)
	expected := 21
	if numRules != expected {
		t.Errorf("Expected %d rules but got %d", expected, numRules)
	}
}

func TestGetUpdates(t *testing.T) {
	input := createTestInput()
	updates, err := GetUpdates(input)

	if err != nil {
		t.Errorf("got error: %v", err)
	}

	numRules := len(updates)
	expected := 6
	if numRules != expected {
		t.Errorf("Expected %d rules but got %d", expected, numRules)
	}

	expectedUpdate := []int{75, 47, 61, 53, 29}
	for i, num := range updates[0] {
		if num != expectedUpdate[i] {
			t.Errorf("Expected %v at index %v but got %v", expectedUpdate[i], i, num)
		}
	}
}
