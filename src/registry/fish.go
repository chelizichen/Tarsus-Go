package registry

import (
	"tarsus/go/src/pkg"
)

type FishMethods interface {
	Swim()
	Fly()
}

type Fish struct {
	message string
	name    string
	// 结构体继承 当前 Methods的所有方法 加上 注册结构体
	RegisterMethods FishMethods
	tarsus          pkg.TarsusRegister
}

func (h *Fish) Swim(args []any) any {
	println("--Swim--")
	println("name", args)
	//println(h.name, h.message)
	println("**********")
	return ""
}

func (h *Fish) Fly(args []any) any {
	println("--Fly--")
	println(h.name, h.message)
	println("**********")
	return ""
}

func (h *Fish) InitRegister() {
	h.tarsus.Register("Swim", h.Swim)
	h.tarsus.Register("Fly", h.Fly)
	h.tarsus.TarsusInit("Fish Register")
}
