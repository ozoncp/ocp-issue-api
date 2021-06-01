package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/ozoncp/ocp-issue-api/internal/flusher"
	"github.com/ozoncp/ocp-issue-api/internal/mocks"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"time"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		flusher  Flusher

		issues []models.Issue
		rest   []models.Issue

		chunkSize int
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

		chunkSize = 2

		issues = []models.Issue{
			{Id: 1, ClassroomId: 1, TaskId: 1, UserId: 42, Deadline: time.Now()},
			{Id: 2, ClassroomId: 1, TaskId: 2, UserId: 42, Deadline: time.Now()},
			{Id: 3, ClassroomId: 2, TaskId: 3, Deadline: time.Now()},
			{Id: 4, ClassroomId: 1, TaskId: 3, UserId: 23, Deadline: time.Now()},
			{Id: 5, ClassroomId: 3, TaskId: 4, UserId: 24, Deadline: time.Now()},
			{Id: 6, ClassroomId: 3, TaskId: 5, UserId: 25, Deadline: time.Now()},
		}
	})

	JustBeforeEach(func() {
		flusher = New(mockRepo, chunkSize)
		rest = flusher.Flush(issues)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save all issues", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddIssues(gomock.Len(chunkSize)).Return(nil).Times(3)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo don't save issues", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddIssues(gomock.Len(chunkSize)).Return(errors.New("can't save issues"))
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(issues))
		})
	})

	Context("repo save not all issues", func() {
		var addedIssuesCount int

		BeforeEach(func() {
			addedIssuesCount = len(issues) - chunkSize

			gomock.InOrder(
				mockRepo.EXPECT().AddIssues(gomock.Len(chunkSize)).Return(nil).Times(2),
				mockRepo.EXPECT().AddIssues(gomock.Len(chunkSize)).Return(errors.New("can't save issues")),
			)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(issues[addedIssuesCount:]))
		})
	})
})
