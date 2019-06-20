package main

import "fmt"

func classifier(params ...interface{}) {
	for i, x := range params {
		switch x.(type) {
		case int64:
			fmt.Printf("Param #%d is a int64\n", i)
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
func main() {
	classifier(int64(13), -14.3, "BELGIUM", complex(1, 2), nil, false)
}
