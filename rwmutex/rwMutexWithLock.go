package main

// cannot do anything when you do a Lock()
// 當你使用Lock()時，無法做任何事，包含RLock(), Lock()

import (
	"sync"
	"time"
)

type rwMutexWithLock struct {
	m *sync.RWMutex
}

func main() {

	rwLock := &rwMutexWithLock{m: new(sync.RWMutex)}

	go rwLock.write(1)
	go rwLock.read(2)
	go rwLock.write(3)

	time.Sleep(2 * time.Second)
}

func (self *rwMutexWithLock) read(i int) {
	println(i, "read start")

	self.m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	self.m.RUnlock()

	println(i, "read over")
}

func (self *rwMutexWithLock) write(i int) {
	println(i, "write start")

	self.m.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	self.m.Unlock()

	println(i, "write over")
}
