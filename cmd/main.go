package main

import (
	"fmt"
	"hellowasm/warppers"
	"syscall/js"
)

func main() {
	js.Global().Set("ToPlaybook", warppers.ToPlaybookWrapper())
	js.Global().Set("ToPlaybookFromURL", warppers.ToPlaybookFromURLWrapper())
	fmt.Println("Go Web Assembly")

	<-make(chan bool)
	<-make(chan bool)
}
