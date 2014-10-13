package builder

import (
	"time"
)

/*
 Builder接口
*/

type Builder interface {
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
	Vacation Vacation
}

/*
 具体方法实现
*/

func (V *VacationBuilder) BuildDay(date time.Time) {
	V.Vacation.BuildDay = date
}

func (V *VacationBuilder) Days(days int8) {
	V.Vacation.Days = days
}

func (V *VacationBuilder) AddHotel(name string, date time.Time) {
	V.Vacation.Hotel, V.Vacation.StartDate = name, date
}

func (V *VacationBuilder) AddReservation(reserved bool) {
	V.Vacation.Reservation = reserved
}

/*
 生成产品
*/

func (V *VacationBuilder) GetVacation() *Vacation {
	return &V.Vacation
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
	self.Builder.AddHotel("香格里拉酒店", time.Date(2014, 10, 1, 18, 0, 0, 0, time.UTC))
	self.Builder.AddReservation(true)
	self.Builder.BuildDay(time.Date(2014, 9, 15, 12, 0, 0, 0, time.UTC))
	self.Builder.Days(7)
	return self.Builder.GetVacation()
}
