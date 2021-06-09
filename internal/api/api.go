package api

import (
	"context"
	desc "github.com/ozoncp/ocp-issue-api/pkg/ocp-issue-api"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOcpIssueApiServer
}

func (a *api) ListIssuesV1(ctx context.Context, req *desc.ListIssuesV1Request) (*desc.ListIssuesV1Response, error) {
	log.Info().Msg("ListIssues V1")

	if err := req.Validate(); err != nil {
		return nil, desc.ListIssuesV1RequestValidationError{}
	}

	return &desc.ListIssuesV1Response{Issues: []*desc.Issue{{Id: 1}, {Id: 2}}}, nil
}

func (a *api) CreateIssueV1(ctx context.Context, req *desc.CreateIssueV1Request) (*desc.CreateIssueV1Response, error) {
	log.Info().Msg("CreateIssue V1")

	if err := req.Validate(); err != nil {
		return nil, desc.CreateIssueV1RequestValidationError{}
	}

	return &desc.CreateIssueV1Response{IssueId: 1}, nil
}

func (a *api) DescribeIssueV1(ctx context.Context, req *desc.DescribeIssueV1Request, ) (*desc.DescribeIssueV1Response, error) {
	log.Info().Msg("DescribeIssue V1")

	if err := req.Validate(); err != nil {
		return nil, desc.DescribeIssueV1RequestValidationError{}
	}

	issue := &desc.Issue{Id: 1, ClassroomId: 1, TaskId: 1, UserId: 1, Deadline: timestamppb.Now()}

	return &desc.DescribeIssueV1Response{Issue: issue}, nil
}

func (a *api) UpdateIssueV1(ctx context.Context, req *desc.UpdateIssueV1Request) (*desc.UpdateIssueV1Response, error) {
	log.Info().Msg("UpdateIssue V1")

	if err := req.Validate(); err != nil {
		return nil, desc.UpdateIssueV1RequestValidationError{}
	}

	return &desc.UpdateIssueV1Response{Found: false}, nil
}

func (a *api) RemoveIssueV1(ctx context.Context, req *desc.RemoveIssueV1Request) (*desc.RemoveIssueV1Response, error) {
	log.Info().Msg("RemoveIssue V1")

	if err := req.Validate(); err != nil {
		return nil, desc.RemoveIssueV1RequestValidationError{}
	}

	return &desc.RemoveIssueV1Response{Found: false}, nil
}

func New() desc.OcpIssueApiServer {
	return &api{}
}
