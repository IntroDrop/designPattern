package factory

import "fmt"

// 简单工厂：把所有的对象的构造都放在一个工厂里
type FruitFactory struct{}

type Fruit interface {
	Eat()
}

type Apple struct{}
type Banana struct{}
type Orange struct{}

func (apple Apple) Eat() {
	fmt.Println("吃苹果")
}

func (banana Banana) Eat() {
	fmt.Println("吃Banana")
}

func (o Orange) Eat() {
	fmt.Println("吃Orange")
}

func (FruitFactory) create(Type string) Fruit {
	switch Type {
	case "Apple":
		return new(Apple)
	case "Banana":
		return new(Banana)
	case "Orange":
		return new(Orange)
	default:
		return nil
	}
}

// 普通工厂：每个产品都有一个专属工厂

//AppleFactory 苹果工厂
type AppleFactory struct{}

//BananaFactory 香蕉工厂
type BananaFactory struct{}

//OrangeFactory 橘子工厂
type OrangeFactory struct{}

// CreateFruit 苹果工厂生产苹果
func (appleFactory AppleFactory) CreateFruit() Fruit {
	return &Apple{}
}

// CreateFruit 香蕉工厂生产香蕉
func (bananaFactory BananaFactory) CreateFruit() Fruit {
	return &Banana{}
}

// CreateFruit 橘子工厂生产橘子
func (orangeFactory OrangeFactory) CreateFruit() Fruit {
	return &Orange{}
}
