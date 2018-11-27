package main

import (
	"fmt"
	"sync"
	"time"
)

type chopS struct{ sync.Mutex }

type philo struct {
	leftCS, rightCS *chopS
	number          int
}

func (p philo) eat() {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat", p.number)
		fmt.Println(time.Now())
		fmt.Println("finishing to eat", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	maxGoroutines := 2
	host := make(chan struct{}, maxGoroutines)

	cSticks := make([]*chopS, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopS)
	}

	philos := make([]*philo, 5)
	for i := 0; i < 5; i++ {
		fmt.Println(i, (i+1)%5)
		philos[i] = &philo{cSticks[i], cSticks[(i+1)%5], i}
	}

	for i := 0; i < 5; i++ {
		host <- struct{}{}
		go philos[i].eat()
		<-host
	}

	time.Sleep(3 * time.Second)
}
