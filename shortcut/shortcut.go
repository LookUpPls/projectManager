package shortcut

import (
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
		log.Fatal(err)
	}

	wshell, err := shell.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		log.Fatal(err)
	}

	return &ShortcutCreator{
		wshell: wshell,
	}
}

func (sc *ShortcutCreator) CreateShortcut(shortcutPath, targetPath string) {
	cs, err := oleutil.CallMethod(sc.wshell, "CreateShortcut", shortcutPath)
	if err != nil {
		log.Fatal(err)
	}

	idispatch := cs.ToIDispatch()
	oleutil.PutProperty(idispatch, "TargetPath", targetPath)
	oleutil.PutProperty(idispatch, "WorkingDirectory", targetPath)
	oleutil.CallMethod(idispatch, "Save")
}

func (sc *ShortcutCreator) LoadShortcutTarget(shortcutPath string) string {
	cs, err := oleutil.CallMethod(sc.wshell, "CreateShortcut", shortcutPath)
	if err != nil {
		log.Fatal(err)
	}

	idispatch := cs.ToIDispatch()
	target, err := oleutil.GetProperty(idispatch, "TargetPath")
	if err != nil {
		log.Fatal(err)
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
