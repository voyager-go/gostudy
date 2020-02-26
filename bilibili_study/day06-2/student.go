// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/26 10:33 下午
package main

type Student struct {
	Name, Class, Gender string
	Age                 int8
	Address
}
type Address struct {
	Province, City string
}

// newStudent 新建一个学生
func (s *Student) newStudent(name, class, gender string, age int8, address Address) Student {
	return Student{
		Name:    name,
		Class:   class,
		Gender:  gender,
		Age:     age,
		Address: address,
	}
}

// newAddress 新建一个学生地址
func (a *Address) newAddress(province, city string) Address {
	return Address{
		Province: province,
		City:     city,
	}
}
