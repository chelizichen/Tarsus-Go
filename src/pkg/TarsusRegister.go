package pkg

type TarsusRegister struct {
	// 注冊接口名
	register string
}

func (t *TarsusRegister) TarsusInit(registerName string) {
	t.register = registerName
	println("Tarsus 服務初始化", t.register)
}
