package metrics

import (
	"log"

	lxd "github.com/lxc/lxd/client"
	lxdapi "github.com/lxc/lxd/shared/api"
	"github.com/prometheus/client_golang/prometheus"
)

// Collector collects metrics to be sent to Prometheus.
type collector struct {
	logger *log.Logger
	server lxd.InstanceServer
}

// NewCollector creates a new collector with logger and LXD connection.
func NewCollector(
	logger *log.Logger, server lxd.InstanceServer) prometheus.Collector {
	return &collector{logger: logger, server: server}
}

var (
	cpuUsageDesc = prometheus.NewDesc("lxd_instance_cpu_usage",
		"instanceCPU Usage in Seconds",
		[]string{"instance_name","instance_type"}, nil,
	)
	diskUsageDesc = prometheus.NewDesc("lxd_instance_disk_usage",
		"instanceDisk Usage",
		[]string{"instance_name", "disk_device","instance_type"}, nil,
	)
	networkUsageDesc = prometheus.NewDesc("lxd_instance_network_usage",
		"instanceNetwork Usage",
		[]string{"instance_name", "interface", "operation","instance_type"}, nil,
	)

	memUsageDesc = prometheus.NewDesc("lxd_instance_mem_usage",
		"instanceMemory Usage",
		[]string{"instance_name","instance_type"}, nil,
	)
	memUsagePeakDesc = prometheus.NewDesc("lxd_instance_mem_usage_peak",
		"instanceMemory Usage Peak",
		[]string{"instance_name","instance_type"}, nil,
	)
	swapUsageDesc = prometheus.NewDesc("lxd_instance_swap_usage",
		"instanceSwap Usage",
		[]string{"instance_name","instance_type"}, nil,
	)
	swapUsagePeakDesc = prometheus.NewDesc("lxd_instance_swap_usage_peak",
		"instanceSwap Usage Peak",
		[]string{"instance_name","instance_type"}, nil,
	)

	processCountDesc = prometheus.NewDesc("lxd_instance_process_count",
		"instancenumber of process Running",
		[]string{"instance_name","instance_type"}, nil,
	)
	InstancePIDDesc = prometheus.NewDesc("lxd_instance_pid",
		"instancePID",
		[]string{"instance_name","instance_type"}, nil,
	)
	runningStatusDesc = prometheus.NewDesc("lxd_instance_running_status",
		"instanceRunning Status",
		[]string{"instance_name","instance_type"}, nil,
	)
)

// Describe fills given channel with metrics descriptor.
func (collector *collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- cpuUsageDesc
	ch <- memUsageDesc
	ch <- memUsagePeakDesc
	ch <- swapUsageDesc
	ch <- swapUsagePeakDesc
	ch <- processCountDesc
	ch <- InstancePIDDesc
	ch <- runningStatusDesc
	ch <- diskUsageDesc
	ch <- networkUsageDesc
}

// Collect fills given channel with metrics data.
func (collector *collector) Collect(ch chan<- prometheus.Metric) {
	//log.Print(collector.server.GetInstances("virtual-machine"))
	InstanceTypes := []lxdapi.InstanceType{"container", "virtual-machine"}
	for _, InstanceType := range InstanceTypes {
		//log.Print(InstanceType)
		InstanceNames, err := collector.server.GetInstanceNames(InstanceType)
		if err != nil {
			collector.logger.Printf("Can't query instancenames: %s", err)
			return
		}

		for _, InstanceName := range InstanceNames {
			state, _, err := collector.server.GetInstanceState(InstanceName)
			if err != nil {
				collector.logger.Printf(
					"Can't query instancestate for `%s`: %s", InstanceName, err)
				continue
			}
			
			collector.collectInstanceMetrics(ch, InstanceName, state, string(InstanceType) )
		}
	}
}

func (collector *collector) collectInstanceMetrics(
	ch chan<- prometheus.Metric,
	InstanceName string,
	state *lxdapi.InstanceState,
	InstanceType string,
) {
	//log.Print(InstanceType, InstanceName)
	ch <- prometheus.MustNewConstMetric(cpuUsageDesc,
		prometheus.GaugeValue, float64(state.CPU.Usage), InstanceName, InstanceType)
	ch <- prometheus.MustNewConstMetric(processCountDesc,
		prometheus.GaugeValue, float64(state.Processes), InstanceName, InstanceType)
	ch <- prometheus.MustNewConstMetric(
		InstancePIDDesc, prometheus.GaugeValue, float64(state.Pid), InstanceName, InstanceType)

	ch <- prometheus.MustNewConstMetric(memUsageDesc,
		prometheus.GaugeValue, float64(state.Memory.Usage), InstanceName, InstanceType)
	ch <- prometheus.MustNewConstMetric(memUsagePeakDesc,
		prometheus.GaugeValue, float64(state.Memory.UsagePeak), InstanceName, InstanceType)
	ch <- prometheus.MustNewConstMetric(swapUsageDesc,
		prometheus.GaugeValue, float64(state.Memory.SwapUsage), InstanceName, InstanceType)
	ch <- prometheus.MustNewConstMetric(swapUsagePeakDesc,
		prometheus.GaugeValue, float64(state.Memory.SwapUsagePeak), InstanceName, InstanceType)

	runningStatus := 0
	if state.Status == "Running" {
		runningStatus = 1
	}
	ch <- prometheus.MustNewConstMetric(runningStatusDesc,
		prometheus.GaugeValue, float64(runningStatus), InstanceName, InstanceType)

	for diskName, state := range state.Disk {
		ch <- prometheus.MustNewConstMetric(diskUsageDesc,
			prometheus.GaugeValue, float64(state.Usage), InstanceName, diskName, InstanceType)
	}

	for ethName, state := range state.Network {
		networkMetrics := map[string]int64{
			"BytesReceived":   state.Counters.BytesReceived,
			"BytesSent":       state.Counters.BytesSent,
			"PacketsReceived": state.Counters.PacketsReceived,
			"PacketsSent":     state.Counters.PacketsSent,
		}

		for opName, value := range networkMetrics {
			ch <- prometheus.MustNewConstMetric(networkUsageDesc,
				prometheus.GaugeValue, float64(value), InstanceName, ethName, opName, InstanceType)
		}
	}
}
