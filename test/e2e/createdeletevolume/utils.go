package createdeletevolume

import (
	"fmt"
	"os/exec"
)

func int32Ptr(i int32) *int32 { return &i }

func boolPter(i bool) *bool { return &i }

func runInLinux(cmd string) string {
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("ERROR:%+v \n", err)
	}
	return string(result)
}
