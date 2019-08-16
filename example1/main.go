// using example: GOSSAFUNC=stringByReflect go tool compile main.go

package main

import (
	"fmt"
	"reflect"
)

func stringByReflect(v interface{}) string {
	rv := reflect.ValueOf(v)
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return string(rv.Bytes())
		}
	case reflect.String:
		return rv.String()
	}
	return ""
}

func stringIdeomatic(v interface{}) string {
	switch rv := v.(type) {
	case []byte:
		return string(rv)
	case string:
		return rv
	}
	return ""
}

func main() {
	fmt.Printf("%s %s\n%s %s\n",
		stringIdeomatic("test"), stringByReflect("test"),
		stringIdeomatic([]byte("test")), stringByReflect([]byte("test")))
}
