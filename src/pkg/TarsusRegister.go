package pkg

// 新增注册表方法Map
var registerMap = make(map[string]func(args ...any) any)

type TarsusInvoke interface {
	// 	在调用前处理
	//beforeInvoke(args []string)

	// 调用
	Call(methodName string, args any) any

	// 初始化注册
	InitRegister()

	// 初始化
	TarsusInit(registerName string)

	// 注册 map
	Register(methodName string, method func(args ...any))

	// 拿到方法表
	GetRegisterMap() map[string]func(args ...any) any
}

type TarsusRegister struct {
	// 注冊接口名
	register  string
	canInvoke bool
	TarsusInvoke
}

func (t *TarsusRegister) TarsusInit(registerName string) {
	t.register = registerName
	println("Tarsus 服務初始化 |", t.register, "| 已加载")
}

func (t *TarsusRegister) Register(methodName string, method func(args ...any) any) {
	registerMap[methodName] = method
}

func (t *TarsusRegister) Call(methodName string, args ...any) any {
	return registerMap[methodName](args)
}

func (t *TarsusRegister) GetRegisterMap() map[string]func(args ...any) any {
	println("map 长度", len(registerMap))
	return registerMap
}
