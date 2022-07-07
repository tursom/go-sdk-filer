//go:build js && wasm

package filer

import (
	"encoding/hex"
	"sync"
	"syscall/js"
)

var (
	Console = ConsoleUtil(js.Global().Get("console"))
	Filer   = js.Global().Get("window").Get("Filer")
	Fs      = Filer.Get("fs")
	Buffer  = Filer.Get("Buffer")
)

type (
	ConsoleUtil js.Value

	File struct {
		path string
	}
)

func (c ConsoleUtil) Log(values ...any) {
	js.Value(c).Call("log", values...)
}

func Open(path string) *File {
	return &File{
		path: path,
	}
}

func (f *File) Write(p []byte) (n int, err error) {
	var wait sync.WaitGroup
	wait.Add(1)
	Fs.Call("appendFile", f.path, NewBuffer(p), js.FuncOf(func(this js.Value, args []js.Value) any {
		for _, arg := range args {
			js.Global().Get("console").Call("log", arg)
			js.Global().Get("console").Call("log", arg.Type().String())

			//bytes, _ := json.Marshal(arg)
			//arg.Type()
			//fmt.Println(string(bytes))
		}
		wait.Done()
		return nil
	}))
	wait.Wait()
	return len(p), nil
}

func (f *File) Read(p []byte) (n int, err error) {
	f.readFile()
	return
}

func (f *File) readFile() (bytes []byte, err error) {
	var wait sync.WaitGroup
	wait.Add(1)

	Fs.Call("readFile", f.path, js.FuncOf(func(this js.Value, args []js.Value) any {
		defer wait.Done()

		if args[0].Type() != js.TypeNull {
			Console.Log("err", args[0])
			return nil
		}
		Console.Log("read file: ", args[1])
		bytes, err = hex.DecodeString(args[1].String())
		return nil
	}))

	wait.Wait()
	return
}
