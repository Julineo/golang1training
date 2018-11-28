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
	permis chan int
	done chan int
}

func (p *philo) eat() {
	for i := 0; i < 3; i++ {
		permis := <-p.permis
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat", p.number)
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
		fmt.Println("finishing to eat", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		p.done <- permis
	}
}

// limits number of eating philos
func host(permis chan int, done chan int) {
	permis <- 1
	permis <- 1

	for {
		permis <- <-done
	}
}

func main() {
	var permis = make(chan int, 2)
	var done = make(chan int, 2)

	cSticks := make([]*chopS, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopS)
	}

	philos := make([]*philo, 5)
	for i := 0; i < 5; i++ {
		fmt.Println(i, (i+1)%5)
		philos[i] = &philo{cSticks[i], cSticks[(i+1)%5], i, permis, done}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	go host(permis, done)

	time.Sleep(10 * time.Second)
}
