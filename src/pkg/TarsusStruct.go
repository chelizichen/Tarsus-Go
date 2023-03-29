package pkg

// TarsusStructArg 字段
type TarsusStructArg struct {
	// 下标
	streamIndex int
	// 字段名
	streamName string
	// 字段类型
	streamType string
}

// TarsusStructArgs 所有字段
type TarsusStructArgs struct {
	structs []TarsusStructArg
}

// TarsusStruct 结构体
type TarsusStruct struct {
	name       string
	streamArgs TarsusStructArgs
}

type TarsusStructMaps struct {
	StructMaps map[string]TarsusStruct
}

//type TarsusInf interface {
//	SetStruct(name string,TStruct TarsusStruct)
//	GetStruct(name string)
//}

var TarSusStructMAPS = make(map[string]any)

//func (t TarsusStructMaps) SetStruct(name string,TStruct TarsusStruct)  {
//	TarSusStructMAPS[name] = TStruct;
//}
//
//func (t TarsusStructMaps) GetStruct(name string)TarsusStruct  {
//	tarsusStruct := TarSusStructMAPS[name]
//	return tarsusStruct;
//}
