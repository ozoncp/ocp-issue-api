package repo

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"github.com/rs/zerolog/log"
)

type Repo interface {
	AddIssue(ctx context.Context, issue models.Issue) (uint64, error)
	AddIssues(ctx context.Context, issues []models.Issue) ([]uint64, error)
	UpdateIssue(ctx context.Context, issue models.Issue) error
	RemoveIssue(ctx context.Context, issueId uint64) error
	DescribeIssue(ctx context.Context, issueId uint64) (*models.Issue, error)
	ListIssues(ctx context.Context, limit uint64, offset uint64) ([]models.Issue, error)
}

const tableName = "issues"

type repo struct {
	db      *sql.DB
}

func (r *repo) AddIssue(ctx context.Context, issue models.Issue) (uint64, error) {
	query := sq.Insert(tableName).
		Columns("class_room_id", "task_id", "user_id", "deadline").
		Values(issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	if err := query.QueryRowContext(ctx).Scan(&issue.Id); err != nil {
		return 0, err
	}

	return issue.Id, nil
}

func (r *repo) AddIssues(ctx context.Context, issues []models.Issue) ([]uint64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AddIssues")
	defer span.Finish()

	query := sq.Insert(tableName).
		Columns("class_room_id", "task_id", "user_id", "deadline").
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, issue := range issues {
		query = query.Values(issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline)
	}

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	var issueIds []uint64

	for rows.Next() {
		var issueId uint64
		err = rows.Scan(&issueId)

		if err == nil {
			issueIds = append(issueIds, issueId)
		}
	}

	span.SetTag("issues-count", len(issues))

	return issueIds, err
}

func (r *repo) UpdateIssue(ctx context.Context, issue models.Issue) error {
	query := sq.Update(tableName).
		Where(sq.Eq{"id": issue.Id}).
		Set("class_room_id", issue.ClassroomId).
		Set("task_id", issue.TaskId).
		Set("user_id", issue.UserId).
		Set("deadline", issue.Deadline).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	res, err := query.ExecContext(ctx)

	if res == nil {
		return err
	}

	updatedRows, err := res.RowsAffected()

	switch updatedRows {
	case 0:
		return errors.New("issue not found")
	case 1:
		return nil
	}

	return err
}

func (r *repo) RemoveIssue(ctx context.Context, issueId uint64) error {
	query := sq.Delete(tableName).
		Where(sq.Eq{"id": issueId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	res, err := query.ExecContext(ctx)

	if res == nil {
		return err
	}

	removedRows, err := res.RowsAffected()

	switch removedRows {
	case 0:
		return errors.New("issue not found")
	case 1:
		return nil
	}

	return err
}

func (r *repo) DescribeIssue(ctx context.Context, issueId uint64) (*models.Issue, error) {
	query := sq.Select("*").
		From(tableName).
		Where(sq.Eq{"id": issueId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var issue models.Issue

	err := query.QueryRowContext(ctx).Scan(&issue.Id, &issue.ClassroomId, &issue.TaskId, &issue.UserId, &issue.Deadline)

	if err != nil {
		return nil, errors.New("issue not found")
	}

	return &issue, nil
}

func (r *repo) ListIssues(ctx context.Context, limit uint64, offset uint64) ([]models.Issue, error) {
	query := sq.Select("*").
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	var issues []models.Issue

	for rows.Next() {
		var issue models.Issue
		err = rows.Scan(&issue.Id, &issue.ClassroomId, &issue.TaskId, &issue.UserId, &issue.Deadline)

		if err == nil {
			issues = append(issues, issue)
		} else {
			log.Error().Msgf("ailed to scan issue: %v", err)
		}
	}

	return issues, err
}

func NewRepo(db *sql.DB) Repo {
	return &repo{db: db}
}
