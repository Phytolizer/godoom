package system

type atExitFunc = func()

type atExitListEntry struct {
	Func       atExitFunc
	RunOnError bool
	Next       *atExitListEntry
}

var ExitFuncs *atExitListEntry = nil

func AtExit(fn atExitFunc, runOnError bool) {
	entry := new(atExitListEntry)
	entry.Func = fn
	entry.RunOnError = runOnError
	entry.Next = ExitFuncs
	ExitFuncs = entry
}
