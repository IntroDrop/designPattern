package adapter

import "fmt"

type Volts220 struct {
}

// OutputPower 接口实现
func (v Volts220) OutputPower() {
	fmt.Println("输出220v电压")
}

type Volts110 struct {
}

// OutputPower 接口实现
func (v Volts110) OutputPower() {
	fmt.Println("输出110v电压")
}

// Adaptee 源接口
type Adaptee interface {
	OutputPower()
}

// 目的接口
type Target interface {
	ConvertTo5V()
}

// Adapter 适配器结构体，嵌套源接口Adaptee，实现Target接口
type Adapter struct {
	// 继承
	Adaptee
}

func (a Adapter) ConvertTo5V() {
	// 调用源结构体的方法
	a.OutputPower()
	fmt.Println("适配5v电压")
}
