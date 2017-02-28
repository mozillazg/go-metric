package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	isatty "github.com/mattn/go-isatty"
	metric "github.com/mozillazg/go-metric"
)

const version = "0.1.0"
const defaultAppDiskPath = "/root/executor"
const timeLayout = "2006-01-02 15:04:05"

func calculateCPUPercentUnix(previousCPU, previousSystem uint64, m *metric.Metric) float64 {
	var (
		cpuPercent = 0.0
		// calculate the change for the cpu usage of the container in between readings
		cpuDelta = float64(m.CPUStats.CPUUsage.TotalUsage) - float64(previousCPU)
		// calculate the change for the entire system between readings
		systemDelta = float64(m.CPUStats.SystemCPUUsage) - float64(previousSystem)
	)

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(m.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return cpuPercent
}

func bytesToGB(n uint64) float64 {
	return float64(n) / 1024.0 / 1024.0 / 1024.0
}

func formatOuput(mArray []metric.Metric, appDiskPath string) {
	firstM := mArray[0]
	memSize := firstM.MemoryStats.Limit
	appDiskSzie := firstM.DisksUsage[appDiskPath].Total
	cpuNumber := len(firstM.CPUStats.CPUUsage.PercpuUsage)

	fmt.Printf("#Time\tCPU(%%)(T: %d)\tRAM(%%)(T: %.1f GB)\tDisk(%%)(T: %.1f GB)\n",
		cpuNumber, bytesToGB(memSize), bytesToGB(appDiskSzie))
	for _, m := range mArray {
		t := metric.ParseTime(m.Read)
		memPercent := float64(m.MemoryStats.Usage) / float64(m.MemoryStats.Limit) * 100.0
		disk := m.DisksUsage[appDiskPath]
		diskPercent := float64(disk.Total-disk.Free) / float64(disk.Total) * 100.0
		previousCPU := m.PreCPUStats.CPUUsage.TotalUsage
		previousSystem := m.PreCPUStats.SystemCPUUsage
		cpuPercent := calculateCPUPercentUnix(previousCPU, previousSystem, &m)
		fmt.Printf("%s\t%.2f\t%.2f\t%.2f\t\n", t.Format(timeLayout),
			cpuPercent, memPercent, diskPercent)
	}
}

func main() {
	flag.Parse()
	metricStrArray := flag.Args()
	stdin := []byte{}
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		stdin, _ = ioutil.ReadAll(os.Stdin)
	}
	if len(stdin) > 0 {
		stdinStr := strings.TrimSpace(string(stdin))
		for _, line := range strings.Split(stdinStr, "\n") {
			metricStrArray = append(metricStrArray, line)
		}
	}

	if len(metricStrArray) == 0 {
		os.Exit(1)
	}

	mArray := []metric.Metric{}
	for _, line := range metricStrArray {
		m, err := metric.ParseJSON(line)
		if err != nil {
			fmt.Printf("Parse json failed!\nJSON:\n%#v\nError:\n%s\n",
				line, err)
			fmt.Println(m.MemoryStats.Stats)
		} else {
			mArray = append(mArray, m)
		}
	}

	if len(mArray) == 0 {
		os.Exit(1)
	}
	formatOuput(mArray, defaultAppDiskPath)
}
