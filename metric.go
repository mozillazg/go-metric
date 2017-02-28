package metric

import (
	"encoding/json"
	"strings"
	"time"
)

const readTimeLayout = "2006-01-02T15:04:05.999999999Z07:00"

// Metric type
type Metric struct {
	Read        string                  `json:"read"`
	Networks    map[string]NetworkValue `json:"networks"`
	DisksUsage  map[string]DiskValue    `json:"disks_usage"`
	BlkioStats  BlkioStats              `json:"blkio_stats"`
	PreCPUStats CPUStats                `json:"precpu_stats"`
	MemoryStats MemoryStats             `json:"memory_stats"`
	CPUStats    CPUStats                `json:"cpu_stats"`
}

// NetworkValue type for Metric.Networks
type NetworkValue struct {
	TxDropped uint64 `json:"tx_dropped"`
	RxPackets uint64 `json:"rx_packets"`
	RxBytes   uint64 `json:"rx_bytes"`
	TxErrors  uint64 `json:"tx_errors"`
	RxErrors  uint64 `json:"rx_errors"`
	TxBytes   uint64 `json:"tx_bytes"`
	RxDropped uint64 `json:"rx_dropped"`
	TxPackets uint64 `json:"tx_packets"`
}

// DiskValue type for Metric.DisksUsage
type DiskValue struct {
	Total         uint64 `json:"total"`
	ReservedSpace uint64 `json:"reserved_space"`
	Free          uint64 `json:"free"`
}

// BlkioStats type for Metric.BlkioStats
type BlkioStats struct {
	IoServiceTimeRecursive  []BlkioRecursiveItem `json:"io_service_time_recursive"`
	SectorsRecursive        []BlkioRecursiveItem `json:"sectors_recursive"`
	IoServiceBytesRecursive []BlkioRecursiveItem `json:"io_service_bytes_recursive"`
	IoServicedRecursive     []BlkioRecursiveItem `json:"io_serviced_recursive"`
	IoTimeRecursive         []BlkioRecursiveItem `json:"io_time_recursive"`
	IoQueueRecursive        []BlkioRecursiveItem `json:"io_queue_recursive"`
	IoMergedRecursive       []BlkioRecursiveItem `json:"io_merged_recursive"`
	IoWaitTimeRecursive     []BlkioRecursiveItem `json:"io_wait_time_recursive"`
}

// BlkioRecursiveItem type for BlkioStats.XX
type BlkioRecursiveItem struct {
	Major uint64 `json:"major"`
	Value uint64 `json:"value"`
	Minor uint64 `json:"minor"`
	Op    string `json:"op"`
}

// CPUStats type for Metric.CPUStats
type CPUStats struct {
	CPUUsage       CPUUsage       `json:"cpu_usage"`
	SystemCPUUsage uint64         `json:"system_cpu_usage"`
	ThrottlingData ThrottlingData `json:"throttling_data"`
}

// CPUUsage type for CPUStats.CPUUsage
type CPUUsage struct {
	UsageInUsermode   uint64 `json:"usage_in_usermode"`
	TotalUsage        uint64 `json:"total_usage"`
	PercpuUsage       []int  `json:"percpu_usage"`
	UsageInKernelmode uint64 `json:"usage_in_kernelmode"`
}

// ThrottlingData type for CPUStats.ThrottlingData
type ThrottlingData struct {
	ThrottledTime    uint64 `json:"throttled_time"`
	Periods          uint64 `json:"periods"`
	ThrottledPeriods uint64 `json:"throttled_periods"`
}

// MemoryStats type for Metric.MemoryStats
type MemoryStats struct {
	Usage    uint64     `json:"usage"`
	Limit    uint64     `json:"limit"`
	Failcnt  uint64     `json:"failcnt"`
	Stats    MemoryStat `json:"stats"`
	MaxUsage uint64     `json:"max_usage"`
}

// MemoryStat type for MemoryStats.Stats
type MemoryStat struct {
	Unevictable             uint64 `json:"unevictable"`
	TotalInactiveFile       uint64 `json:"total_inactive_file"`
	TotalRssHuge            uint64 `json:"total_rss_huge"`
	Writeback               uint64 `json:"writeback"`
	TotalCache              uint64 `json:"total_cache"`
	TotalMappedFile         uint64 `json:"total_mapped_file"`
	MappedFile              uint64 `json:"mapped_file"`
	Pgfault                 uint64 `json:"pgfault"`
	TotalWriteback          uint64 `json:"total_writeback"`
	HierarchicalMemoryLimit uint64 `json:"hierarchical_memory_limit"`
	TotalActiveFile         uint64 `json:"total_active_file"`
	RssHuge                 uint64 `json:"rss_huge"`
	Cache                   uint64 `json:"cache"`
	ActiveAnon              uint64 `json:"active_anon"`
	Pgmajfault              uint64 `json:"pgmajfault"`
	TotalPgpgout            uint64 `json:"total_pgpgout"`
	Pgpgout                 uint64 `json:"pgpgout"`
	TotalActiveAnon         uint64 `json:"total_active_anon"`
	TotalUnevictable        uint64 `json:"total_unevictable"`
	TotalPgfault            uint64 `json:"total_pgfault"`
	TotalPgmajfault         uint64 `json:"total_pgmajfault"`
	TotalInactiveAnon       uint64 `json:"total_inactive_anon"`
	InactiveFile            uint64 `json:"inactive_file"`
	Pgpgin                  uint64 `json:"pgpgin"`
	TotalPgpgin             uint64 `json:"total_pgpgin"`
	Rss                     uint64 `json:"rss"`
	ActiveFile              uint64 `json:"active_file"`
	InactiveAnon            uint64 `json:"inactive_anon"`
	TotalRss                uint64 `json:"total_rss"`
}

// Version return version
func Version() string {
	return "0.1.0"
}

// ParseTime parse time string
func ParseTime(value string) time.Time {
	t, _ := time.Parse(readTimeLayout, value)
	return t
}

// ParseJSON parse JSON string to Metric
func ParseJSON(s string) (m Metric, err error) {
	s = strings.TrimSpace(s)
	err = json.Unmarshal([]byte(s), &m)
	return
}
