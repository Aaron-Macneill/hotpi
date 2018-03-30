package hotpi

import (
	"testing"
	"fmt"
)

func TestCpuTemp(t *testing.T) {
	temp, err := CpuTemp()
	if err != nil {
		t.Error("CpuTemp, ", err)
	}
	fmt.Println(temp)
}
