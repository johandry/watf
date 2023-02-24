package main

import (
	"log"
	"os"
	"strconv"

	"syscall/js"

	"github.com/johandry/terranova"
	"github.com/johandry/terranova/logger"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

var code string

const stateFilename = "simple.tfstate"

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

func apply(this js.Value, i []js.Value) interface{} {
	count := 0
	keyName := "demo"

	log := log.New(os.Stderr, "", log.LstdFlags)
	logMiddleware := logger.NewMiddleware()
	defer logMiddleware.Close()

	platform, err := terranova.NewPlatform(code).
		SetMiddleware(logMiddleware).
		AddProvider("aws", aws.Provider()).
		Var("c", count).
		Var("key_name", keyName).
		PersistStateToFile(stateFilename)

	if err != nil {
		log.Fatalf("Fail to create the platform using state file %s. %s", stateFilename, err)
	}

	terminate := (count == 0)
	if err := platform.Apply(terminate); err != nil {
		log.Fatalf("Fail to apply the changes to the platform. %s", err)
	}

	println("Completed")

	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("apply", js.FuncOf(apply))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}

func init() {
	code = `
  variable "c"    { default = 2 }
  variable "key_name" {}
  provider "aws" {
    region        = "us-west-2"
  }
  resource "aws_instance" "server" {
    instance_type = "t2.micro"
    ami           = "ami-6e1a0117"
    count         = "${var.c}"
    key_name      = "${var.key_name}"
  }
`
}
