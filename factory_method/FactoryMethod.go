package factory_method

import (
	"fmt"
)

/*
 产品接口定义，产品需要支持Run函数
*/

type ICar interface {
	Run()
}

/*
 产品BMW Car
*/

type BMWCar struct {
	Model string
}

/*
 产品BMW Car实现Run函数
*/

func (self *BMWCar) Run() {
	fmt.Println("BMWCar " + self.Model + " run!")
}

/*
 产品Benz Car
*/

type BenzCar struct {
	Model string
}

/*
 产品Benz Car实现Run函数
*/

func (self *BenzCar) Run() {
	fmt.Println("BenzCar " + self.Model + " run!")
}

/*
 定义工厂接口，工厂方法为DriverCar
*/

type IDriver interface {
	DriverCar() ICar
}

/*
 BMW的司机，即为工厂
*/

type BMWCarDriver struct {
	Gender string
}

/*
 BMW的司机(工厂)实现工厂方法DriverCar
*/

func (self *BMWCarDriver) DriverCar(Model string) ICar {
	return &BMWCar{Model: Model}
}

/*
 Benz的司机，即为工厂
*/

type BenzCarDriver struct {
	Gender string
}

/*
 Benz的司机(工厂)实现工厂方法DriverCar
*/

func (self *BenzCarDriver) DriverCar(Model string) ICar {
	return &BenzCar{Model: Model}
}

type Client struct {
}

func (*Client) Verify() {
	//创建工厂实例
	MyDriver := BenzCarDriver{Gender: "Male"}
	//通过工厂的工厂方法创建产品
	MyCar := MyDriver.DriverCar("SLR")
	//验证产品属性
	MyCar.Run()
}
