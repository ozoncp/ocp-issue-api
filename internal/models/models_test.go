package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"time"
)

var _ = Describe("Models", func() {
	var deadline time.Time
	var issue models.Issue

	BeforeEach(func() {
		deadline, _ = time.Parse(time.RFC3339, "2021-09-11T08:35:05+07:00")
		issue = models.Issue{Id: 1, ClassroomId: 2, TaskId: 3, UserId: 4, Deadline: deadline}
	})

	Context("Issues", func() {
		var issueString string

		BeforeEach(func() {
			issueString = "Issue(Id=1, ClassroomId=2, TaskId=3, UserId=4, Deadline=2021-09-11 08:35:05 +0700 +07)"
		})

		It("conversion to string", func() {
			Expect(issue.String()).Should(BeEquivalentTo(issueString))
		})
	})
})
