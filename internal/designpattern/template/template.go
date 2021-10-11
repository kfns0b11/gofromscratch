package template

import "fmt"

type Cooker interface {
	fire()
	cook()
	outfire()
}

type CookMeau struct {
}

func (CookMeau) fire() {
	fmt.Println("fire")
}

func (CookMeau) outfire() {
	fmt.Println("outfire")
}

func DoCook(cooker Cooker) {
	cooker.fire()
	cooker.cook()
	cooker.outfire()
}

type XiHongShi struct {
	CookMeau
}

func (XiHongShi) cook() {
	fmt.Println("cook XiHongShi")
}

func DoCookXiHongShi() {
	DoCook(XiHongShi{})
}
