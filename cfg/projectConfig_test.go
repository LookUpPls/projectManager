package cfg

import (
	"fmt"
	"os"
	"testing"
)

func TestTagCreate(t *testing.T) {
	fmt.Println("run here")
	fmt.Println("run here1")
	if _, err := os.Stat("C:\\WorkSpace1\\java\\english.lnk"); os.IsNotExist(err) {
		t.Errorf("The java tag was not created")
	}
	if _, err := os.Stat("C:\\WorkSpace1\\my\\english.lnk"); os.IsNotExist(err) {
		t.Errorf("The my tag was not created")
	}
}
