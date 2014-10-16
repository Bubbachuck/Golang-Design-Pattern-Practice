package singleton

/*
 要实现单一实例的类
*/

type singleton struct {
	name       string
	attributes string
}

/*
 单一实例初始化
*/

var single singleton = singleton{}

/*
 获取单一实例
*/

func GetInstance() *singleton {
	return &single
}

/*
 设置单一实例属性的方法
*/

func SingletonSetName(Name string) {
	single.name = Name
}

func SingletonSetAttributes(Attributes string) {
	single.attributes = Attributes
}

/*
 获取单一实例属性的方法
*/

func SingletonGetName() string {
	return single.name
}

func SingletonGetAttributes() string {
	return single.attributes
}
