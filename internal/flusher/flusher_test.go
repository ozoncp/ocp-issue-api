package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-issue-api/internal/flusher"
	"github.com/ozoncp/ocp-issue-api/internal/mocks"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"time"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		f        flusher.Flusher

		issues []models.Issue
		rest   []models.Issue

		chunkSize int
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

		chunkSize = 2

		f = flusher.New(mockRepo, chunkSize)

		issues = []models.Issue{
			{1, 1, 1, 42, time.Now()},
			{2, 1, 2, 42, time.Now()},
			{3, 2, 3, 0, time.Now()},
			{4, 1, 3, 23, time.Now()},
			{5, 3, 4, 24, time.Now()},
			{6, 3, 5, 25, time.Now()},
		}

		rest = f.Flush(issues)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save all models", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddIssues(gomock.Any()).Return(nil).MinTimes(3)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})
})
