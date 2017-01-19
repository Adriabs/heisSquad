// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i int = 0

type empty {}
type semaphore chan empty


func thread1() {
    for j := 0; j<1e6; j++ {
	Lock(sem)
        i++
	Unlock(sem)
    } 
}

func thread2() {
    for j := 0; j<1e6;j++ {
	Lock(sem)
        i--
	Unlock(sem)
    } 
}

/* Resource aquiring and releasing */

// acquire n resources
func (s semaphore) P(n int) {
    e := empty{}
    for i := 0; i < n; i++ {
        s <- e
    }
}

// release n resources
func (s semaphore) V(n int) { 
    for i := 0; i < n; i++ {
        <-s
    }
}



/* mutexes */

func (s semaphore) Lock() {
    s.P(1)
}

func (s semaphore) Unlock() {
    s.V(1)
}

/* signal-wait */

func (s semaphore) Signal() {
    s.V(1)
}

func (s semaphore) Wait(n int) {
    s.P(n)
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!

    sem = make(semaphore, 1)		// making semaphore channel of length 1
    go thread1()                      // spawning threads
    go thread2()
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(100*time.Millisecond)
    Println(i)
}