//go:build windows && amd64

package implementations

/*
#include <stdlib.h>

#include "../event/callback.h"
*/
import "C"

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/brunoga/unitybridge/unity/event"

	internal_event "github.com/brunoga/unitybridge/internal/event"
)

var (
	libPath = "./lib/windows/amd64/unitybridge.dll"

	// Singleton.
	UnityBridgeImpl *loadLibraryUnityBridgeImpl = &loadLibraryUnityBridgeImpl{}
)

func init() {
	var err error

	UnityBridgeImpl.handle, err = syscall.LoadDLL(libPath)
	if err != nil {
		panic(fmt.Sprintf("Could not load Unity Bridge library at \"%s\": %s",
			libPath, err))
	}

	UnityBridgeImpl.createUnityBridge =
		UnityBridgeImpl.getSymbol("CreateUnityBridge")
	UnityBridgeImpl.destroyUnityBridge =
		UnityBridgeImpl.getSymbol("DestroyUnityBridge")
	UnityBridgeImpl.unityBridgeInitialize =
		UnityBridgeImpl.getSymbol("UnityBridgeInitialize")
	UnityBridgeImpl.unityBridgeUninitialize =
		UnityBridgeImpl.getSymbol("UnityBridgeUninitialze") // Typo in library.
	UnityBridgeImpl.unitySendEvent =
		UnityBridgeImpl.getSymbol("UnitySendEvent")
	UnityBridgeImpl.unitySendEventWithString =
		UnityBridgeImpl.getSymbol("UnitySendEventWithString")
	UnityBridgeImpl.unitySendEventWithNumber =
		UnityBridgeImpl.getSymbol("UnitySendEventWithNumber")
	UnityBridgeImpl.unitySetEventCallback =
		UnityBridgeImpl.getSymbol("UnitySetEventCallback")
	UnityBridgeImpl.UnityGetSecurityKeyByKeyChainIndex =
		UnityBridgeImpl.getSymbol("UnityGetSecurityKeyByKeyChainIndex")
}

type loadLibraryUnityBridgeImpl struct {
	handle *syscall.DLL

	createUnityBridge                  *syscall.Proc
	destroyUnityBridge                 *syscall.Proc
	unityBridgeInitialize              *syscall.Proc
	unityBridgeUninitialize            *syscall.Proc
	unitySendEvent                     *syscall.Proc
	unitySendEventWithString           *syscall.Proc
	unitySendEventWithNumber           *syscall.Proc
	unitySetEventCallback              *syscall.Proc
	UnityGetSecurityKeyByKeyChainIndex *syscall.Proc
}

func (u *loadLibraryUnityBridgeImpl) Create(name string, debuggable bool,
	logPath string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	intDebuggable := 0
	if debuggable {
		intDebuggable = 1
	}

	cLogPath := C.CString(logPath)
	defer C.free(unsafe.Pointer(cLogPath))

	_, _, _ = u.createUnityBridge.Call(
		uintptr(unsafe.Pointer(cName)),
		uintptr(intDebuggable),
		uintptr(unsafe.Pointer(cLogPath)),
	)
}

func (u *loadLibraryUnityBridgeImpl) Initialize() bool {
	ret, _, _ := u.unityBridgeInitialize.Call()
	return ret == 1
}

func (u *loadLibraryUnityBridgeImpl) SetEventCallback(t event.Type,
	callback event.Callback) {
	var eventCallbackUintptr uintptr
	if callback != nil {
		eventCallbackUintptr = uintptr(C.eventCallbackC)
	}

	eventCode := event.NewFromType(t).Code()

	_, _, _ = u.unitySetEventCallback.Call(
		uintptr(eventCode),
		eventCallbackUintptr,
	)

	internal_event.SetEventCallback(t, callback)
}

func (u *loadLibraryUnityBridgeImpl) SendEvent(e *event.Event, data uintptr,
	tag uint64) {
	_, _, _ = u.unitySendEvent.Call(
		uintptr(e.Code()),
		data,
		uintptr(tag),
	)

}

func (u *loadLibraryUnityBridgeImpl) SendEventWithString(e *event.Event,
	data string, tag uint64) {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))

	_, _, _ = u.unitySendEventWithString.Call(
		uintptr(e.Code()),
		uintptr(unsafe.Pointer(cData)),
		uintptr(tag),
	)
}

func (u *loadLibraryUnityBridgeImpl) SendEventWithNumber(e *event.Event, data,
	tag uint64) {
	_, _, _ = u.unitySendEventWithNumber.Call(
		uintptr(e.Code()),
		uintptr(data),
		uintptr(tag),
	)
}

func (u *loadLibraryUnityBridgeImpl) GetSecurityKeyByKeyChainIndex(
	index int) string {
	cKeyUintptr, _, _ := u.UnityGetSecurityKeyByKeyChainIndex.Call(
		uintptr(index),
	)
	defer C.free(unsafe.Pointer(cKeyUintptr))

	return C.GoString((*C.char)(unsafe.Pointer(cKeyUintptr)))
}

func (u *loadLibraryUnityBridgeImpl) Uninitialize() {
	_, _, _ = u.unityBridgeUninitialize.Call()
}

func (u *loadLibraryUnityBridgeImpl) Destroy() {
	_, _, _ = u.destroyUnityBridge.Call()
}

func (h *loadLibraryUnityBridgeImpl) getSymbol(name string) *syscall.Proc {
	symbol, err := h.handle.FindProc(name)
	if err != nil {
		panic(fmt.Sprintf("Could not load symbol \"%s\": %s", name, err))
	}

	return symbol
}
