package saver

import (
	"context"
	"github.com/ozoncp/ocp-issue-api/internal/flusher"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"sync"
	"time"
)

type OverflowPolicy uint8

const (
	ClearBuffer OverflowPolicy = iota
	RemoveOldest
)

type Saver interface {
	Init(ctx context.Context)
	SetOverflowPolicy(policy OverflowPolicy)
	Save(issue models.Issue) error
	Close(ctx context.Context)
}

type saver struct {
	flusher        flusher.Flusher
	timeout        time.Duration
	buffer         []models.Issue
	bufferCapacity uint
	overflowPolicy OverflowPolicy
	mutex          sync.Mutex
	issueCh        chan models.Issue
	doneCh         chan struct{}
}

func (s *saver) Init(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(s.timeout)
		defer ticker.Stop()

		for {
			select {
			case issue := <-s.issueCh:
				s.appendIssue(issue)

			case <-ticker.C:
				s.flushBuffer(ctx)

			case <-s.doneCh:
				return
			}
		}
	}()
}

func (s *saver) SetOverflowPolicy(policy OverflowPolicy) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.overflowPolicy = policy
}

func (s *saver) Save(issue models.Issue) error {
	s.issueCh <- issue
	return nil
}

func (s *saver) Close(ctx context.Context) {
	s.doneCh <- struct{}{}
	close(s.doneCh)
	close(s.issueCh)

	s.flushBuffer(ctx)
}

func New(flusher flusher.Flusher, timeout time.Duration, capacity uint) Saver {
	return &saver{
		flusher:        flusher,
		timeout:        timeout,
		buffer:         make([]models.Issue, 0, capacity),
		bufferCapacity: capacity,
		overflowPolicy: ClearBuffer,
		mutex:          sync.Mutex{},
		issueCh:        make(chan models.Issue),
		doneCh:         make(chan struct{}),
	}
}

func (s *saver) appendIssue(issue models.Issue) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.buffer) == int(s.bufferCapacity) {
		switch s.overflowPolicy {
		case ClearBuffer:
			s.buffer = make([]models.Issue, 0, s.bufferCapacity)

		case RemoveOldest:
			s.buffer = s.buffer[1:]
		}
	}

	s.buffer = append(s.buffer, issue)
}

func (s *saver) flushBuffer(ctx context.Context) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.buffer) > 0 {
		s.buffer = s.flusher.Flush(ctx, s.buffer)
	}
}
