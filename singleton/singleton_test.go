package singleton

import (
	"fmt"
	"testing"
)

func TestGetSingleton(t *testing.T) {
	/* 执行结果：
	   执行获取单例对象
	   instance1 = 0xc0000584a0, name=标题1
	   instance2 = 0xc0000584a0, name=标题1
	*/

	// 第一次调用，instance里是"标题1"
	instance1 := GetInstance("标题1")
	fmt.Println(fmt.Sprintf("instance1 = %p, name=%s", instance1, instance1.Title))

	// 之后调用，instance并没有重新被创建，拿到的还是原来的instance，取出的title还是"标题1"
	for i := 0; i < 10; i++ {
		go func() {
			instance2 := GetInstance(fmt.Sprintf("标题%d", i))
			fmt.Println(fmt.Sprintf("instance2 = %p, name=%s", instance2, instance2.Title))
		}()
	}

}
