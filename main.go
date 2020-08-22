package main

import (
	"strconv"
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	println("Value 1:" + value1)
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	println("Value 2:" + value2)

	int1, err := strconv.Atoi(value1)
	if err != nil {
		println("Error" + err.Error())
		return nil
	}
	int2, err := strconv.Atoi(value2)
	if err != nil {
		println("Error" + err.Error())
		return nil
	}

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)

	return nil
}

func subtract(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	println("Value 1:" + value1)
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	println("Value 2:" + value2)

	int1, err := strconv.Atoi(value1)
	if err != nil {
		println("Error" + err.Error())
		return nil
	}
	int2, err := strconv.Atoi(value2)
	if err != nil {
		println("Error" + err.Error())
		return nil
	}

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)

	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
