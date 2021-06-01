package repo

import "github.com/ozoncp/ocp-issue-api/internal/issues"

type Repo interface {
	AddIssue(issue issues.Issue) (uint64, error)
	AddIssues(issues []issues.Issue) error
	RemoveIssue(issueId uint64) error
	DescribeIssue(issueId uint64) (*issues.Issue, error)
	ListIssues(limit uint64, offset uint64) ([]issues.Issue, error)
}
