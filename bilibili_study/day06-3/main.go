// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/26 11:36 下午
package main

import (
	"fmt"
	"os"
)

// ShowMenu 打印提示信息
func ShowMenu() {
	fmt.Println("Welcome to student management system...")
	fmt.Println("0. Add student")
	fmt.Println("1. Edit student")
	fmt.Println("2. Show students")
	fmt.Println("3. Exit system")
}

// HandleUserInput 处理用户输入
func HandleUserInput() *Student {
	var (
		id                  int
		age                 int8
		name, class, gender string
	)
	fmt.Println("Please input student information as required.")
	fmt.Print("0. Please enter the student's id:")
	_, _ = fmt.Scanf("%d\n", &id)
	fmt.Print("1. Please enter the student's name:")
	_, _ = fmt.Scanf("%s\n", &name)
	fmt.Print("2. Please enter the student's class:")
	_, _ = fmt.Scanf("%s\n", &class)
	fmt.Print("3. Please enter the student's gender:")
	_, _ = fmt.Scanf("%s\n", &gender)
	fmt.Print("4. Please enter the student's age:")
	_, _ = fmt.Scanf("%d\n", &age)
	return newStudent(id, name, class, gender, age)
}
func main() {
	var student *Student
	var studentHandle *StudentHandel
	studentHandle = newStudentHandle()
	for true {
		ShowMenu()
		var userInput int
		_, err := fmt.Scanf("%d\n", &userInput)
		if err != nil {
			panic("Input error")
		}
		switch userInput {
		case 0:
			student = HandleUserInput()
			studentHandle.createStudent(student)
			break
		case 1:
			student = HandleUserInput()
			studentHandle.modifyStudent(student)
		case 2:
			studentHandle.showStudents()
			break
		case 3:
			os.Exit(0)
		}
	}
}
