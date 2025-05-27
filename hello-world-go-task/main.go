package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	for i, arg := range os.Args[1:] {
		fmt.Println("Arg", i, ": ", arg)
	}

	// Print time, sleep 60 seconds, repeat 30 times
	for i := 0; i < 30; i++ {
		fmt.Printf("Time %d: %s\n", i+1, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(60 * time.Second)
	}
}
