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

	f := ahk.NewProc("ahkExec")

	ret, ret2, err := f.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("MsgBox  run in go"))))
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
