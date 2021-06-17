package api

import (
	"context"
	"github.com/ozoncp/ocp-issue-api/internal/repo"
	desc "github.com/ozoncp/ocp-issue-api/pkg/ocp-issue-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOcpIssueApiServer
	repo repo.Repo
}

func (a *api) ListIssuesV1(ctx context.Context, req *desc.ListIssuesV1Request) (*desc.ListIssuesV1Response, error) {
	log.Debug().
		Str("request", "ListIssues").
		Uint("version", 1).
		Uint64("limit", req.Limit).
		Uint64("offset", req.Offset).
		Msg("invoke handle")

	if err := req.Validate(); err != nil {
		return nil, desc.ListIssuesV1RequestValidationError{}
	}

	issues, err := a.repo.ListIssues(ctx, req.Limit, req.Offset)

	if err != nil {
		log.Error().Msgf("failed to list issues: %v", err)
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &desc.ListIssuesV1Response{Issues: mapFromIssueModels(issues)}, nil
}

func (a *api) CreateIssueV1(ctx context.Context, req *desc.CreateIssueV1Request) (*desc.CreateIssueV1Response, error) {
	log.Debug().
		Str("request", "CreateIssue").
		Uint("version", 1).
		Uint64("classroom_id", req.ClassroomId).
		Uint64("task_id", req.TaskId).
		Uint64("user_id", req.UserId).
		Time("deadline", req.Deadline.AsTime()).
		Msg("invoke handle")

	if err := req.Validate(); err != nil {
		return nil, desc.CreateIssueV1RequestValidationError{}
	}

	issueId, err := a.repo.AddIssue(ctx, mapFromCreateIssueRequest(req))

	if err != nil {
		log.Error().Msgf("failed to create issue: %v", err)
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &desc.CreateIssueV1Response{IssueId: issueId}, nil
}

func (a *api) DescribeIssueV1(ctx context.Context, req *desc.DescribeIssueV1Request) (*desc.DescribeIssueV1Response, error) {
	log.Debug().
		Str("request", "DescribeIssue").
		Uint("version", 1).
		Uint64("issue_id", req.IssueId).
		Msg("invoke handle")

	if err := req.Validate(); err != nil {
		return nil, desc.DescribeIssueV1RequestValidationError{}
	}

	issue, err := a.repo.DescribeIssue(ctx, req.IssueId)

	if err != nil {
		log.Error().
			Uint64("issue_id", req.IssueId).
			Msgf("failed to describe issue: %v", err)

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &desc.DescribeIssueV1Response{Issue: mapFromIssueModel(*issue)}, nil
}

func (a *api) UpdateIssueV1(ctx context.Context, req *desc.UpdateIssueV1Request) (*desc.UpdateIssueV1Response, error) {
	log.Debug().
		Str("request", "UpdateIssue").
		Uint("version", 1).
		Uint64("issue_id", req.IssueId).
		Uint64("classroom_id", req.ClassroomId).
		Uint64("task_id", req.TaskId).
		Uint64("user_id", req.UserId).
		Time("deadline", req.Deadline.AsTime()).
		Msg("invoke handle")

	if err := req.Validate(); err != nil {
		return nil, desc.UpdateIssueV1RequestValidationError{}
	}

	err := a.repo.UpdateIssue(ctx, mapFromUpdateIssueRequest(req))

	if err != nil {
		log.Error().
			Uint64("issue_id", req.IssueId).
			Msgf("failed to update issue: %v", err)

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &desc.UpdateIssueV1Response{Found: true}, nil
}

func (a *api) RemoveIssueV1(ctx context.Context, req *desc.RemoveIssueV1Request) (*desc.RemoveIssueV1Response, error) {
	log.Debug().
		Str("request", "RemoveIssue").
		Uint("version", 1).
		Uint64("issue_id", req.IssueId).
		Msg("invoke handle")

	if err := req.Validate(); err != nil {
		return nil, desc.RemoveIssueV1RequestValidationError{}
	}

	err := a.repo.RemoveIssue(ctx, req.IssueId)

	if err != nil {
		log.Error().
			Uint64("issue_id", req.IssueId).
			Msgf("failed to remove issue: %v", err)

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &desc.RemoveIssueV1Response{Found: true}, nil
}

func New(repo repo.Repo) desc.OcpIssueApiServer {
	return &api{repo: repo}
}
