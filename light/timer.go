package main

import (
	"fmt"
	"time"
)

func runTimer() {
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println(1)
		}
	}
}
