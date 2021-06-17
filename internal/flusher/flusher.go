package flusher

import (
	"context"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"github.com/ozoncp/ocp-issue-api/internal/repo"
	"github.com/ozoncp/ocp-issue-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, issues []models.Issue) []models.Issue
}

type flusher struct {
	repo      repo.Repo
	chuckSize int
}

func (f *flusher) Flush(ctx context.Context, issues []models.Issue) []models.Issue {
	for index, chunk := range utils.SplitIssuesToChunks(issues, f.chuckSize) {
		if _, err := f.repo.AddIssues(ctx, chunk); err != nil {
			return issues[index*f.chuckSize:]
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
