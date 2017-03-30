package main

import (
	"fmt"
	"time"
)

// Event is just a generic event
type Event struct {
	slow bool
}

func main() {

	var after, before time.Time

	before = time.Now()

	// a loop where an event is being waited for, in this case using time.Sleep
	for {

		time.Sleep(time.Second * 1)

		var event *Event // nil by default

		after = time.Now()

		// diff since last event
		diff := after.Sub(before)

		// now using diff in some logic
		if diff > time.Minute*30 {
			// save event to database
			event = &Event{slow: true}
		} else if diff > time.Minute*5 {
			event = &Event{slow: false}
		}

		// use event in some logic, like saving event to database
		fmt.Println(event.slow)

		before = after
	}
}
