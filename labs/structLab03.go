package main
// Go 语言的类型或结构体没有构造函数的功能。结构体的初始化过程可以使用函数封装实现


// 多种方式创建和初始化结构体——模拟构造函数重载
type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

// 带有父子关系的结构体的构造和初始化——模拟父级构造调用
type BlackCat struct {
	Cat // 嵌入Cat, 类似于派生
}

// 构造基类
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// 构造子类

func NewBlockCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}