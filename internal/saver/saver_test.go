package saver_test

import (
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
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher
		timeout     time.Duration
		capacity    uint
		s           saver.Saver
		issues      []models.Issue
	)

	BeforeEach(func() {
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
			s.Init()
		})

		AfterEach(func() {
			s.Close()
		})

		Context("flush by closing saver", func() {
			BeforeEach(func() {
				mockFlusher.EXPECT().Flush([]models.Issue{{Id: 1}, {Id: 2}, {Id: 3}}).Return(nil).Times(1)
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
					mockFlusher.EXPECT().Flush([]models.Issue{{Id: 1}}).Return(nil).Times(1),
					mockFlusher.EXPECT().Flush([]models.Issue{{Id: 2}}).Return(nil).Times(1),
					mockFlusher.EXPECT().Flush([]models.Issue{{Id: 3}}).Return(nil).Times(1),
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
			s.Init()

			mockFlusher.EXPECT().Flush([]models.Issue{{Id: 3}}).Return(nil).Times(1)
		})

		AfterEach(func() {
			s.Close()
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
			s.Init()

			mockFlusher.EXPECT().Flush([]models.Issue{{Id: 2}, {Id: 3}}).Return(nil).Times(1)
		})

		AfterEach(func() {
			s.Close()
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
