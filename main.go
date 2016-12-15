package main

import (
	"fmt"
	"github.com/advantageous/go-logback/logging"
)

func main() {

	buffer := make([]interface{}, 3)
	buffer[0] = "a"
	buffer[1] = "b"
	buffer[2] = "c"

	logging.NewSimpleLogger("foo")

	bufferCopy1 := buffer[1:]
	bufferCopy1[0] = "B"
	fmt.Printf("BUFFER CONTENTS %v\n", buffer)

	ch := make(chan interface{})

	go func() {
		ch <- buffer[1:]
	}()

	val := <-ch
	bufferCopy2 := val.([]interface{})
	bufferCopy2[0] = "BEE"
	fmt.Printf("BUFFER CONTENTS %v\n", buffer)

}
