package models

import (
	"fmt"
	"time"
)

type Issue struct {
	Id          uint64
	ClassroomId uint64
	TaskId      uint64
	UserId      uint64
	Deadline    time.Time
}

func (issue *Issue) String() string {
	return fmt.Sprintf(
		"Issue(Id=%d, ClassroomId=%d, TaskId=%d, UserId=%d, Deadline=%s)",
		issue.Id, issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline.String(),
	)
}
