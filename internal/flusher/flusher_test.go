package flusher

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-issue-api/internal/events"
	"github.com/ozoncp/ocp-issue-api/internal/mocks"
	"github.com/ozoncp/ocp-issue-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		ctx context.Context

		ctrl              *gomock.Controller
		mockRepo          *mocks.MockRepo
		mockEventNotifier *mocks.MockEventNotifier

		f Flusher

		issues []models.Issue
		rest   []models.Issue

		chunkSize int
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		mockEventNotifier = mocks.NewMockEventNotifier(ctrl)

		chunkSize = 2
		issues = []models.Issue{{Id: 1}, {Id: 2}, {Id: 3}, {Id: 4}, {Id: 5}, {Id: 6}}
	})

	JustBeforeEach(func() {
		f = NewFlusher(mockRepo, mockEventNotifier, chunkSize)
		rest = f.Flush(ctx, issues)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save all issues", func() {
		BeforeEach(func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 1}, {Id: 2}}).Return([]uint64{1, 2}, nil),
				mockEventNotifier.EXPECT().Notify(uint64(1), events.Created),
				mockEventNotifier.EXPECT().Notify(uint64(2), events.Created),

				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 3}, {Id: 4}}).Return([]uint64{3, 4}, nil),
				mockEventNotifier.EXPECT().Notify(uint64(3), events.Created),
				mockEventNotifier.EXPECT().Notify(uint64(4), events.Created),

				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 5}, {Id: 6}}).Return([]uint64{5, 6}, nil),
				mockEventNotifier.EXPECT().Notify(uint64(5), events.Created),
				mockEventNotifier.EXPECT().Notify(uint64(6), events.Created),
			)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo don't save issues", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 1}, {Id: 2}}).
				Return(nil, errors.New("failed to save issues"))
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(issues))
		})
	})

	Context("repo save not all issues", func() {
		BeforeEach(func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 1}, {Id: 2}}).Return([]uint64{1, 2}, nil),
				mockEventNotifier.EXPECT().Notify(uint64(1), events.Created),
				mockEventNotifier.EXPECT().Notify(uint64(2), events.Created),

				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 3}, {Id: 4}}).Return([]uint64{3, 4}, nil),
				mockEventNotifier.EXPECT().Notify(uint64(3), events.Created),
				mockEventNotifier.EXPECT().Notify(uint64(4), events.Created),

				mockRepo.EXPECT().AddIssues(ctx, []models.Issue{{Id: 5}, {Id: 6}}).
					Return([]uint64{6}, errors.New("failed to save issues")),

				mockEventNotifier.EXPECT().Notify(uint64(6), events.Created),
			)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo([]models.Issue{{Id: 5}, {Id: 6}}))
		})
	})
})
