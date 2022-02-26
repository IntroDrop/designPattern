package DesignPattern

import (
	"fmt"
	"math"
	"testing"
)

func Test(t *testing.T) {
	bitset := Constructor(2)
	bitset.Flip()
	bitset.Fix(3)
	bitset.Fix(1)

	print(bitset.All())
	bitset.Unfix(0)
	bitset.Flip()
	print(bitset.One())
	bitset.Unfix(0)
	print(bitset.Count())
	fmt.Println(bitset.ToString())
	fmt.Println(len(bitset.ToString()))
}

type Bitset struct {
	data []uint64
	size int
	bia  int
}

func Constructor(size int) Bitset {
	var v int
	if size%64 == 0 {
		v = 0
	} else {
		v = 1
	}
	return Bitset{
		data: make([]uint64, size/64+v),
		size: size,
		bia:  (size/64+v)*64 - size,
	}
}

func (this *Bitset) Fix(idx int) {
	i := len(this.data) - 1 - idx/64 // 数组
	idx %= 64                        // 具体
	idx += this.bia
	this.data[i] |= 1 << idx
}

func (this *Bitset) Unfix(idx int) {
	i := len(this.data) - 1 - idx/64 // 数组
	idx %= 64                        // 具体
	idx += this.bia                  // 具体
	this.data[i] &= ^(1 << idx)      // 置0
}

func (this *Bitset) Flip() {
	for i, i2 := range this.data {
		this.data[i] = ^i2
	}
}

func (this *Bitset) All() bool {
	// 不全为1
	for i := 0; i < len(this.data); i++ {
		var data uint64
		if i == len(this.data)-1 {
			data = this.data[i] >> this.bia
			cout := 0
			for data != 0 {
				cout++
				data = data & (data - 1)
			}
			if cout != 64-this.bia {
				return false
			}
		} else {
			data = this.data[i]
			if data != math.MaxUint64 {
				return false
			}
		}
	}
	return true
}

func (this *Bitset) One() bool {
	for i := 0; i < len(this.data); i++ {
		if i == len(this.data)-1 {
			if this.data[i]>>this.bia != 0 {
				return true
			}
		} else {
			if this.data[i] != 0 {
				return true
			}
		}
	}
	return false
}

func (this *Bitset) Count() int {
	ans := 0
	for i := 0; i < len(this.data); i++ {
		cout := 0
		var data uint64
		if i == len(this.data)-1 {
			data = this.data[i] >> this.bia
		} else {
			data = this.data[i]
		}
		for data != 0 {
			cout++
			data = data & (data - 1)
		}
		ans += cout
	}
	return ans
}

func (this *Bitset) ToString() string {
	ans := ""
	for i := 0; i < len(this.data); i++ {
		if i == len(this.data)-1 {
			s := fmt.Sprintf("%b", this.data[i]>>this.bia)
			for j := 0; j < (64 - this.bia - len(s)); j++ {
				ans += "0"
			}
			ans += s
		} else {
			s := fmt.Sprintf("%b", this.data[i])
			for j := 0; j < (64 - len(s)); j++ {
				ans += "0"
			}
			ans += s
		}
	}

	return Reverse(ans)
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
