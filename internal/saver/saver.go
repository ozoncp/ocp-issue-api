package saver

import (
	"github.com/ozoncp/ocp-issue-api/internal/flusher"
	"github.com/ozoncp/ocp-issue-api/internal/models"
	"time"
)

type OverflowPolicy uint8

const (
	ClearBuffer OverflowPolicy = iota
	RemoveOldest
)

type Saver interface {
	Init()
	SetOverflowPolicy(policy OverflowPolicy)
	Save(issue models.Issue) error
	Close()
}

type saver struct {
	flusher        flusher.Flusher
	timeout        time.Duration
	buffer         []models.Issue
	overflowPolicy OverflowPolicy
	issueCh        chan models.Issue
	doneCh         chan struct{}
}

func (s *saver) Init() {
	go func() {
		ticker := time.NewTicker(s.timeout)

		for {
			select {
			case issue, ok := <-s.issueCh:
				if ok {
					if len(s.buffer) == cap(s.buffer) {
						switch s.overflowPolicy {
						case ClearBuffer:
							s.buffer = make([]models.Issue, 0, cap(s.buffer))
							break

						case RemoveOldest:
							s.buffer = s.buffer[1:]
							break
						}
					}

					s.buffer = append(s.buffer, issue)
				} else {
					if len(s.buffer) > 0 {
						_ = s.flusher.Flush(s.buffer)
					}

					s.doneCh <- struct{}{}
					return
				}

			case <-ticker.C:
				s.buffer = s.flusher.Flush(s.buffer)
			}
		}
	}()
}

func (s *saver) SetOverflowPolicy(policy OverflowPolicy) {
	s.overflowPolicy = policy
}

func (s *saver) Save(issue models.Issue) error {
	s.issueCh <- issue
	return nil
}

func (s *saver) Close() {
	close(s.issueCh)
	<-s.doneCh
}

func New(flusher flusher.Flusher, timeout time.Duration, capacity uint) Saver {
	return &saver{
		flusher:        flusher,
		timeout:        timeout,
		buffer:         make([]models.Issue, 0, capacity),
		overflowPolicy: ClearBuffer,
		issueCh:        make(chan models.Issue),
		doneCh:         make(chan struct{}),
	}
}
