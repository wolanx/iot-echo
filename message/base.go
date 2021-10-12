package message

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/zx5435/iot-echo/config"
)

var params *config.Params

func init() {
	params = config.NewParams()
}

func GetMetric() string {
	//arr := make(map[string]interface{})
	arr := params.LoadData()

	cpuPct, memPct := getCpuMem()
	arr["sn"] = config.GetConfig().Device.DeviceName
	arr["cpu"] = f2(cpuPct)
	arr["mem"] = f2(memPct)
	ret, _ := json.Marshal(arr)
	return string(ret)
}

func getCpuMem() (float64, float64) {
	percent, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	return percent[0], memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(false)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

func f2(in float64) float64 {
	out, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", in), 64)
	return out
}
