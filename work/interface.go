package work

import (
	"fmt"
	"unsafe"
)

type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bycycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}

func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

func (v *bycycle) speedUp() int {
	v.speed += 3 * v.humanPower
	return v.speed
}

func (v *bycycle) speedDown() int {
	v.speed -= 1 * v.humanPower
	return v.speed
}

func speedUpAndDonw(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

// Stringerのインターフェースを満たすものとしている
func (v vehicle) String() string {
	return fmt.Sprintf("Vehicle current speed is %v (enginePower %v)", v.speed, v.enginePower)
}

func interface_() {
	v := &vehicle{0, 5}
	speedUpAndDonw(v)
	b := &bycycle{0, 5}
	speedUpAndDonw(b)
	fmt.Println(v)

	var i1 interface{}
	var i2 any
	fmt.Printf("%[1]v %[1]T %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]v %[1]T %v\n", i1, unsafe.Sizeof(i2))
	checkType(i2)
	i2 = 1
	checkType(i2)
	i2 = "hello"
	checkType(i2)
}

func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Printf("nil")
	case int:
		fmt.Printf("int")
	case string:
		fmt.Printf("string")
	default:
		fmt.Printf("unknown")
	}
}
