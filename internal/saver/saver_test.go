package saver_test

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-issue-api/internal/mocks"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"github.com/ozoncp/ocp-issue-api/internal/saver"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		ctx context.Context

		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher

		timeout  time.Duration
		capacity uint
		s        saver.Saver
		issues   []models.Issue
	)

	BeforeEach(func() {
		ctx = context.Background()

		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)

		timeout = 10 * time.Millisecond
		issues = []models.Issue{{Id: 1}, {Id: 2}, {Id: 3}}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("with big enough capacity", func() {
		BeforeEach(func() {
			capacity = 10
			s = saver.New(mockFlusher, timeout, capacity)
			s.Init(ctx)
		})

		AfterEach(func() {
			s.Close(ctx)
		})

		Context("flush by closing saver", func() {
			BeforeEach(func() {
				mockFlusher.EXPECT().Flush(ctx, []models.Issue{{Id: 1}, {Id: 2}, {Id: 3}}).Return(nil).Times(1)
			})

			It("", func() {
				for _, issue := range issues {
					Expect(s.Save(issue)).Should(BeNil())
				}
			})
		})

		Context("flush by timeout", func() {
			BeforeEach(func() {
				gomock.InOrder(
					mockFlusher.EXPECT().Flush(ctx, []models.Issue{{Id: 1}}).Return(nil).Times(1),
					mockFlusher.EXPECT().Flush(ctx, []models.Issue{{Id: 2}}).Return(nil).Times(1),
					mockFlusher.EXPECT().Flush(ctx, []models.Issue{{Id: 3}}).Return(nil).Times(1),
				)
			})

			It("", func() {
				for _, issue := range issues {
					Expect(s.Save(issue)).Should(BeNil())
					time.Sleep(timeout)
				}
			})
		})
	})

	Context("ClearBuffer overflow policy", func() {
		BeforeEach(func() {
			capacity = 2
			s = saver.New(mockFlusher, timeout, capacity)
			s.SetOverflowPolicy(saver.ClearBuffer)
			s.Init(ctx)

			mockFlusher.EXPECT().Flush(ctx, []models.Issue{{Id: 3}}).Return(nil).Times(1)
		})

		AfterEach(func() {
			s.Close(ctx)
		})

		It("flush by closing saver", func() {
			for _, issue := range issues {
				Expect(s.Save(issue)).Should(BeNil())
			}
		})

		It("flush by timeout", func() {
			for _, issue := range issues {
				Expect(s.Save(issue)).Should(BeNil())
			}

			time.Sleep(timeout)
		})
	})

	Context("RemoveOldest overflow policy", func() {
		BeforeEach(func() {
			capacity = 2
			s = saver.New(mockFlusher, timeout, capacity)
			s.SetOverflowPolicy(saver.RemoveOldest)
			s.Init(ctx)

			mockFlusher.EXPECT().Flush(ctx, []models.Issue{{Id: 2}, {Id: 3}}).Return(nil).Times(1)
		})

		AfterEach(func() {
			s.Close(ctx)
		})

		It("flush by closing saver", func() {
			for _, issue := range issues {
				Expect(s.Save(issue)).Should(BeNil())
			}
		})

		It("flush by timeout", func() {
			for _, issue := range issues {
				Expect(s.Save(issue)).Should(BeNil())
			}

			time.Sleep(timeout)
		})
	})
})
