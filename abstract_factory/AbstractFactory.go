package abstract_factory

import (
	"fmt"
)

/*
 定义第一类抽象产品接口
*/

type IBusinessCar interface {
	Run()
}

/*
 定义第二类抽象产品接口
*/

type ISportCar interface {
	Run()
}

/*
 定义第一家工厂BMW的第一类产品BMWBusinessCar
*/

type BMWBusinessCar struct {
	Model string
}

/*
 实现第一家工厂BMW的第一类产品BMWBusinessCar的属性方法
*/

func (self *BMWBusinessCar) Run() {
	fmt.Println("BMWBusinessCar " + self.Model + " run!")
}

/*
 定义第二家工厂Benz的第一类产品BenzBusinessCar
*/

type BenzBusinessCar struct {
	Model string
}

/*
 实现第二家工厂Benz的第一类产品BenzBusinessCar的属性方法
*/

func (self *BenzBusinessCar) Run() {
	fmt.Println("BenzBusinessCar " + self.Model + " run!")
}

/*
 定义第一家工厂BMW的第二类产品BMWSportCar
*/

type BMWSportCar struct {
	Model string
}

/*
 实现第一家工厂BMW的第二类产品BMWSportCar的属性方法
*/

func (self *BMWSportCar) Run() {
	fmt.Println("BMWSportCar " + self.Model + " run!")
}

/*
 定义第二家工厂Benz的第二类产品BMWSportCar
*/

type BenzSportCar struct {
	Model string
}

/*
 实现第二家工厂Benz的第二类产品BMWSportCar的属性方法
*/

func (self *BenzSportCar) Run() {
	fmt.Println("BenzSportCar " + self.Model + " run!")
}

/*
 定义抽象工厂接口
*/

type IDriver interface {
	BusinessCarDriver() IBusinessCar
	SportCarDriver() ISportCar
}

/*
 具体工厂 —— 第一家工厂BMWCarDriver
*/

type BMWCarDriver struct {
	Gender string
}

/*
 实现第一家工厂BMWCarDriver的方法
*/

func (self *BMWCarDriver) BusinessCarDriver(Model string) IBusinessCar {
	return &BMWBusinessCar{Model: Model}
}

func (self *BMWCarDriver) SportCarDriver(Model string) ISportCar {
	return &BMWSportCar{Model: Model}
}

/*
 具体工厂 —— 第二家工厂BenzCarDriver
*/

type BenzCarDriver struct {
	Gender string
}

/*
 实现第二家工厂BenzCarDriver的方法
*/

func (self *BenzCarDriver) BusinessCarDriver(Model string) IBusinessCar {
	return &BenzBusinessCar{Model: Model}
}

func (self *BenzCarDriver) SportCarDriver(Model string) ISportCar {
	return &BenzSportCar{Model: Model}
}

/*
 App (一个应用)
*/

type App struct {
}

func (*App) Verify() {
	//创建一个具体工厂
	MyDriver := BenzCarDriver{Gender: "male"}
	//工厂生产一类产品
	MyCar1 := MyDriver.SportCarDriver("SLR")
	//产品验收
	MyCar1.Run()
	MyCar2 := MyDriver.BusinessCarDriver("350L")
	MyCar2.Run()
}
