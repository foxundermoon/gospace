package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

func ahkTest() {
	ahk := syscall.NewLazyDLL("AutoHotkey.dll")

	log.Println("call dll", ahk.Name)

	//f := ahk.NewProc("ahkdll")
	ahktextdll := ahk.NewProc("ahktextdll")
	f := ahk.NewProc("ahkReady")
	add := ahk.NewProc("addScript")

	ret, ret2, err := ahktextdll.Call(0)
	log.Println("ret1:", ret, "   ret2:", ret2, "  last error:", err)
	ret, ret2, err = f.Call(0)
	log.Println("ret1:", ret, "   ret2:", ret2, "  last error:", err)
	ret, ret2, err = add.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Msgbox Hello World \n#z::Run www.qq.com"))), 1) // {{{// }}}
	log.Println("ret1:", ret, "   ret2:", ret2, "  last error:", err)

}

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
	fmt.Println("stop..")

}
