package flusher

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-issue-api/internal/mocks"
	"github.com/ozoncp/ocp-issue-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		ctx context.Context

		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		f        Flusher

		issues []models.Issue
		rest   []models.Issue

		chunkSize int
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

		chunkSize = 2
		issues = []models.Issue{{Id: 1}, {Id: 2}, {Id: 3}, {Id: 4}, {Id: 5}, {Id: 6}}
	})

	JustBeforeEach(func() {
		f = New(mockRepo, chunkSize)
		rest = f.Flush(ctx, issues)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save all issues", func() {
		BeforeEach(func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 1}, {Id: 2}}).Return(nil),
				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 3}, {Id: 4}}).Return(nil),
				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 5}, {Id: 6}}).Return(nil),
			)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo don't save issues", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 1}, {Id: 2}}).Return(errors.New("can't save issues"))
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(issues))
		})
	})

	Context("repo save not all issues", func() {
		BeforeEach(func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 1}, {Id: 2}}).Return(nil),
				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 3}, {Id: 4}}).Return(nil),
				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 5}, {Id: 6}}).Return(errors.New("can't save issues")),
			)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo([]models.Issue{{Id: 5}, {Id: 6}}))
		})
	})
})
