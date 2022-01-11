package adapter

import "testing"

func Test(t *testing.T) {
	t.Run("adaptor:", ConvertVolts)
}

func ConvertVolts(t *testing.T) {
	// 源接口的实现类：Volts220
	var adaptee Adaptee
	adaptee = &Volts220{}

	// target的实现类：Adapter（初始化Adaptee为adaptee对象）（继承自Volts220）
	var target Target
	target = &Adapter{adaptee}

	target.ConvertTo5V()
}
