package containsDuplicate

func containsDuplicate(nums []int) bool {
	intMap := map[int]bool{}
	for _, num := range nums {
		_, ok := intMap[num]
		if ok {
			return true
		} else {
			intMap[num] = true
		}
	}
	return false
}
