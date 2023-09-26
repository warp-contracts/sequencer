package tools

import (
	"time"
	"fmt"
)

func Timer(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %d ms\n", name, time.Since(start).Milliseconds())
    }
}