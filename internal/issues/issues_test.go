package issues_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/ozoncp/ocp-issue-api/internal/issues"
	"time"
)

var _ = Describe("Issues", func() {
	Context("conversion to string", func() {
		It("", func() {
			issue := New(1, 2, 3, 4, time.Now())

			excepted :=  fmt.Sprintf(
				"Issue(Id=%d, ClassroomId=%d, TaskId=%d, UserId=%d, Deadline=%s)",
				issue.Id, issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline,
			)

			Expect(excepted).Should(BeEquivalentTo(issue.String()))
		})
	})
})
