package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.2
	fmt.Println(reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println(v)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	fmt.Println(v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println(v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.14150)
	fmt.Println(v.Interface())
	fmt.Println(v)

	fmt.Println("----------- reflect struct -----------")
	value := reflect.ValueOf(secret)
	fmt.Println(value)
	fmt.Println(reflect.TypeOf(secret))
	kin := value.Kind()
	fmt.Println(kin)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
	results := value.Method(0).Call(nil)
	fmt.Println(results)
	fmt.Println("------------结构中只有被导出字段（首字母大写）才是可设置的--------------")
	t := T{23, "Bob"}
	val := reflect.ValueOf(&t).Elem()
	typeOfT := val.Type()
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	val.Field(0).SetInt(56)
	val.Field(1).SetString("mmp")
	fmt.Println("t is now", t)
}

type UnknownType struct {
	s1, s2, s3 string
}

func (n UnknownType) String() string  {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = UnknownType{"Ada", "Go", "Python"}

type T struct {
	A int
	B string
}