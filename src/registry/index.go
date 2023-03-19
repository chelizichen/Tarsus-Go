package registry

// 记录所有服务
func LoadRegister() {
	h := new(Hello)
	h.InitRegister()
	f := new(Fish)
	f.InitRegister()

	h.tarsus.Call("Notify",[]any{1,2})
}
