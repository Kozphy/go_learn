package dsp

import "fmt"

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

type Adaptee struct {
	adapterType int
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func Execute_Adapter() {
	var processer IProcess = Adapter{}
	processer.process()
}
