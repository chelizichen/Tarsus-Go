package registry

// 记录所有服务
func LoadRegister() {
	h := new(Hello)
	h.InitRegister()
	f := new(Fish)
	f.InitRegister()

	map1 := h.tarsus.GetRegisterMap()
	//h.tarsus.Call()
}
