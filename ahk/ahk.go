package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

var (
//ahk, _ = syscall.LoadLibrary("AutoHotkey.dll")
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}
func ahkTest() {
	ahk := syscall.NewLazyDLL("AutoHotkey.dll")
	log.Println("call dll", ahk)

	//f := ahk.NewProc("ahkdll")
	ahktextdll := ahk.NewProc("ahktextdll")
	f := ahk.NewProc("ahkReady")
	add := ahk.NewProc("ahkExec")

	ret, ret2, err := ahktextdll.Call(IntPtr(0))
	log.Println("ret1:", ret, "   ret2:", ret2, "  last error:", err)
	ret, ret2, err = f.Call()
	log.Println("ret1:", ret, "   ret2:", ret2, "  last error:", err)
	ret, ret2, err = add.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Msgbox Hello World \n#z::Run www.qq.com"))))
	log.Println("ret1:", ret, "   ret2:", ret2, "  last error:", err)

}

//func ahkTest2() {
//log.Println("ahk dll:", ahk)
//ahktextdll, _ := syscall.GetProcAddress(ahk, "ahktextdll")
//log.Println("ahktextdll proc:", ahktextdll)
//ret1, _, err := syscall.Syscall(ahktextdll, 0, 0, 0, 0)
//log.Println("ret:", ret1, "  err:", err)

//addScript, _ := syscall.GetProcAddress(ahk, "addScript")
//log.Println("addScript proc:", addScript)
//ret2, _, err2 := syscall.Syscall(addScript, StrPtr("Msgbox Hello World \n#z::Run www.qq.com"), 1, 0, 0)
//log.Println("ret:", ret2, "err:", err2)

//}
func main() {
	ahkTest()

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		cmd := string(data)
		if cmd == "stop" {
			break
		}
	}
	//defer syscall.FreeLibrary(ahk)
	fmt.Println("stop..")

}
