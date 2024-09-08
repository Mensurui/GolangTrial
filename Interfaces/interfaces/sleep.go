package interfaces

import (
	"flag"
	"fmt"
	"time"
)

var Period = flag.Duration("Period", 1*time.Second, "Sleep Period")

func Sleep() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *Period)
	time.Sleep(*Period)
	fmt.Println()
}
