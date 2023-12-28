package reflex

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	fields := getValue(x)

	switch fields.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < fields.Len(); i++ {
			walk(fields.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(fields.String())
	case reflect.Struct:
		for i := 0; i < fields.NumField(); i++ {
			walk(fields.Field(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range fields.MapKeys() {
			walk(fields.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for field, ok := fields.Recv(); ok; field, ok = fields.Recv() {
			walk(field.Interface(), fn)
		}
	case reflect.Func:
		fnResult := fields.Call(nil)
		for _, res := range fnResult {
			walk(res.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	fields := reflect.ValueOf(x)

	if fields.Kind() == reflect.Pointer {
		fields = fields.Elem()
	}
	return fields
}
