package message

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func F2(in float64) float64 {
	out, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", in), 64)
	return out
}
func GetCpuMem() (float64, float64) {
	percent, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	return percent[0], memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(false)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}
