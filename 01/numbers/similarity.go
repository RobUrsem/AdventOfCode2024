package numbers

func setupCountMap(nums []int) map[int]int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	return countMap
}

func CalcSimilarity(a []int, b []int) int {
	if len(a) == 0 {
		return 0
	}

	similarity := 0
	countMap := setupCountMap(b)
	for _, num := range a {
		count := countMap[num]
		similarity += num * count
	}

	return similarity
}
