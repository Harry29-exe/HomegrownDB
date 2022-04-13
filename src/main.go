package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	println("started")

	rwMutex := sync.RWMutex{}

	wg.Add(1)
	go func() {
		rwMutex.RLock()
		println("locked r")
		rwMutex.Lock()
		println("locked w")

		println("doing something")

		rwMutex.Unlock()
		println("unlocked w")

		rwMutex.Unlock()
		println("unlocked r")
		wg.Done()
	}()

	wg.Wait()
	println("finished")
}
