package hotpi
import (
	"io/ioutil"
	"strconv"
)


var (
	cpuTempDir = "/sys/class/thermal/thermal_zone0"
	cpuTempFile = cpuTempDir + "/temp"
	cpuCurrentFreq = "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq"
)

func CpuTemp() (float64, error) {
	temp, err := ioutil.ReadFile(cpuTempFile)
	if err != nil {
		return 0, err
	}
	temp = temp[:len(temp)-2]
	tempFloat, err := strconv.ParseFloat(string(temp), 64)
	if err != nil {
		return 0, err
	}
	return tempFloat / 100, nil
}

func CpuFreq() (float64, error) {
	freq, err := ioutil.ReadFile(cpuCurrentFreq)
	if err != nil {
		return 0, err
	}
	freq = freq[:len(freq)-2]
	freqFloat , err := strconv.ParseFloat(string(freq), 64)
	if err != nil {
		return 0, err
	}
	return freqFloat / 100, nil
}
