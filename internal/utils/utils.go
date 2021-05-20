package utils

func SplitToChunks(slice []int, chunkSize int) (chunks [][]int) {
	if chunkSize <= 0 {
		panic("Only positive value of chunkSize is possible!")
	}

	for begin := 0; begin < len(slice); {
		end := min(begin + chunkSize, len(slice))
		chunks = append(chunks, slice[begin:end])
		begin = end
	}

	return chunks
}

func SwapMapPairs(source map[int]string) map[string]int {
	swapped := map[string]int{}

	for key, value := range source {
		swapped[value] = key
	}

	return swapped
}

func DeleteValues(source []int, values []int) (filtered []int) {
	for _, value := range source {
		if !contains(values, value) {
			filtered = append(filtered, value)
		}
	}

	return filtered
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
