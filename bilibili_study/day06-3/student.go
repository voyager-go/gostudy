// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/26 11:36 下午
package main

import "fmt"

type Student struct {
	id                  int
	age                 int8
	name, class, gender string
}
type StudentHandel struct {
	students []*Student
}

// newStudent 新构建一个学生
func newStudent(id int, name, class, gender string, age int8) *Student {
	return &Student{
		id:     id,
		age:    age,
		name:   name,
		class:  class,
		gender: gender,
	}
}

// newStudentHandle 初始化一个学生管理句柄
func newStudentHandle() *StudentHandel {
	return &StudentHandel{
		students: make([]*Student, 0, 100),
	}
}

// createStudent 追加一个学生
func (s *StudentHandel) createStudent(student *Student) {
	s.students = append(s.students, student)
}

// modifyStudent 修改学生信息
func (s *StudentHandel) modifyStudent(student *Student) {
	for index, stu := range s.students {
		if student.id == stu.id {
			s.students[index] = student
			return
		}
	}
	fmt.Printf("未找到学号为：%d 的学生\n", student.id)
}

// showStudents 展示全部学生信息
func (s *StudentHandel) showStudents() {
	for _, val := range s.students {
		fmt.Printf("学号为：%d, 姓名为：%s, 班级为：%s, 性别为：%s, 年龄为：%d\n", val.id, val.name, val.class, val.gender, val.age)
	}
}
