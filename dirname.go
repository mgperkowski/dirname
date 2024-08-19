package dirname

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetDirname() (string, error) {

	execPath, err := os.Executable()

	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return "", err
	}

	dirname := filepath.Dir(execPath)

	if strings.Contains(dirname, os.TempDir()) {
		dirname, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current working directory:", err)
			return "", err
		}
	}

	return dirname, nil
}
