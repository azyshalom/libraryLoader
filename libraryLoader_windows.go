// +build windows

package libraryLoader

import (
    "sync"
    "syscall"
)

type Function struct {
    name string
    proc *syscall.Proc
}

type Library struct {
    name string
    lib  *syscall.DLL
    functions map[string]*Function
}

type LibraryLoader struct {
    libraries map[string]*Library
}

var initCtx sync.Once
var libraryLoaderCtx *LibraryLoader

func Instance() *LibraryLoader {
    initCtx.Do(func() {
        libraryLoaderCtx = &LibraryLoader {
            libraries: make(map[string]*Library),
        }
    })
    return libraryLoaderCtx
}

func (ctx *LibraryLoader) LoadLibrary(Name string) (*Library, error) {
    if _, ok := ctx.libraries[Name]; ok  {
        return ctx.libraries[Name], nil
    }
    lib, err := syscall.LoadDLL(Name)
    if err != nil {
        return nil, err
    }
    library := &Library {
        name: Name,
        lib: lib,
        functions: make(map[string]*Function),
    }
    ctx.libraries[Name] = library
    return library, nil
}

func (ctx *Library) GetProc(Name string) (*Function, error) {
    if _, ok := ctx.functions[Name]; ok  {
        return ctx.functions[Name], nil
    }
    proc, err := ctx.lib.FindProc(Name)
    if err != nil {
        return nil, err
    }
    function := &Function {
        name: Name,
        proc: proc,
    }
    ctx.functions[Name] = function
    return function, nil
}

func (ctx *Function) Call(a ...uintptr)  (r1, r2 uintptr, lastErr error) {
    return ctx.proc.Call(a...)
}

func (ctx *LibraryLoader) Call(library, function string, a ...uintptr) (r1, r2 uintptr, lastErr error) {
    l, err := ctx.LoadLibrary(library)
    if err != nil {
        return r1, r2, err
    }
    f, err := l.GetProc(function)
    if err != nil {
        return r1, r2, err
    }
    return f.Call(a...)
}


