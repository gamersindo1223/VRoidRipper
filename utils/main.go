package utils

// import (
// 	"os"
// 	"os/exec"
// 	"syscall"
// 	"unsafe"
// 	// "golang.org/x/sys/windows"
// )

// func SetConsoleTitle(title string) (int, error) {
// 	handle, err := windows.LoadLibrary("Kernel32.dll")
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer windows.FreeLibrary(handle)

// 	proc, err := windows.GetProcAddress(handle, "SetConsoleTitleW")
// 	if err != nil {
// 		return 0, err
// 	}

// 	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
// 	return int(r), err
// }

// func Clr() {
// 	cmd := exec.Command("cmd", "/c", "cls")
// 	cmd.Stdout = os.Stdout
// 	_ = cmd.Run()
// }
