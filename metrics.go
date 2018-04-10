package main

import (
	"fmt"

	zsend "github.com/blacked/go-zabbix"
)

var (
	globalMetrics = []string{
		"Aborted_clients",
		"Aborted_connects",
		"Innodb_rows_deleted",
		"Innodb_rows_inserted",
		"Innodb_rows_read",
		"Innodb_rows_updated",
		"Com_begin",
		"Com_commit",
		"Com_rollback",
		"Com_delete",
		"Com_insert",
		"Com_select",
		"Com_update",
		"Queries",
		"Slow_queries",
		"Uptime",
		"Threads_running",
		"Bytes_received",
		"Bytes_sent",
	}

	galeraMetrics = []string{
		"wsrep_cluster_size",
		"wsrep_cluster_state_uuid",
		"wsrep_gcomm_uuid",
		"wsrep_cluster_status",
		"wsrep_connected",
		"wsrep_evs_state",
		"wsrep_last_committed",
		"wsrep_local_bf_aborts",
		"wsrep_local_cert_failures",
		"wsrep_local_state",
		"wsrep_local_state_comment",
		"wsrep_local_state_uuid",
		"wsrep_protocol_version",
		"wsrep_provider_name",
		"wsrep_ready",
	}

	slaveMetrics = []string{
		"Slave_IO_Running",
		"Slave_SQL_Running",
		"Seconds_Behind_Master",
		"Master_Host",
		"Master_Port",
	}
	processMetrics = []string{
		"processlist_count",
		"query_max_time",
	}
)

func makePrefix(prefix, key string) string {
	return fmt.Sprintf(
		"%s.%s", prefix, key,
	)
}

func createMetrics(
	hostname string,
	stats map[string]string,
	statsType string,
	filterMetrics []string,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	for _, filterMetric := range filterMetrics {

		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf("%s.[%s]", filterMetric, statsType),
				),
				stats[filterMetric],
			),
		)
	}

	return metrics
}
