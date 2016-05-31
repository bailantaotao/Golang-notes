package main

// multiple goroutine read data at the same time in RLock()
// 允許同時多個goroutine進行讀取

import (
	"sync"
	"time"
)

type rwMutexWithRLock struct {
	m *sync.RWMutex
}

func main() {
	rwRLock := &rwMutexWithRLock{m: new(sync.RWMutex)}

	// read at the same time
	go rwRLock.read(1)
	go rwRLock.read(2)

	time.Sleep(2 * time.Second)
}

func (self *rwMutexWithRLock) read(i int) {
	println(i, "read start")

	self.m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	self.m.RUnlock()

	println(i, "read over")
}
