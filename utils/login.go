package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetLastLogin() string {
	cmd := exec.Command("last", "-n", "1")
	output, _ := cmd.Output()

	fields := strings.Fields(string(output))

	return fmt.Sprintf("%s(%s)", fields[2], fields[0])
}
