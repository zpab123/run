package main

import (
	"github.com/zpab123/sco"
)

func main() {
	m := newmatch()
	go m.run()

	sco.Run()
}
