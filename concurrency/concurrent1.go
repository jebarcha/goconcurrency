package main

import (
	"fmt"
	"os"
)

func main_concurrent1() {
	f, err := os.Create("my_file.txt")
	if err != nil {
		panic(err)
	}

	numGoRoutines := 10
	doneCh := make(chan struct{})

	final := 16777215
	for i := 0; i <= final; i = i + (final / numGoRoutines) + 1 {
		paso := i + (final / numGoRoutines)
		if paso > final {
			paso = final
		}
		fmt.Printf("executing %d %d\n", i, paso)
		go calcNums1(i, paso, f, doneCh)
	}

	doneNum := 0
	for doneNum < numGoRoutines {
		<-doneCh
		fmt.Println("Completed")
		doneNum++
	}
	fmt.Println("Finish!! :D")
}

func calcNums1(start, end int, f *os.File, doneCh chan struct{}) {
	for i := start; i <= end; i++ {
		_, err := f.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}
	}
	doneCh <- struct{}{}
}
