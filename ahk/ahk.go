package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}
func ahkTest() {
	ahk := syscall.NewLazyDLL("AutoHotkey.dll")

	log.Println("NewLazyDll:", ahk.Name)

	//f := ahk.NewProc("ahkdll")
	ahktextdll := ahk.NewProc("ahkdll")
	f := ahk.NewProc("ahkReady")
	add := ahk.NewProc("addScript")

	ret, ret2, err := f.Call(0)
	log.Println("ret1:", ret, "   ret2:", ret2, "  last error:", err)
	ret, ret2, err = ahktextdll.Call(StrPtr("test中文.ahk"), 0, 0)
	log.Println("ret1:", ret, "   ret2:", ret2, "  last error:", err)
	ret, ret2, err = add.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Msgbox Hello World \n#z::Run www.qq.com"))), 1) // {{{// }}}
	log.Println("ret1:", ret, "   ret2:", ret2, "  last error:", err)

}

type ulong int32
type ulong_ptr uintptr

type PROCESSENTRY32 struct {
	dwSize              ulong
	cntUsage            ulong
	th32ProcessID       ulong
	th32DefaultHeapID   ulong_ptr
	th32ModuleID        ulong
	cntThreads          ulong
	th32ParentProcessID ulong
	pcPriClassBase      ulong
	dwFlags             ulong
	szExeFile           [260]byte
}

func DllTest() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
	pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
	if int(pHandle) == -1 {
		return
	}
	Process32Next := kernel32.NewProc("Process32Next")
	for {
		var proc PROCESSENTRY32
		proc.dwSize = ulong(unsafe.Sizeof(proc))
		if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
			fmt.Println("ProcessName : " + string(proc.szExeFile[0:]))
			fmt.Println("ProcessID : " + strconv.Itoa(int(proc.th32ProcessID)))
		} else {
			break
		}
	}
	//CloseHandle := kernel32.NewProc("CloseHandle")
	//_, _, _ = CloseHandle.Call(pHandle)
}

func main() {
	DllTest()
	ahkTest()
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		cmd := string(data)
		if cmd == "stop" {
			break
		}
	}
	fmt.Println("stop..")

}
