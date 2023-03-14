package registry

import (
	"tarsus/go/src/pkg"
)

type Hello struct {
	message string
	name    string
	tarsus  pkg.TarsusRegister
}

func (u *Hello) Notify() {
	println(u.name, u.message)
}

func init() {
	user := new(Hello)
	user.message = "134@qq.com"
	user.name = "你好"
	user.tarsus.TarsusInit("測試Hello 接口")
	println("測試初始化")
}
