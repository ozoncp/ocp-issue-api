package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/ozoncp/ocp-issue-api/internal/models"
	. "github.com/ozoncp/ocp-issue-api/internal/utils"
	"time"
)

var _ = Describe("Utils", func() {
	Context("SplitToChunks", func() {
		var integers []int
		var issues []Issue
		var chunkSize int

		BeforeEach(func() {
			integers = []int{1, 2, 3, 4, 5, 6}

			issues = []Issue{
				{1, 1, 1, 42, time.Now()},
				{2, 1, 2, 42, time.Now()},
				{3, 2, 3, 0, time.Now()},
				{4, 1, 3, 23, time.Now()},
				{5, 3, 4, 24, time.Now()},
				{6, 3, 5, 25, time.Now()},
			}
		})

		Context("slice split to chunks of 2", func() {
			BeforeEach(func() {
				chunkSize = 2
			})

			It("slice of integers", func() {
				chunks := SplitIntegersToChunks(integers, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(3))
				Expect(chunks[0]).Should(BeEquivalentTo([]int{1, 2}))
				Expect(chunks[1]).Should(BeEquivalentTo([]int{3, 4}))
				Expect(chunks[2]).Should(BeEquivalentTo([]int{5, 6}))
			})

			It("slice of issues", func() {
				chunks := SplitIssuesToChunks(issues, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(3))
				Expect(chunks[0]).Should(BeEquivalentTo([]Issue{issues[0], issues[1]}))
				Expect(chunks[1]).Should(BeEquivalentTo([]Issue{issues[2], issues[3]}))
				Expect(chunks[2]).Should(BeEquivalentTo([]Issue{issues[4], issues[5]}))
			})
		})

		Context("slice split to chunks of 3", func() {
			BeforeEach(func() {
				chunkSize = 3
			})

			It("slice of integers", func() {
				chunks := SplitIntegersToChunks(integers, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(2))
				Expect(chunks[0]).Should(BeEquivalentTo([]int{1, 2, 3}))
				Expect(chunks[1]).Should(BeEquivalentTo([]int{4, 5, 6}))
			})

			It("slice of issues", func() {
				chunks := SplitIssuesToChunks(issues, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(2))
				Expect(chunks[0]).Should(BeEquivalentTo([]Issue{issues[0], issues[1], issues[2]}))
				Expect(chunks[1]).Should(BeEquivalentTo([]Issue{issues[3], issues[4], issues[5]}))
			})
		})

		Context("slice split to chunks of 4", func() {
			BeforeEach(func() {
				chunkSize = 4
			})

			It("slice of integers", func() {
				chunks := SplitIntegersToChunks(integers, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(2))
				Expect(chunks[0]).Should(BeEquivalentTo([]int{1, 2, 3, 4}))
				Expect(chunks[1]).Should(BeEquivalentTo([]int{5, 6}))
			})

			It("slice of issues", func() {
				chunks := SplitIssuesToChunks(issues, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(2))
				Expect(chunks[0]).Should(BeEquivalentTo([]Issue{issues[0], issues[1], issues[2], issues[3]}))
				Expect(chunks[1]).Should(BeEquivalentTo([]Issue{issues[4], issues[5]}))
			})
		})

		Context("slice split to chunks of size", func() {
			BeforeEach(func() {
				chunkSize = len(integers)
			})

			It("slice of integers", func() {
				chunks := SplitIntegersToChunks(integers, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(1))
				Expect(chunks[0]).Should(BeEquivalentTo(integers))
			})

			It("slice of issues", func() {
				chunks := SplitIssuesToChunks(issues, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(1))
				Expect(chunks[0]).Should(BeEquivalentTo(issues))
			})
		})

		Context("slice split to chunks of more than size", func() {
			BeforeEach(func() {
				chunkSize = len(integers) + 1
			})

			It("slice of integers", func() {
				chunks := SplitIntegersToChunks(integers, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(1))
				Expect(chunks[0]).Should(BeEquivalentTo(integers))
			})

			It("slice of issues", func() {
				chunks := SplitIssuesToChunks(issues, chunkSize)
				Expect(len(chunks)).Should(BeEquivalentTo(1))
				Expect(chunks[0]).Should(BeEquivalentTo(issues))
			})
		})

		Context("empty slice split to chunks", func() {
			BeforeEach(func() {
				chunkSize = 42
			})

			It("slice of integers", func() {
				Expect(SplitIntegersToChunks([]int{}, chunkSize)).Should(BeEmpty())
			})

			It("slice of issues", func() {
				Expect(SplitIssuesToChunks([]Issue{}, chunkSize)).Should(BeEmpty())
			})
		})
	})

	Context("SwapMapPairs", func() {
		It("swap key-value pairs of map", func() {
			source := map[int]string{
				1: "one",
				2: "two",
				3: "three",
				4: "four",
				5: "five",
			}

			Expect(SwapMapPairs(source)).Should(BeEquivalentTo(map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
				"four":  4,
				"five":  5,
			}))
		})

		It("swap key-value pairs of empty map", func() {
			Expect(SwapMapPairs(map[int]string{})).Should(BeEmpty())
		})
	})

	Context("DeleteValues", func() {
		var slice []int

		BeforeEach(func() {
			slice = []int{1, 2, 3, 4, 5, 6}
		})

		It("delete values from the source slice", func() {
			Expect(DeleteValues(slice, []int{1, 3, 4, 5, 6})).Should(BeEquivalentTo([]int{2}))
			Expect(DeleteValues(slice, []int{1, 3, 4, 5})).Should(BeEquivalentTo([]int{2, 6}))
			Expect(DeleteValues(slice, []int{5, 6, 6, 7, 4})).Should(BeEquivalentTo([]int{1, 2, 3}))
		})

		It("delete all values from the source slice", func() {
			Expect(DeleteValues(slice, slice)).Should(BeEmpty())
		})

		It("no delete values from the source slice", func() {
			Expect(DeleteValues(slice, []int{})).Should(BeEquivalentTo(slice))
		})
	})

	Context("SliceToMap", func() {
		It("", func() {
			issues := []Issue{
				{1, 1, 1, 42, time.Now()},
				{2, 1, 2, 42, time.Now()},
				{3, 2, 3, 0, time.Now()},
				{4, 1, 3, 23, time.Now()},
			}

			issuesMap := SliceToMap(issues)

			Expect(len(issuesMap)).Should(BeEquivalentTo(len(issues)))

			for i := 0; i < len(issues); i++ {
				issue := issues[i]
				Expect(issuesMap[issue.Id]).Should(BeEquivalentTo(issue))
			}
		})
	})
})
