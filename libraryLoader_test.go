package libraryLoader

import (
    "fmt"
    "syscall"
    "unsafe"
    "testing"
)

func TestNetstat(t *testing.T) {
    var MB_YESNOCANCEL = 0x00000003
    ret, _, _ := Instance().Call("user32.dll", "MessageBoxW", 0,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("MessageBox"))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Test"))),
        uintptr(MB_YESNOCANCEL))
    fmt.Printf("Return: %d\n", ret)
}
