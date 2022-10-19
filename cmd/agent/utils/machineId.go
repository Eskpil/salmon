package utils

import (
	"os"
	"strings"
)

func GetMachineId() (string, error) {
	file, err := os.Open("/etc/machine-id")

	if err != nil {
		return "", err
	}

	data := make([]byte, 32)
	_, err = file.Read(data)

	return strings.TrimSuffix(string(data), "\n"), err
}
