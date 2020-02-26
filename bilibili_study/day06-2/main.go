// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/26 10:33 下午
package main

import (
	"fmt"
	"os"
)

var StudentsList = make([]Student, 0, 100)

func ShowMenu() {
	fmt.Println("Welcome to student management system...")
	fmt.Println("0. Add student")
	fmt.Println("1. Edit student")
	fmt.Println("2. Show students")
	fmt.Println("3. Exit system")
}
func main() {
	for true {
		ShowMenu()
		var userInput int
		_, err := fmt.Scanf("%d\n", &userInput)
		if err != nil {
			panic("Input error")
		}
		switch userInput {
		case 0:
			var stu Student
			fmt.Println("0. Please enter the student's name")
			_, _ = fmt.Scanf("%s\n", &stu.Name)
			fmt.Println("1. Please enter the student's gender")
			_, _ = fmt.Scanf("%s\n", &stu.Gender)
			fmt.Println("2. Please enter the student's class")
			_, _ = fmt.Scanf("%s\n", &stu.Class)
			fmt.Println("3. Please enter the student's age")
			_, _ = fmt.Scanf("%d\n", &stu.Age)
			fmt.Println("3. Please enter the student's province")
			_, _ = fmt.Scanf("%s\n", &stu.Province)
			fmt.Println("3. Please enter the student's city")
			_, _ = fmt.Scanf("%s\n", &stu.City)
			StudentsList = append(StudentsList, stu)
			break
		case 1:
		case 2:
			for index, student := range StudentsList {
				fmt.Println(index, ":", student)
			}
			break
		case 3:
			os.Exit(0)

		}
	}
}
