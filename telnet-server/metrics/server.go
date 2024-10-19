package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	connectionsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "telnet_server_connection_total",
		Help: "The total number of connections",
	})
	connectionErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "telnet_server_connection_errors_total",
		Help: "The total number of errors",
	})
	unknownCommands = promauto.NewCounter(prometheus.CounterOpts{
		Name: "telnet_server_unknown_commands_total",
		Help: "The total number of unknown commands entered",
	})

	connectionsActive = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "telnet_server_active_connections",
		Help: "The number of active connections",
	})
)

// MetricServer holds state for our Prometheus metrics server
type MetricServer struct {
	port     string
	endPoint string
	logger   *log.Logger
}

// New creates a new metric server
func New(port string, logger *log.Logger) *MetricServer {
	return &MetricServer{port: port, endPoint: "/metrics", logger: logger}
}

// ListenAndServeMetrics runs our metrics server
func (m *MetricServer) ListenAndServeMetrics() {
	http.Handle(m.endPoint, promhttp.Handler())
	m.logger.Printf("Metrics endpoint listening on %s\n", m.port)
	http.ListenAndServe(m.port, nil)
}

// Increment_Connection_Errors += 1
func (m *MetricServer) IncrementConnectionErrors() {
	connectionErrors.Inc()
}

//Increment_Connections_Processed += 1
func (m *MetricServer) IncrementConnectionsProcessed() {
	connectionsProcessed.Inc()
}

// Increment_Unknown_Commands += 1
func (m *MetricServer) IncrementUnknownCommands(cmd string) {
	unknownCommands.Inc()
}

//Increment_Active_Connections += 1
func (m *MetricServer) IncrementActiveConnections() {
	connectionsActive.Inc()
}

//DecrementActiveConnections -= 1
func (m *MetricServer) DecrementActiveConnections() {
	connectionsActive.Dec()
}
