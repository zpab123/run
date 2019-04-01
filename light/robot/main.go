package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	num      = 10           // 机器人数量
	sendCont = 100          // 发送次数
	wg       sync.WaitGroup // 等待组
)

func main() {

	t1 := time.Now()

	for i := 0; i < num; i++ {
		wg.Add(1)

		go newRobot()

		// time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()

	elapsed := time.Since(t1)

	fmt.Println("Total count:", num*sendCont)
	fmt.Println("Success count:", num*sendCont)
	fmt.Println("Cysle TPS:", float64(num*sendCont)/elapsed.Seconds())
	fmt.Println("Taken Time(s) :", elapsed)
	fmt.Println("Average Latency time(ms):", elapsed.Seconds()*1000/(float64(num*sendCont)))
}

func newRobot() {
	defer wg.Done()

	r := NewRobot()
	r.count = sendCont

	r.Run()
}
