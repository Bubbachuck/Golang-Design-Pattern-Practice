package builder

import (
	"time"
)

/*
 Builder接口
*/

type Builder interface {
	BuildNewVacation()
	BuildDay(date time.Time)
	Days(days int8)
	AddHotel(name string, date time.Time)
	AddReservation(reserved bool)
	GetVacation() *Vacation
}

/*
 Product,构建的产品
*/

type Vacation struct {
	BuildDay    time.Time
	Days        int8
	StartDate   time.Time
	Hotel       string
	Reservation bool
}

/*
 Concrete Builder, 具体建造者
*/

type VacationBuilder struct {
	Vacation *Vacation
}

/*
 具体方法实现
*/

func (self *VacationBuilder) BuildNewVacation() {
	self.Vacation = &Vacation{}
}

func (self *VacationBuilder) BuildDay(date time.Time) {
	self.Vacation.BuildDay = date
}

func (self *VacationBuilder) Days(days int8) {
	self.Vacation.Days = days
}

func (self *VacationBuilder) AddHotel(name string, date time.Time) {
	self.Vacation.Hotel, self.Vacation.StartDate = name, date
}

func (self *VacationBuilder) AddReservation(reserved bool) {
	self.Vacation.Reservation = reserved
}

/*
 生成产品
*/

func (self *VacationBuilder) GetVacation() *Vacation {
	return self.Vacation
}

/*
 Director,或者Client
*/

type Director struct {
	Builder Builder
}

/*
 Director或Client构建产品的实现
*/

func (self *Director) Construct() *Vacation {
	self.Builder.BuildNewVacation()
	self.Builder.AddHotel("香格里拉酒店", time.Date(2014, 10, 1, 18, 0, 0, 0, time.UTC))
	self.Builder.AddReservation(true)
	self.Builder.BuildDay(time.Date(2014, 9, 15, 12, 0, 0, 0, time.UTC))
	self.Builder.Days(7)
	return self.Builder.GetVacation()
}
