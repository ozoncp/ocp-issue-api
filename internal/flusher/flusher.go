package flusher

import (
	"context"
	"github.com/ozoncp/ocp-issue-api/internal/events"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"github.com/ozoncp/ocp-issue-api/internal/repo"
	"github.com/ozoncp/ocp-issue-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, issues []models.Issue) []models.Issue
}

type flusher struct {
	repo      repo.Repo
	notifier  events.EventNotifier
	chuckSize int
}

func (f *flusher) Flush(ctx context.Context, issues []models.Issue) []models.Issue {
	for index, chunk := range utils.SplitIssuesToChunks(issues, f.chuckSize) {
		issueIds, err := f.repo.AddIssues(ctx, chunk)

		for _, issueId := range issueIds {
			f.notifier.Notify(issueId, events.Created)
		}

		if err != nil {
			return issues[index*f.chuckSize:]
		}
	}

	return nil
}

func NewFlusher(repo repo.Repo, notifier events.EventNotifier, chunkSize int) Flusher {
	return &flusher{
		repo:      repo,
		notifier:  notifier,
		chuckSize: chunkSize,
	}
}
