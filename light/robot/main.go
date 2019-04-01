package main

import (
	"sync"
	"time"
)

var (
	num = 5            // 机器人数量
	wg  sync.WaitGroup // 等待组
)

func main() {
	for i := 0; i < num; i++ {
		wg.Add(1)

		go newRobot()

		time.Sleep(3 * time.Second)
	}

	wg.Wait()
}

func newRobot() {
	defer wg.Done()

	r := NewRobot()
	r.Run()
}
