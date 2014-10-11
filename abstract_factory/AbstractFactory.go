package abstract_factory

import (
	"fmt"
)

/*
 Abstract Factory 抽象工厂 (一个 Interface)
*/

type GUIFactory interface {
	CreateButton() Button
}

/*
 Concrete Factory 具体工厂 (WinFactory,要实现 Interface 中的函数)
*/

type WinFactory struct {
	Os string
}

/*
 Concrete Factory 具体工厂 (OSXFactory,要实现 Interface 中的函数)
*/

type OSXFactory struct {
	Os string
}

/*
 Abstract Product 抽象产品 (一个 Interface)
*/

type Button interface {
	Paint()
}

/*
 Concrete Product 具体产品 (WinButton,要实现 Interface 中的函数)
*/

type WinButton struct {
}

/*
 Concrete Product 具体产品 (OSXButton,要实现 Interface 中的函数)
*/

type OSXButton struct {
}

/*
 具体函数实现
*/

func (*WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (*OSXFactory) CreateButton() Button {
	return &OSXButton{}
}

func (self *WinButton) Paint() {
	fmt.Println("这是一个Windows按钮")
}

func (self *OSXButton) Paint() {
	fmt.Println("这是一个OSX按钮")
}

/*
 Client (一个应用)
*/

type App struct {
}

func (*App) Verify() {
	f := OSXFactory{}
	b := f.CreateButton()
	b.Paint()
}
