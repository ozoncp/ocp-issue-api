package utils

import "testing"

func assertEqual(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func assertEqualSlice(t *testing.T, expected []int, actual []int) {
	if len(expected) != len(actual) {
		t.Errorf("expected size: %d, actual size: %d", len(expected), len(actual))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("expected: %d, actual: %d at index: %d", expected[i], actual[i], i)
		}
	}
}

func assertEqualMap(t *testing.T, expected map[string]int, actual map[string]int) {
	if len(expected) != len(actual) {
		t.Errorf("expected size: %d, actual size: %d", len(expected), len(actual))
	}

	for key, value := range expected {
		if value != actual[key] {
			t.Errorf("expected value: %d, actual value: %d for key: %s", value, actual[key], key)
		}
	}
}

func TestSplitToChunks(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}

	chunksOf2 := SplitToChunks(slice, 2)
	assertEqual(t, 3, len(chunksOf2))
	assertEqualSlice(t, []int{1, 2}, chunksOf2[0])
	assertEqualSlice(t, []int{3, 4}, chunksOf2[1])
	assertEqualSlice(t, []int{5, 6}, chunksOf2[2])

	chunksOf3 := SplitToChunks(slice, 3)
	assertEqual(t, 2, len(chunksOf3))
	assertEqualSlice(t, []int{1, 2, 3}, chunksOf3[0])
	assertEqualSlice(t, []int{4, 5, 6}, chunksOf3[1])

	chunksOf4 := SplitToChunks(slice, 4)
	assertEqual(t, 2, len(chunksOf3))
	assertEqualSlice(t, []int{1, 2, 3, 4}, chunksOf4[0])
	assertEqualSlice(t, []int{5, 6}, chunksOf4[1])

	chunksOfSize := SplitToChunks(slice, len(slice))
	assertEqual(t, 1, len(chunksOfSize))
	assertEqualSlice(t, slice, chunksOfSize[0])

	chunksOfMoreThanSize := SplitToChunks(slice, len(slice) + 1)
	assertEqual(t, 1, len(chunksOfMoreThanSize))
	assertEqualSlice(t, slice, chunksOfMoreThanSize[0])

	assertEqual(t, 0, len(SplitToChunks([]int{}, 1)))
}

func TestSwapMapPairs(t *testing.T) {
	source := map[int]string {
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	assertEqualMap(t, map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
	}, SwapMapPairs(source))

	assertEqualMap(t, map[string]int{}, SwapMapPairs(map[int]string{}))
}

func TestDeleteValues(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	assertEqualSlice(t, []int{2}, DeleteValues(slice, []int{1, 3, 4, 5, 6}))
	assertEqualSlice(t, []int{2, 6}, DeleteValues(slice, []int{1, 3, 4, 5}))
	assertEqualSlice(t, []int{1, 2, 3}, DeleteValues(slice, []int{5, 6, 6, 7, 4}))
	assertEqualSlice(t, []int{}, DeleteValues(slice, slice))
	assertEqualSlice(t, slice, DeleteValues(slice, []int{}))
}
