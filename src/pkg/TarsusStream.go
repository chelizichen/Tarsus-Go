package pkg

import (
	"fmt"
	"reflect"
	"strconv"
)

var StreamMap = make(map[string]func(args []any) any)

type TarsusStreamInf interface {
	ReadInt(index int) int
	ReadString(index int) string
	ReadList(index int, listType string)
	ReadStruct(index int, tarsusStruct string) any
	TarsusStream(args []any) TarsusStreamStruct
}

type TarsusStreamStruct struct {
	osData    []any
	clazzName string
	TarsusStreamInf
}

func (receiver *TarsusStreamStruct) ReadInt(index int) int {
	index = index - 1
	s := receiver.osData[index]

	if reflect.TypeOf(s).Kind() == reflect.Int {
		return s.(int)
	}
	if reflect.TypeOf(s).Kind() == reflect.String {
		atoi, err := strconv.Atoi(s.(string))
		if err == nil {
			return 0
		}
		return atoi
	}
	return 0
}

func (receiver *TarsusStreamStruct) ReadString(index int) string {
	index = index - 1

	s := receiver.osData[index]

	if reflect.TypeOf(s).Kind() == reflect.Int {
		s2 := fmt.Sprintf("%d", s)
		return s2
	}

	if reflect.TypeOf(s).Kind() == reflect.String {
		return fmt.Sprintf("%v", s)
	}

	return ""
}

func (receiver *TarsusStreamStruct) ReadStruct(index int, tarsusStruct string) any {
	index = index - 1

	data := receiver.osData[index]
	if reflect.TypeOf(data).Kind() == reflect.Array {
		f := StreamMap[tarsusStruct]
		return f(data.([]any))
	} else {
		return ""
	}
}

func TarsusStream(args []any, structType string) TarsusStreamStruct {
	var tarsusStreamStruct TarsusStreamStruct
	tarsusStreamStruct.osData = args
	tarsusStreamStruct.clazzName = structType
	return tarsusStreamStruct
}
