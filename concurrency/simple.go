package main

// 000000
// 000001
// ...
// ffffff -> 16,777,215

import (
	"fmt"
	"os"
)

func main_simple() {
	f, err := os.Create("my_file.txt")
	if err != nil {
		panic(err)
	}

	final := 16777215
	for i := 0; i <= final; i++ {
		_, err = f.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}
	}
}
