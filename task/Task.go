package task

import "time"

func Init() {
	go ticker()
}

func ticker() {
	// 创建一个计时器
	minute := time.Tick(time.Minute)
	for {
		select {
		case <-minute:
			projectTask()
		}
	}
}
