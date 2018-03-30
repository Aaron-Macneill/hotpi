package hotpi

import (
	"io/ioutil"
	"strconv"
)

var (
	cpuTempDir     = "/sys/class/thermal/thermal_zone0"
	cpuTempFile    = cpuTempDir + "/temp"
	cpuFreqFile = "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq"
)

func parseFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data[:len(data)-2], nil
}

func CpuTemp() (float64, error) {
	tempBytes, err := parseFile(cpuTempFile)
	if err != nil {
		return 0, err
	}
	tempFloat, err := strconv.ParseFloat(string(tempBytes), 64)
	if err != nil {
		return 0, err
	}
	return tempFloat / 100, nil
}

func CpuFreq() (float64, error) {
	freqBytes, err := parseFile(cpuFreqFile)
	if err != nil {
		return 0, err
	}
	freqFloat, err := strconv.ParseFloat(string(freqBytes), 64)
	if err != nil {
		return 0, err
	}
	return freqFloat / 100, nil
}
