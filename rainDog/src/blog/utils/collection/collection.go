package collection

import (
	"fmt"
)

const (
	CIntArray 	 = iota
	CString
	CStringArray
	CMap
	CMapArray
	CStructArray
)

type Collection struct {
	value interface{}
	cType interface{}
}

func Collect(obj interface{}, cType int) Collection {
	return Collection{
		value: obj,
		cType: cType,
	}
}

func CollectStruct(obj interface{}, cType interface{}) Collection {
	return Collection{
		value: obj,
		cType: cType,
	}
}

func (c Collection) GroupBy(k string) Collection {
	//var d = make(map[string]interface{}, 0)
	fmt.Println(c.cType)
	//fmt.Println(c.value.(c.cType))
	for {
		fmt.Println(c.value)
	}
	return c
}
