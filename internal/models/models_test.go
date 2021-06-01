package models_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"time"
)

var _ = Describe("Models", func() {
	Context("Issues", func() {
		It("conversion to string", func() {
			issue := models.Issue{Id: 1, ClassroomId: 2, TaskId: 3, UserId: 4, Deadline: time.Now()}

			excepted :=  fmt.Sprintf(
				"Issue(Id=%d, ClassroomId=%d, TaskId=%d, UserId=%d, Deadline=%s)",
				issue.Id, issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline,
			)

			Expect(excepted).Should(BeEquivalentTo(issue.String()))
		})
	})
})
