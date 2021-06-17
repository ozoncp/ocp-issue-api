package metrics

import "github.com/prometheus/client_golang/prometheus"

var createdIssues = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "created_issues",
		Help: "Number of created issues",
	},
)

var updatedIssues = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "updated_issues",
		Help: "Number of updated issues",
	},
)

var removedIssues = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "removed_issues",
		Help: "Number of removed issues",
	},
)

func RegisterMetrics() {
	prometheus.MustRegister(createdIssues)
	prometheus.MustRegister(updatedIssues)
	prometheus.MustRegister(removedIssues)
}

func AddCreatedIssues(value uint64) {
	createdIssues.Add(float64(value))
}

func IncCreatedIssues() {
	createdIssues.Inc()
}

func IncUpdatedIssues() {
	updatedIssues.Inc()
}

func IncRemovedIssues() {
	removedIssues.Inc()
}
