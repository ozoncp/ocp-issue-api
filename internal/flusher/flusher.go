package flusher

import (
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"github.com/ozoncp/ocp-issue-api/internal/repo"
	"github.com/ozoncp/ocp-issue-api/internal/utils"
)

type Flusher interface {
	Flush(issues []models.Issue) []models.Issue
}

type flusher struct {
	repo      repo.Repo
	chuckSize int
}

func (flusher *flusher) Flush(issues []models.Issue) []models.Issue {
	for index, chunk := range utils.SplitIssuesToChunks(issues, flusher.chuckSize) {

		if err := flusher.repo.AddIssues(chunk); err != nil {
			return issues[index*flusher.chuckSize:]
		}
	}

	return nil
}

func New(repo repo.Repo, chunkSize int) Flusher {
	return &flusher{
		repo:      repo,
		chuckSize: chunkSize,
	}
}
