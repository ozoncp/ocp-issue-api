package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-issue-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/event_notifier_mock.go -package=mocks github.com/ozoncp/ocp-issue-api/internal/events EventNotifier
//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-issue-api/internal/repo Repo
