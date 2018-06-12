package simple_impl

import (
	"fmt"
	"reflect"
)

func MeasureObject(obj interface{}) uint {
	if reflect.ValueOf(obj).IsNil() {
		return 0
	}
	value := reflect.ValueOf(obj)
	ve := value.Elem() // value对应的真实Element

	if ve.Kind() != reflect.Struct {
		panic("no struct type")
	}
	var num uint = 0
	measure(ve, &num)
	fmt.Printf("total size: %v\n", num)
	return num
}

func measure(v reflect.Value, num *uint) {
	t := v.Type()
	switch t.Kind() {
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
		*num += uint(t.Size())
		if reflect.Value.IsNil(v) {
			return
		}
		keys := v.MapKeys()
		for i := 0; i < len(keys); i++ {
			measure(keys[i], num)
			hashVal := v.MapIndex(keys[i])
			measure(hashVal, num)
		}

	case reflect.Ptr:
		*num += uint(t.Size())
		if reflect.Value.IsNil(v) {
			return
		}
		// 对于ptr来说, 还是可以进一步求解的.
		measure(v.Elem(), num)
	case reflect.Slice, reflect.Array:
		*num += uint(t.Size())
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			if int(item.Kind()) > 16 && item.Kind() != reflect.String && reflect.Value.IsNil(item) {
				continue
			}
			measure(item, num)
		}
	case reflect.String:
		*num += uint(t.Size())
		*num += uint(len(v.String()))
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			measure(v.Field(i), num)
		}
	case reflect.UnsafePointer:
	default:
		*num += uint(t.Size())
	}
}
