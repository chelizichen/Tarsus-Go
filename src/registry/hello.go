package registry

import (
	"tarsus/go/src/pkg"
)

type HelloMethods interface {
	Notify()
	Say()
}

type Hello struct {
	message string
	name    string
	// 结构体继承 当前 Methods的所有方法 加上 注册结构体
	RegisterMethods HelloMethods
	tarsus          pkg.TarsusRegister
}

type any = interface {}

func (h *Hello) Notify(args ...any) any {
	println("--Notify--")
	println("name", args)
	//println(h.name, h.message)
	println("**********")
	return ""
}

func (h *Hello) Say(args ...any) any {
	println("--Hello--")
	println(h.name, h.message)
	println("**********")
	return ""
}

func (h *Hello) InitRegister() {
	h.tarsus.Register("Notify", h.Notify)
	h.tarsus.Register("Say", h.Say)
	//h.tarsus.TarsusInvoke.
	h.tarsus.TarsusInit("Hello Register")
}
