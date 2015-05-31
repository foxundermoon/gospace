package ahk

import (
	"errors"
	"syscall"
)

type AhkClient struct {
	Loaded     bool
	LastErr    string
	AhkLazyDll *syscall.LazyDLL
	AhkProcs   *ahkProcs
}
type ahkProcs struct {
	ahkReady        *syscall.LazyProc
	ahktextdll      *syscall.LazyProc
	addScript       *syscall.LazyProc
	ahkExec         *syscall.LazyProc
	ahkLabel        *syscall.LazyProc
	ahkFunction     *syscall.LazyProc
	ahkPostFunction *syscall.LazyProc
	ahkassign       *syscall.LazyProc
	ahkgetvar       *syscall.LazyProc
	ahkTerminate    *syscall.LazyProc
	ahkReload       *syscall.LazyProc
	ahkPause        *syscall.LazyProc
}

func (ahk *AhkClient) Load() error {
	ahk.AhkLazyDll = syscall.NewLazyDLL("AutoHotkey.dll")
	ahk.AhkProcs.ahkReady = ahk.AhkLazyDll.NewProc("ahkReady")
	ahk.AhkProcs.ahktextdll = ahk.AhkLazyDll.NewProc("ahktextdll")
	ahk.AhkProcs.addScript = ahk.AhkLazyDll.NewProc("addScript")
	ahk.AhkProcs.ahkExec = ahk.AhkLazyDll.NewProc("ahkExec")
	ahk.AhkProcs.ahkLabel = ahk.AhkLazyDll.NewProc("ahkLabel")
	ahk.AhkProcs.ahkFunction = ahk.AhkLazyDll.NewProc("ahkFunction")
	ahk.AhkProcs.ahkPostFunction = ahk.AhkLazyDll.NewProc("ahkPostFunction")
	ahk.AhkProcs.ahkassign = ahk.AhkLazyDll.NewProc("ahkassign")
	ahk.AhkProcs.ahkgetvar = ahk.AhkLazyDll.NewProc("ahkgetvar")
	ahk.AhkProcs.ahkTerminate = ahk.AhkLazyDll.NewProc("ahkTerminate")
	ahk.AhkProcs.ahkReload = ahk.AhkLazyDll.NewProc("ahkReload")
	ahk.AhkProcs.ahkPause = ahk.AhkLazyDll.NewProc("ahkPause")
	ret1, _, lastErr := ahk.AhkProcs.ahktextdll.Call(0)
	ahk.LastErr = lastErr
	if ret1 < 1 {
		return errors.New(lastErr)
	}
	return nil
}

func (ahk *AhkClient) AhkReady() (bool, error) {
	ret1, _, lastErr := ahk.AhkProcs.ahkReady.Call(0)
	ahk.LastErr = LastErr
	if ret1 == 0 {
		return false, nil
	}
	if ret1 == 1 {
		return true, nil
	}
	return false, errors.New("return param error :", lastErr)
}
