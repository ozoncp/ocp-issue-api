package issues

import (
	"fmt"
	"time"
)

type State uint

const (
	Created State = iota
	Assigned
	InProgress
	Completed
	OnReview
	Approved
)

type Issue struct {
	Id          uint64
	TaskId      uint64
	UserId      uint64
	Description string
	State       State
	Deadline    time.Time
}

func New(id uint64, taskId uint64, description string) *Issue {
	issue := new(Issue)
	issue.Id = id
	issue.TaskId = taskId
	issue.Description = description
	issue.State = Created

	return issue
}

func (issue *Issue) String() string {
	return fmt.Sprintf(
		"Issue(Id=%d, TaskId=%d, UserId=%d, Description=\"%s\", State=%d, Deadline=%s)",
		issue.Id, issue.TaskId, issue.UserId, issue.Description, issue.State, issue.Deadline.String(),
	)
}
