package mean

func sum(array []int) int {
	var result int
	for _, v := range array {
		result += v
	}
	return result
}

func GetMean(listNum []int) int {
	mean := sum(listNum) / len(listNum)
	return mean
}
