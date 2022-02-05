package main

import (
	"fmt"
	"hellowasm/warppers"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("ToPlaybook", warppers.ToPlaybookWrapper())
	//	js.Global().Set("ToPlaybookFromURL", toPlaybookWrapper())

	<-make(chan bool)
}
