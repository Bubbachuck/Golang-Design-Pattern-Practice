package prototype

import (
	"fmt"
)

/*
 Prototype接口定义
*/

type Image interface {
	Clone() Image
	GetImageName()
	SetImageName(Name string)
}

/*
 具体Prototype
*/

type ConcrectImage struct {
	Name string
}

/*
 具体Prototype实现的Clone方法
*/

func (self *ConcrectImage) Clone() Image {
	duplicate := &ConcrectImage{}
	*duplicate = *self
	return duplicate
}

/*
 Prototype类的功能方法
*/

func (self *ConcrectImage) GetImageName() {
	fmt.Println("这个图片是" + self.Name)
}

func (self *ConcrectImage) SetImageName(Name string) {
	self.Name = Name
}

/*
 Client程序
*/

type App struct {
}

func (*App) Verify() {
	image_a := ConcrectImage{Name: "教学楼正面图"}
	image_b := image_a.Clone()
	image_b.SetImageName("教学楼侧面图")
	image_a.GetImageName()
	image_b.GetImageName()
}
