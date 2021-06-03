package repo

import "github.com/ozoncp/ocp-issue-api/internal/models"

type Repo interface {
	AddIssue(issue models.Issue) (uint64, error)
	AddIssues(issues []models.Issue) error
	RemoveIssue(issueId uint64) error
	DescribeIssue(issueId uint64) (*models.Issue, error)
	ListIssues(limit uint64, offset uint64) ([]models.Issue, error)
}
