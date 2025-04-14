package gotypewriter

import (
	"fmt"
	"time"
)

func Print(s string) {
	for _, c := range s {
		fmt.Printf("%c", c)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	Print("Hello, World!")
}
