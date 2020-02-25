// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/26 12:08 上午
package main

import (
	"encoding/json"
	"fmt"
)

// 如果一个结构体中的字段名称首字母是大写的，那么该字段就是对外开放的(别的包可访问)
// 如果这里 ID 改为 iD，对于 json 包来讲，是无法访问到 iD 的
// 如果要修改对外的字段值，可以给字段加标签 如 `json:"my_id" db:"id"`
type student struct {
	ID   int
	Name string
}

type class struct {
	Title    string
	Students []student
}

func (s student) NewStudent(id int, name string) student {
	s.ID = id
	s.Name = name
	return s
}
func main() {
	var stu student
	var cla class
	cla = class{
		Title:    "三年级一班",
		Students: make([]student, 0, 60),
	}
	for i := 1; i <= 50; i++ {
		cla.Students = append(cla.Students, stu.NewStudent(i, fmt.Sprintf("stu%02d", i)))
	}
	//fmt.Printf("%#v", cla)
	jsonData, err := json.Marshal(cla)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%T\n", jsonData)
	fmt.Printf("%s\n", jsonData)
}
