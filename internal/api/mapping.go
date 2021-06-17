package api

import (
	"github.com/ozoncp/ocp-issue-api/internal/models"
	desc "github.com/ozoncp/ocp-issue-api/pkg/ocp-issue-api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapFromIssueModel(issue models.Issue) *desc.Issue {
	return &desc.Issue{
		Id:          issue.Id,
		ClassroomId: issue.ClassroomId,
		TaskId:      issue.TaskId,
		UserId:      issue.UserId,
		Deadline:    timestamppb.New(issue.Deadline),
	}
}

func mapFromIssueModels(issues []models.Issue) []*desc.Issue {
	mapped := make([]*desc.Issue, 0, len(issues))

	for _, issue := range issues {
		mapped = append(mapped, mapFromIssueModel(issue))
	}

	return mapped
}

func mapFromCreateIssueRequest(req *desc.CreateIssueV1Request) models.Issue {
	return models.Issue{
		ClassroomId: req.ClassroomId,
		TaskId:      req.TaskId,
		UserId:      req.UserId,
		Deadline:    req.Deadline.AsTime(),
	}
}

func mapFromUpdateIssueRequest(req *desc.UpdateIssueV1Request) models.Issue {
	return models.Issue{
		Id:          req.IssueId,
		ClassroomId: req.ClassroomId,
		TaskId:      req.TaskId,
		UserId:      req.UserId,
		Deadline:    req.Deadline.AsTime(),
	}
}
