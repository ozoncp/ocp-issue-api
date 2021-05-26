package tests

import (
	"fmt"
	"github.com/ozoncp/ocp-issue-api/internal/issues"
	"testing"
	"time"
)

func TestIssueString(t *testing.T) {
	issue := issues.New(1, 2, 3, 4, time.Now())

	excepted :=  fmt.Sprintf(
		"Issue(Id=%d, ClassroomId=%d, TaskId=%d, UserId=%d, Deadline=%s)",
		issue.Id, issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline,
	)

	assertEqual(t, excepted, issue.String())
}
