package main

import (
	"encoding/json"
	"log"
	"fmt"
)
//主函数
func main() {
	//将结构体，map，切片进行序列化
	testStruct()
	//testMap()
	//testSlice()

}
type jsonSerialize struct {
	//使用tag设置别名，反射机制
	Name string `json:"mingzi"`
	Age  int
	Birthday string
	Skill string
}
// 序列化结构体
func testStruct()  {
	m := jsonSerialize{
		Name: "youeryuan",
		Age: 25,
		Birthday: "1996-02-01",
		Skill: "coding",
	}
	jsonText, err := json.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonText))
}
// 序列化map
func testMap()  {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "youeryuan"
	a["age"] = "25"
	a["birthday"] = "1996-02-01"
	a["skill"] = "code"
	jsonString, err := json.Marshal(&a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonString))
}
// 序列化切片
func textSlice()  {
	var b []map[string]interface{}
	var a   map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "youeryuan"
	a["age"] = "25"
	a["birthday"] = "1996-02-01"
	a["skill"] = "code"
	b = append(b, a)
	var c map[string]interface{}
	c = make(map[string]interface{})
	a["name"] = "youeryuango"
	a["age"] = "29"
	a["birthday"] = "1990-02-01"
	a["skill"] = "play games"
	b = append(b, c)
	jsonString, err := json.Marshal(&b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonString))
}