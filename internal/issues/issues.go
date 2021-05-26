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
	ClassroomId uint64
	TaskId      uint64
	UserId      uint64
	Deadline    time.Time
}

func New(id uint64, classroomId uint64, taskId uint64, userId uint64, deadline time.Time) *Issue {
	issue := new(Issue)
	issue.Id = id
	issue.ClassroomId = classroomId
	issue.TaskId = taskId
	issue.UserId = userId
	issue.Deadline = deadline

	return issue
}

func (issue *Issue) String() string {
	return fmt.Sprintf(
		"Issue(Id=%d, ClassroomId=%d, TaskId=%d, UserId=%d, Deadline=%s)",
		issue.Id, issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline.String(),
	)
}
