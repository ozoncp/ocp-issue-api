package events

import "time"

type EventNotifier interface {
	Notify(issueId uint64, eventType IssueEventType)
	Close()
}

type notifier struct {
	eventCh chan IssueEvent
}

func (n *notifier) Notify(issueId uint64, eventType IssueEventType) {
	n.eventCh <- IssueEvent{
		Type: eventType,
		Body: IssueEventBody{
			IssueId:   issueId,
			Timestamp: time.Now().UTC().Unix(),
		},
	}
}

func (n *notifier) Close() {
	close(n.eventCh)
}

func NewNotifier(eventCh chan IssueEvent) EventNotifier {
	return &notifier{
		eventCh: eventCh,
	}
}
