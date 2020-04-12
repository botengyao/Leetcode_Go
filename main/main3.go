package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(runtime.Version())
	fmt.Println("OK")
}
