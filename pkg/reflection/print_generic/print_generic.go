package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	Display("name", []string{"karthik", "haasini", "chandhu", "bhavi"})

	type sStruct struct {
		name string
		age  int
	}

	Display("sStruct", []sStruct{{"Karthik", 38}, {"haasini", 2}})
}

func Display(name string, val interface{}) {
	display(name, reflect.ValueOf(val))
}

func display(path string, v reflect.Value) {

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)

			display(fPath, v.Field(i))
		}
	default:
		fmt.Printf("%s == %s \n", path, formatAtom(v))
	}

}

func formatAtom(val reflect.Value) string {

	switch val.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())
	case reflect.String:
		return strconv.Quote(val.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return val.Type().String() + " 0x " + strconv.FormatUint(uint64(val.Pointer()), 16)
	default:
		return val.Type().String() + " value "

	}
}
