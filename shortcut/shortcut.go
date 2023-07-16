package shortcut

import (
	"fmt"
	"log"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type ShortcutCreator struct {
	wshell *ole.IDispatch
}

func NewShortcutCreator() *ShortcutCreator {
	ole.CoInitialize(0)
	shell, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		fmt.Println(err)
	}

	wshell, err := shell.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		fmt.Println(err)
	}

	return &ShortcutCreator{
		wshell: wshell,
	}
}

func (sc *ShortcutCreator) CreateShortcut(shortcutPath, targetPath string) {
	fmt.Printf("开始创建快捷方式... %s -> %s\n", shortcutPath, targetPath)
	cs, err := oleutil.CallMethod(sc.wshell, "CreateShortcut", shortcutPath)
	if err != nil {
		fmt.Println(err)
	}

	idispatch := cs.ToIDispatch()
	property, err := oleutil.PutProperty(idispatch, "TargetPath", targetPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(property)
	putProperty, err := oleutil.PutProperty(idispatch, "WorkingDirectory", targetPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(putProperty)
	callMethod, err := oleutil.CallMethod(idispatch, "Save")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(callMethod)

}

func (sc *ShortcutCreator) LoadShortcutTarget(shortcutPath string) string {
	cs, err := oleutil.CallMethod(sc.wshell, "CreateShortcut", shortcutPath)
	if err != nil {
		fmt.Println(err)
	}

	idispatch := cs.ToIDispatch()
	target, err := oleutil.GetProperty(idispatch, "TargetPath")
	if err != nil {
		fmt.Println(err)
	}

	return target.ToString()
}

func (sc *ShortcutCreator) Close() {
	sc.wshell.Release()
	ole.CoUninitialize()
}

func main() {
	sc := NewShortcutCreator()
	defer sc.Close()

	targetPath := "C:\\WorkSpace\\_home\\your_repository"
	shortcutPath := "C:\\WorkSpace\\your_repository.lnk"
	sc.CreateShortcut(shortcutPath, targetPath)

	// Load shortcut target
	target := sc.LoadShortcutTarget(shortcutPath)
	log.Println("Loaded shortcut target:", target)

	// Call sc.CreateShortcut as many times as you need...
}
