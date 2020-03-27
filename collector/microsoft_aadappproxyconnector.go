// +build windows

package collector

import (
	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

func init() {
	registerCollector("microsoft_aadappproxyconnector", NewAddAppProxyConnectorCollector)
}

// A MicrosoftAADAppProxyConnectorCollector is a Prometheus collector for WMI Win32_PerfRawData_NETFramework_NETCLRSecurity metrics
type MicrosoftAADAppProxyConnectorCollector struct {
	Numbercurrentactivebackendwebsockets *prometheus.Desc
	NumbernewbackendwebsocketsPersec     *prometheus.Desc
	Numberrequests                       *prometheus.Desc
	NumberrequestsPersec                 *prometheus.Desc
	Numberresponses                      *prometheus.Desc
	NumberresponsesPersec                *prometheus.Desc
	Numbertransactionscompletions        *prometheus.Desc
	NumbertransactionscompletionsPersec  *prometheus.Desc
	Numbertransactionsfailed             *prometheus.Desc
	NumbertransactionsfailedPersec       *prometheus.Desc
}

// NewAddAppProxyConnectorCollector ...
func NewAddAppProxyConnectorCollector() (Collector, error) {
	const subsystem = "microsoft_aadappproxyconnector"
	return &MicrosoftAADAppProxyConnectorCollector{
		Numbercurrentactivebackendwebsockets: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_current_active_backend_websockets"),
			"Displays the total number of current connected websockets.",
			nil,
			nil,
		),
		NumbernewbackendwebsocketsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_new_backend_websockets_persec"),
			"",
			nil,
			nil,
		),
		Numberrequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_requests"),
			"",
			nil,
			nil,
		),
		NumberrequestsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_requests_persec"),
			"",
			nil,
			nil,
		),
		Numberresponses: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_responses"),
			"",
			nil,
			nil,
		),
		NumberresponsesPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_responses_persec"),
			"",
			nil,
			nil,
		),
		Numbertransactionscompletions: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_transactions_completions"),
			"",
			nil,
			nil,
		),
		NumbertransactionscompletionsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_transactions_completions_persec"),
			"",
			nil,
			nil,
		),
		Numbertransactionsfailed: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_transactions_failed"),
			"",
			nil,
			nil,
		),
		NumbertransactionsfailedPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "number_transactions_failed_Persec"),
			"",
			nil,
			nil,
		),
	}, nil

}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *MicrosoftAADAppProxyConnectorCollector) Collect(ctx *ScrapeContext, ch chan<- prometheus.Metric) error {
	if desc, err := c.collect(ch); err != nil {
		log.Error("failed collecting Win32_PerfFormattedData_MicrosoftAADAppProxyConnector_MicrosoftAADAppProxyConnector metrics:", desc, err)
		return err
	}
	return nil
}

type Win32_PerfFormattedData_MicrosoftAADAppProxyConnector_MicrosoftAADAppProxyConnector struct {
	Numbercurrentactivebackendwebsockets uint64
	NumbernewbackendwebsocketsPersec     uint64
	Numberrequests                       uint64
	NumberrequestsPersec                 uint64
	Numberresponses                      uint64
	NumberresponsesPersec                uint64
	Numbertransactionscompletions        uint64
	NumbertransactionscompletionsPersec  uint64
	Numbertransactionsfailed             uint64
	NumbertransactionsfailedPersec       uint64
}

func (c *MicrosoftAADAppProxyConnectorCollector) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {
	var dst []Win32_PerfFormattedData_MicrosoftAADAppProxyConnector_MicrosoftAADAppProxyConnector

	q := queryAll(&dst)
	if err := wmi.Query(q, &dst); err != nil {
		return nil, err
	}

	for _, process := range dst {

		ch <- prometheus.MustNewConstMetric(
			c.Numbercurrentactivebackendwebsockets,
			prometheus.CounterValue,
			float64(process.Numbercurrentactivebackendwebsockets),
		)

		ch <- prometheus.MustNewConstMetric(
			c.NumbernewbackendwebsocketsPersec,
			prometheus.GaugeValue,
			float64(process.NumbernewbackendwebsocketsPersec),
		)

		ch <- prometheus.MustNewConstMetric(
			c.Numberrequests,
			prometheus.CounterValue,
			float64(process.Numberrequests),
		)

		ch <- prometheus.MustNewConstMetric(
			c.NumberrequestsPersec,
			prometheus.GaugeValue,
			float64(process.NumberrequestsPersec),
		)

		ch <- prometheus.MustNewConstMetric(
			c.Numberresponses,
			prometheus.CounterValue,
			float64(process.Numberresponses),
		)

		ch <- prometheus.MustNewConstMetric(
			c.NumberresponsesPersec,
			prometheus.GaugeValue,
			float64(process.NumberresponsesPersec),
		)

		ch <- prometheus.MustNewConstMetric(
			c.Numbertransactionscompletions,
			prometheus.CounterValue,
			float64(process.Numbertransactionscompletions),
		)

		ch <- prometheus.MustNewConstMetric(
			c.NumbertransactionscompletionsPersec,
			prometheus.GaugeValue,
			float64(process.NumbertransactionscompletionsPersec),
		)

		ch <- prometheus.MustNewConstMetric(
			c.Numbertransactionsfailed,
			prometheus.CounterValue,
			float64(process.Numbertransactionsfailed),
		)

		ch <- prometheus.MustNewConstMetric(
			c.NumbertransactionsfailedPersec,
			prometheus.GaugeValue,
			float64(process.NumbertransactionsfailedPersec),
		)
	}
	return nil, nil
}
