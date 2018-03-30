package hotpi
import (
	"io/ioutil"
	"strconv"
)


var (
	cpuTempDir = "/sys/class/thermal/thermal_zone0"
	cpuTempFile = cpuTempDir + "/temp"
)

func CpuTemp() (float64, error) {
	temp, err := ioutil.ReadFile(cpuTempFile)
	if err != nil {
		return 0, err
	}
	tempFloat, err := strconv.ParseFloat(string(temp), 64)
	if err != nil {
		return 0, err
	}
	return tempFloat, nil
}
