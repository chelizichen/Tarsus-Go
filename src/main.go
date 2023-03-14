package main

import (
	"tarsus/go/src/pkg"
	"tarsus/go/src/registry"
)

func main() {
	hello := registry.Hello{}
	hello.Notify()
	pkg.Tarsus()
}

