package api_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-issue-api/internal/api"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"github.com/ozoncp/ocp-issue-api/internal/repo"
	desc "github.com/ozoncp/ocp-issue-api/pkg/ocp-issue-api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ = Describe("Api", func() {
	var (
		ctx context.Context

		db   *sql.DB
		mock sqlmock.Sqlmock

		a desc.OcpIssueApiServer

		err error
	)

	BeforeEach(func() {
		ctx = context.Background()
		db, mock, _ = sqlmock.New()
		a = api.New(repo.New(db))
	})

	AfterEach(func() {
		mock.ExpectClose()
		_ = db.Close()
	})

	Context("ListIssues", func() {
		var (
			issues []models.Issue
			rows   *sqlmock.Rows

			req *desc.ListIssuesV1Request
			res *desc.ListIssuesV1Response

			limit  uint64
			offset uint64
		)

		BeforeEach(func() {
			issues = []models.Issue{
				{Id: 1, ClassroomId: 1, TaskId: 1, UserId: 1, Deadline: time.Now()},
				{Id: 2, ClassroomId: 1, TaskId: 1, UserId: 1, Deadline: time.Now()},
				{Id: 3, ClassroomId: 1, TaskId: 1, UserId: 1, Deadline: time.Now()},
				{Id: 4, ClassroomId: 1, TaskId: 1, UserId: 1, Deadline: time.Now()},
				{Id: 5, ClassroomId: 1, TaskId: 1, UserId: 1, Deadline: time.Now()},
			}

			rows = sqlmock.NewRows([]string{"id", "classroom_id", "task_id", "user_id", "deadline"})

			limit = 5
			offset = 0
		})

		JustBeforeEach(func() {
			req = &desc.ListIssuesV1Request{
				Limit:  limit,
				Offset: offset,
			}

			res, err = a.ListIssuesV1(ctx, req)
		})

		Context("list first page of issues", func() {
			BeforeEach(func() {
				limit = 3
				offset = 0

				for _, issue := range issues[0:3] {
					rows.AddRow(issue.Id, issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline)
				}

				mock.ExpectQuery(fmt.Sprintf(`SELECT \* FROM issues LIMIT %d OFFSET %d`, limit, offset)).
					WillReturnRows(rows)
			})

			It("list first page of issues", func() {
				Expect(err).Should(BeNil())
				Expect(len(res.Issues)).Should(BeEquivalentTo(limit))
				Expect(res.Issues).Should(BeEquivalentTo(api.MapFromIssueModels(issues[0:3])))
			})
		})

		Context("list last page of issues", func() {
			BeforeEach(func() {
				limit = 3
				offset = 3

				for _, issue := range issues[3:] {
					rows.AddRow(issue.Id, issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline)
				}

				mock.ExpectQuery(fmt.Sprintf(`SELECT \* FROM issues LIMIT %d OFFSET %d`, limit, offset)).
					WillReturnRows(rows)
			})

			It("list first page of issues", func() {
				Expect(err).Should(BeNil())
				Expect(len(res.Issues)).Should(BeEquivalentTo(2))
				Expect(res.Issues).Should(BeEquivalentTo(api.MapFromIssueModels(issues[3:])))
			})
		})

		Context("list empty issues ", func() {
			BeforeEach(func() {
				mock.ExpectQuery(fmt.Sprintf(`SELECT \* FROM issues LIMIT %d OFFSET %d`, limit, offset)).
					WillReturnRows(rows)
			})

			It("", func() {
				Expect(err).Should(BeNil())
				Expect(res.Issues).Should(BeEmpty())
			})
		})

		Context("failed to list issues", func() {
			BeforeEach(func() {
				mock.ExpectQuery(fmt.Sprintf(`SELECT \* FROM issues LIMIT %d OFFSET %d`, limit, offset)).
					WillReturnError(errors.New("failed to list issues"))
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})
	})

	Context("CreateIssue", func() {
		var (
			req *desc.CreateIssueV1Request
			res *desc.CreateIssueV1Response
		)

		BeforeEach(func() {
			req = &desc.CreateIssueV1Request{
				ClassroomId: 1,
				TaskId:      1,
				UserId:      1,
				Deadline:    timestamppb.New(time.Now()),
			}
		})

		JustBeforeEach(func() {
			res, err = a.CreateIssueV1(ctx, req)
		})

		Context("create an issue", func() {
			var issueId uint64

			BeforeEach(func() {
				issueId = 1
				rows := sqlmock.NewRows([]string{"id"}).AddRow(issueId)

				mock.ExpectQuery("INSERT INTO issues").
					WithArgs(req.ClassroomId, req.TaskId, req.UserId, req.Deadline.AsTime()).
					WillReturnRows(rows)
			})

			It("", func() {
				Expect(err).Should(BeNil())
				Expect(res.IssueId).Should(BeEquivalentTo(issueId))
			})
		})

		Context("don't create an issue with incomplete request data", func() {
			BeforeEach(func() {
				req = &desc.CreateIssueV1Request{}
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})

		Context("failed to create issue", func() {
			BeforeEach(func() {
				mock.ExpectQuery("INSERT INTO issues").
					WithArgs(req.ClassroomId, req.TaskId, req.UserId, req.Deadline.AsTime()).
					WillReturnError(errors.New("failed to create issue"))
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})
	})

	Context("DescribeIssue", func() {
		var (
			rows *sqlmock.Rows
			req  *desc.DescribeIssueV1Request
			res  *desc.DescribeIssueV1Response
		)

		BeforeEach(func() {
			req = &desc.DescribeIssueV1Request{IssueId: 1}

			rows = sqlmock.NewRows([]string{"id", "classroom_id", "task_id", "user_id", "deadline"})

		})

		JustBeforeEach(func() {
			res, err = a.DescribeIssueV1(ctx, req)
		})

		Context("describe issue", func() {
			var issue models.Issue

			BeforeEach(func() {
				issue = models.Issue{
					Id:          req.IssueId,
					ClassroomId: 1,
					TaskId:      1,
					UserId:      1,
					Deadline:    time.Now(),
				}

				rows.AddRow(issue.Id, issue.ClassroomId, issue.TaskId, issue.UserId, issue.Deadline)

				mock.ExpectQuery(`SELECT \* FROM issues WHERE`).
					WithArgs(req.IssueId).
					WillReturnRows(rows)
			})

			It("", func() {
				Expect(err).Should(BeNil())
				Expect(res.Issue).Should(BeEquivalentTo(api.MapFromIssueModel(issue)))
			})
		})

		Context("don't describe not existent issue", func() {
			BeforeEach(func() {
				mock.ExpectQuery(`SELECT \* FROM issues WHERE`).
					WithArgs(req.IssueId).
					WillReturnRows(rows)
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})

		Context("failed to describe issue", func() {
			BeforeEach(func() {
				mock.ExpectQuery(`SELECT \* FROM issues WHERE`).
					WithArgs(req.IssueId).
					WillReturnError(errors.New("failed to describe issue"))
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})
	})

	Context("UpdateIssue", func() {
		var (
			req *desc.UpdateIssueV1Request
			res *desc.UpdateIssueV1Response
		)

		BeforeEach(func() {
			req = &desc.UpdateIssueV1Request{
				IssueId:     1,
				ClassroomId: 1,
				TaskId:      1,
				UserId:      1,
				Deadline:    timestamppb.New(time.Now()),
			}
		})

		JustBeforeEach(func() {
			res, err = a.UpdateIssueV1(ctx, req)
		})

		Context("update issue", func() {
			BeforeEach(func() {
				mock.ExpectExec("UPDATE issues").
					WillReturnResult(sqlmock.NewResult(0, 1))
			})

			It("", func() {
				Expect(err).Should(BeNil())
				Expect(res.Found).Should(BeTrue())
			})
		})

		Context("don't update not existent issue", func() {
			BeforeEach(func() {
				mock.ExpectExec("UPDATE issues").
					WillReturnResult(sqlmock.NewResult(0, 0))
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})

		Context("don't update issue with incomplete request data", func() {
			BeforeEach(func() {
				req = &desc.UpdateIssueV1Request{}
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})

		Context("failed to update issue", func() {
			BeforeEach(func() {
				mock.ExpectExec("UPDATE issues").
					WillReturnError(errors.New("failed to update issue"))
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})
	})

	Context("RemoveIssue", func() {
		var (
			req *desc.RemoveIssueV1Request
			res *desc.RemoveIssueV1Response
		)

		BeforeEach(func() {
			req = &desc.RemoveIssueV1Request{IssueId: 1}
		})

		JustBeforeEach(func() {
			res, err = a.RemoveIssueV1(ctx, req)
		})

		Context("remove issue", func() {
			BeforeEach(func() {
				mock.ExpectExec("DELETE FROM issues WHERE").
					WithArgs(req.IssueId).
					WillReturnResult(sqlmock.NewResult(0, 1))
			})

			It("", func() {
				Expect(err).Should(BeNil())
				Expect(res.Found).Should(BeTrue())
			})
		})

		Context("don't remove not existent issue", func() {
			BeforeEach(func() {
				mock.ExpectExec("DELETE FROM issues WHERE").
					WithArgs(req.IssueId).
					WillReturnResult(sqlmock.NewResult(0, 0))
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})

		Context("failed to remove issue", func() {
			BeforeEach(func() {
				mock.ExpectExec("DELETE FROM issues WHERE").
					WithArgs(req.IssueId).
					WillReturnError(errors.New("failed to remove issue"))
			})

			It("", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(res).Should(BeNil())
			})
		})
	})
})
