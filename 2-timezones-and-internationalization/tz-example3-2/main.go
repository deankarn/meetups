package main

import (
	"fmt"
	"time"
)

func main() {

	before := time.Date(2016, 03, 13, 1, 59, 59, 0, time.Local)
	after := time.Date(2016, 03, 13, 2, 0, 0, 0, time.Local)

	// diff since last event
	diff := after.Sub(before)

	fmt.Printf("Before:%s\n", before)
	fmt.Printf("After:%s\n", after)
	fmt.Printf("Diff:%s\n", diff)
}
