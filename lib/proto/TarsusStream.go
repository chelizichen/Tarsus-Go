package proto

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type any = interface{}

type TarsusStream struct {
	Arguments   []any
	ClazzName   string
	TargetClass reflect.Type
}

type ClassBasic struct {
	Clazz           reflect.Type
	Deserialization func(intermediate []interface{}) interface{}
}

var StreamMap = make(map[string]ClassBasic)

func SetClass(clazzName string, clazz ClassBasic) {
	StreamMap[clazzName] = clazz
}

func GetClass(clazzName string) ClassBasic {
	return StreamMap[clazzName]
}

func NewTarsusStream(arguments []interface{}, clazzName string) *TarsusStream {
	return &TarsusStream{
		Arguments:   arguments,
		ClazzName:   clazzName,
		TargetClass: GetClass(clazzName).Clazz,
	}
}

func (t *TarsusStream) ReadString(index int) string {
	return fmt.Sprintf("%v", t.Arguments[index-1])
}

func (t *TarsusStream) ReadInt(index int) int {
	val, _ := strconv.Atoi(fmt.Sprintf("%v", t.Arguments[index-1]))
	return val
}

func (t *TarsusStream) ReadStruct(index int, className string) interface{} {
	Clazz := GetClass(className)
	listArgs, ok := t.Arguments[index-1].([]interface{})
	if !ok {
		return nil
	}
	// Here, you'd need to call the appropriate constructor for the type.
	// This is a placeholder and will need to be adapted to your specific needs.
	parse := Clazz.Deserialization(listArgs)
	return parse
}

func (t *TarsusStream) ReadList(index int, className string) interface{} {
	getType := GetType(className)
	listArgs, ok := t.Arguments[index-1].([]interface{})
	if !ok {
		return nil
	}

	if isBaseType(getType) {
		args := make([]interface{}, 0)
		for _, arg := range listArgs {
			switch getType {
			case "int":
				val, _ := strconv.Atoi(fmt.Sprintf("%v", arg))
				args = append(args, val)
			case "string":
				args = append(args, fmt.Sprintf("%v", arg))
			}
		}
		return args
	} else {
		clazz := GetClass(getType)
		sliceType := reflect.SliceOf(clazz.Clazz.Elem()) // 注意这里的修改
		sliceValue := reflect.MakeSlice(sliceType, len(listArgs), len(listArgs))
		for i, arg := range listArgs {
			constructorArgs, ok := arg.([]interface{})
			if !ok {
				continue
			}
			deserialization := clazz.Deserialization(constructorArgs)
			sliceValue.Index(i).Set(reflect.ValueOf(deserialization).Elem())
		}
		return sliceValue.Interface()
	}
}

func GetType(typeStr string) string {
	re := regexp.MustCompile(`<(.+?)>`)
	matches := re.FindStringSubmatch(typeStr)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
func isBaseType(typeStr string) bool {
	baseTypes := []string{"int", "string", "bool"}
	for _, t := range baseTypes {
		if t == typeStr {
			return true
		}
	}
	return false
}
