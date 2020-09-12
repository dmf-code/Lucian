package collection

import (
	"fmt"
	"reflect"
)

type Collection struct {
	value interface{}
}

func Collect(obj interface{}) Collection {
	return Collection{
		value: obj,
	}
}

func (c Collection) GroupBy(k string) Collection {
	//var d = make(map[string]interface{}, 0)
	elements := reflect.ValueOf(c.value)
	for i := 0; i < elements.Len(); i++ {
		element := elements.Index(i)
		fmt.Println(element)
		fmt.Println(element.Type())
		fmt.Println(element.Kind())
		fmt.Println(element.NumField())
		fmt.Println(element.FieldByName(k))
		fmt.Println(element.FieldByName(k).String())
	}
	return c
}
