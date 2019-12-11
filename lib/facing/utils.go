package facing

import "reflect"

func defaultNewObj() interface{} {
	return "."
}

var defaultType string

func typeEqual(a, b interface{}) bool {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)
	return t1 == t2
}
