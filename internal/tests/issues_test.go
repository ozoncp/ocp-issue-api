package tests

import (
	"fmt"
	"github.com/ozoncp/ocp-issue-api/internal/issues"
	"testing"
	"time"
)

func TestIssueString(t *testing.T) {
	issue := issues.New(1, 2, "An issue")
	issue.UserId = 42
	issue.Deadline = time.Now()
	issue.State = issues.InProgress

	excepted :=  fmt.Sprintf(
		"Issue(Id=%d, TaskId=%d, UserId=%d, Description=\"%s\", State=%d, Deadline=%s)",
		issue.Id, issue.TaskId, issue.UserId, issue.Description, issue.State, issue.Deadline,
	)

	assertEqual(t, excepted, issue.String())
}
