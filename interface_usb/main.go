package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

type Cermra struct {
}

type Compter struct {
}

func (p Phone) Start() {
	fmt.Println("Phone is start working")
}

func (p Phone) Stop() {
	fmt.Println("Phone is stop working")
}

func (c Cermra) Start() {
	fmt.Println("Cremra is start working")
}

func (c Cermra) Stop() {
	fmt.Println("cermra is stop working")
}

func (c Compter) Working(usb Usb) {
	usb.Start()
	usb.Stop()

}

func main() {

	phone := Phone{}

	cermra := Cermra{}

	computer := Compter{}

	computer.Working(phone)
	computer.Working(cermra)

}
