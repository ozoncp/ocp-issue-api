package tests

import (
	"github.com/ozoncp/ocp-issue-api/internal/issues"
	"github.com/ozoncp/ocp-issue-api/internal/utils"

	"testing"
	"time"
)

func TestSplitToChunks(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}

	chunksOf2 := utils.SplitToChunks(slice, 2)
	assertEqual(t, 3, len(chunksOf2))
	assertEqualSlice(t, []int{1, 2}, chunksOf2[0])
	assertEqualSlice(t, []int{3, 4}, chunksOf2[1])
	assertEqualSlice(t, []int{5, 6}, chunksOf2[2])

	chunksOf3 := utils.SplitToChunks(slice, 3)
	assertEqual(t, 2, len(chunksOf3))
	assertEqualSlice(t, []int{1, 2, 3}, chunksOf3[0])
	assertEqualSlice(t, []int{4, 5, 6}, chunksOf3[1])

	chunksOf4 := utils.SplitToChunks(slice, 4)
	assertEqual(t, 2, len(chunksOf3))
	assertEqualSlice(t, []int{1, 2, 3, 4}, chunksOf4[0])
	assertEqualSlice(t, []int{5, 6}, chunksOf4[1])

	chunksOfSize := utils.SplitToChunks(slice, len(slice))
	assertEqual(t, 1, len(chunksOfSize))
	assertEqualSlice(t, slice, chunksOfSize[0])

	chunksOfMoreThanSize := utils.SplitToChunks(slice, len(slice)+1)
	assertEqual(t, 1, len(chunksOfMoreThanSize))
	assertEqualSlice(t, slice, chunksOfMoreThanSize[0])

	assertEqual(t, 0, len(utils.SplitToChunks([]int{}, 1)))
}

func TestSplitIssuesToChunks(t *testing.T) {
	slice := []issues.Issue{
		{1, 1, 1, 42, time.Now()},
		{2, 1, 2, 42, time.Now()},
		{3, 2, 3, 0, time.Now()},
		{4, 1, 3, 23, time.Now()},
	}

	chunksOf2 := utils.SplitIssuesToChunks(slice, 2)
	assertEqual(t, 2, len(chunksOf2))
	assertEqualIssueSlice(t, []issues.Issue{slice[0], slice[1]}, chunksOf2[0])
	assertEqualIssueSlice(t, []issues.Issue{slice[2], slice[3]}, chunksOf2[1])

	chunksOf3 := utils.SplitIssuesToChunks(slice, 3)
	assertEqual(t, 2, len(chunksOf2))
	assertEqualIssueSlice(t, []issues.Issue{slice[0], slice[1], slice[2]}, chunksOf3[0])
	assertEqualIssueSlice(t, []issues.Issue{slice[3]}, chunksOf3[1])

	chunksOfSize := utils.SplitIssuesToChunks(slice, len(slice))
	assertEqual(t, 1, len(chunksOfSize))
	assertEqualIssueSlice(t, slice, chunksOfSize[0])

	chunksOfMoreThanSize := utils.SplitIssuesToChunks(slice, len(slice)+1)
	assertEqual(t, 1, len(chunksOfMoreThanSize))
	assertEqualIssueSlice(t, slice, chunksOfMoreThanSize[0])

	assertEqual(t, 0, len(utils.SplitIssuesToChunks([]issues.Issue{}, 1)))
}

func TestSwapMapPairs(t *testing.T) {
	source := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	assertEqualMap(t, map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}, utils.SwapMapPairs(source))

	assertEqualMap(t, map[string]int{}, utils.SwapMapPairs(map[int]string{}))
}

func TestDeleteValues(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	assertEqualSlice(t, []int{2}, utils.DeleteValues(slice, []int{1, 3, 4, 5, 6}))
	assertEqualSlice(t, []int{2, 6}, utils.DeleteValues(slice, []int{1, 3, 4, 5}))
	assertEqualSlice(t, []int{1, 2, 3}, utils.DeleteValues(slice, []int{5, 6, 6, 7, 4}))
	assertEqualSlice(t, []int{}, utils.DeleteValues(slice, slice))
	assertEqualSlice(t, slice, utils.DeleteValues(slice, []int{}))
}

func TestSliceToMap(t *testing.T) {
	issuesSlice := []issues.Issue{
		{1, 1, 1, 42, time.Now()},
		{2, 1, 2, 42, time.Now()},
		{3, 2, 3, 0, time.Now()},
		{4, 1, 3, 23, time.Now()},
	}

	issuesMap := utils.SliceToMap(issuesSlice)

	assertEqual(t, len(issuesSlice), len(issuesMap))

	for i := 0; i < len(issuesSlice); i++ {
		issue := issuesSlice[i]
		assertEqualIssue(t, issue, issuesMap[issue.Id])
	}
}
