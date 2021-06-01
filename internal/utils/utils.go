package utils

import "github.com/ozoncp/ocp-issue-api/internal/issues"

func SplitIntegersToChunks(slice []int, chunkSize int) (chunks [][]int) {
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

func SplitIssuesToChunks(issues []issues.Issue, chunkSize int) (chunks [][]issues.Issue) {
	if chunkSize <= 0 {
		panic("Only positive value of chunkSize is possible!")
	}

	for begin := 0; begin < len(issues); {
		end := min(begin + chunkSize, len(issues))
		chunks = append(chunks, issues[begin:end])
		begin = end
	}

	return chunks
}

func SwapMapPairs(source map[int]string) map[string]int {
	swapped := make(map[string]int, len(source))

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

func SliceToMap(slice []issues.Issue) map[uint64]issues.Issue {
	result := make(map[uint64]issues.Issue, len(slice))

	for _, issue := range slice {
		result[issue.Id] = issue
	}

	return result
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func contains(slice []int, value int) bool {
	for _, _value := range slice {
		if _value == value {
			return true
		}
	}

	return false
}
