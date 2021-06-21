package metrics

import "github.com/prometheus/client_golang/prometheus"

func newStatusCounters(name string, help string) (prometheus.Counter, prometheus.Counter) {
	var okCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
			ConstLabels: map[string]string{
				"status": "ok",
			},
		},
	)

	var errorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
			ConstLabels: map[string]string{
				"status": "error",
			},
		},
	)

	return okCounter, errorCounter
}

var okCreatedIssues, errorCreatedIssues = newStatusCounters("created_issues", "Number of created issues")
var okUpdatedIssues, errorUpdatedIssues = newStatusCounters("updated_issues", "Number of updated issues")
var okRemovedIssues, errorRemovedIssues = newStatusCounters("removed_issues", "Number of removed issues")

func RegisterMetrics() {
	prometheus.MustRegister(okCreatedIssues, errorCreatedIssues)
	prometheus.MustRegister(okUpdatedIssues, errorUpdatedIssues)
	prometheus.MustRegister(okRemovedIssues, errorRemovedIssues)
}

type Status = uint8

const(
	Ok Status = iota
	Error
)

func AddCreatedIssues(value uint64, status Status) {
	switch status {
	case Ok:
		okCreatedIssues.Add(float64(value))
	case Error:
		errorCreatedIssues.Add(float64(value))
	}
}

func IncCreatedIssues(status Status) {
	switch status {
	case Ok:
		okCreatedIssues.Inc()
	case Error:
		errorCreatedIssues.Inc()
	}
}

func IncUpdatedIssues(status Status) {
	switch status {
	case Ok:
		okUpdatedIssues.Inc()
	case Error:
		errorUpdatedIssues.Inc()
	}
}

func IncRemovedIssues(status Status) {
	switch status {
	case Ok:
		okRemovedIssues.Inc()
	case Error:
		errorRemovedIssues.Inc()
	}
}
