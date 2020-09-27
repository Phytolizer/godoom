// +build windows

package system

func consoleStdout() bool {
	// SDL "helpfully" always redirects stdout to a file. thanks sdl love you <3
	return false
}
