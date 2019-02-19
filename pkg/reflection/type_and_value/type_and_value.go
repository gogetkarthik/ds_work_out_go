package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	fmt.Println(formatAtom(reflect.ValueOf(3)))
	fmt.Println(formatAtom(reflect.ValueOf("karthik")))
	fmt.Println(formatAtom(reflect.ValueOf(true)))
	fmt.Println(formatAtom(reflect.ValueOf([]string{"kk"})))
	fmt.Println(formatAtom(reflect.ValueOf(map[string]string{"KK1": "KK1"})))

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
