package foo

import (
	"fmt"
	"os"
)

func Crasher() {
	fmt.Println("Going down in flames!")
	os.Exit(1)
}

func NotCrasher() {
	fmt.Println("Not crashing")
}
