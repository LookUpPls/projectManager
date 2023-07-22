package cfg

import (
	"fmt"
	"testing"
)

func TestLoadConfigPath(t *testing.T) {
	fmt.Println(loadConfigPath())
}

func TestLoadHomePath(t *testing.T) {
	fmt.Println(loadHomePath())
}

func TestLoadExePath(t *testing.T) {
	fmt.Println(loadExecutablePath())
}
