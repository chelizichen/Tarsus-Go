package invoke

import (
	"errors"
	"reflect"
)
type Method struct {
	name string
	fn   interface{}
	args []reflect.Type
}

type Class struct {
	name    string
	methods []*Method
}

var clazzMap = make(map[string]*Class)

func GetClass(obj interface{}) (*Class, error) {
	t := reflect.TypeOf(obj)

	name := t.Name()
	if(clazzMap[name] != nil){
		println("走单例")
		return clazzMap[name], nil;
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("obj is not struct")
	}
	class := &Class{
		name:    t.Name(),
		methods: []*Method{},
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fn := m.Func.Interface()
		args := make([]reflect.Type, m.Type.NumIn()-1)
		for i := 1; i < m.Type.NumIn(); i++ {
			args[i-1] = m.Type.In(i)
		}
		method := &Method{
			name: m.Name,
			fn:   fn,
			args: args,
		}
		class.methods = append(class.methods, method)
	}
	clazzMap[name] = class
	return class, nil
}

func Invoke(obj interface{}, methodName string, params ...interface{}) (result []interface{}, err error) {
	class, err := GetClass(obj)
	if err != nil {
		return nil, err
	}
	for _, method := range class.methods {
		// 方法是否对应
		if method.name == methodName {
			if len(params) != len(method.args) {
				return nil, errors.New("params mismatch")
			}
			in := make([]reflect.Value, len(params))
			for i := 0; i < len(params); i++ {
				in[i] = reflect.ValueOf(params[i])
			}
			fn := reflect.ValueOf(method.fn)
			out := fn.Call(append([]reflect.Value{reflect.ValueOf(obj)}, in...))
			result = make([]interface{}, len(out))
			for i := 0; i < len(out); i++ {
				result[i] = out[i].Interface()
			}
			return result, nil
		}
	}
	return nil, errors.New("method not found")
}
