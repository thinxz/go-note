// 调用windows的api函数
package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 加载 DLL
var (
	//kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	//getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")
	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND
	MB_DEFBUTTON1        = 0x00000000
	MB_DEFBUTTON2        = 0x00000100
	MB_DEFBUTTON3        = 0x00000200
	MB_DEFBUTTON4        = 0x00000300
)

func init() {
	fmt.Print("Starting Up\n")
}

// 错误处理
func abort(funcName string, err int) {
	panic(funcName + " failed: " + syscall.Errno(err).Error())
}

// 消息窗口
func MessageBox(caption, text string, style uintptr) int {
	// var hwnd HWND
	ret, _, callErr := syscall.Syscall6(
		//
		uintptr(messageBox), 4,
		// HWND 句柄
		0,
		// Text
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		// Caption
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		// type
		style, 0, 0)
	if callErr != 0 {
		abort("Call MessageBox", int(callErr))
	}

	return int(ret)
}

func main() {
	//defer syscall.FreeLibrary(kernel32)
	defer syscall.FreeLibrary(user32)

	fmt.Printf("Retern: %d\n", MessageBox("Done Title", "This test is Done.", MB_OKCANCEL))
}
