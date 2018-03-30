package hotpi

import (
	"fmt"
	"testing"
)

func TestCpuTemp(t *testing.T) {
	temp, err := CpuTemp()
	if err != nil {
		t.Error("CpuTemp: ", err)
	}
	fmt.Println("CPU temperature: ", temp, "°C")
}

func TestCpuFreq(t *testing.T) {
	freq, err := CpuFreq()
	if err != nil {
		t.Error("CpuFreq: ", err)
	}
	fmt.Println("Current CPU frequency: ", freq, "MHz")
}
