package message

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetMetric() string {
	cpuPct, memPct := getCpuMem()
	ret := make(map[string]float64)
	ret["cpu"] = float64(cpuPct)
	ret["mem"] = float64(memPct)
	ret["rpm"] = 123
	marshal, _ := json.Marshal(ret)
	return string(marshal)
}

func getCpuMem() (int, int) {
	percent, _ := cpu.Percent(time.Second, false)
	a := int(percent[0])
	memInfo, _ := mem.VirtualMemory()
	b := int(memInfo.UsedPercent)
	return a, b
}
