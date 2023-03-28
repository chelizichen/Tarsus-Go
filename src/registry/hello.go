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
}

type any = interface {}
