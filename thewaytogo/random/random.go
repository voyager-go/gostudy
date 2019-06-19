package main

import (
	"math/rand"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		aa := rand.Int()
		fmt.Printf("%d\n", aa)
	}
	for ii := 0; ii < 10; ii++ {
		bb := rand.Intn(5)
		fmt.Printf("%d\n", bb)
	}

	//timens := int64(time.Now().Nanosecond())
	//rand.Seed(timens)
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%2.2f / ", 100*rand.Float32())
	//}
}
