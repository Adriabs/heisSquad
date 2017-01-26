// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i int = 0;


func thread1() {
    for j := 0; j<1e6; j++ {
        i++;
    } 
}

func thread2() {
    for j := 0; j<1e6;j++ {
        i--;
    } 
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!
    go thread1()                      // spawning threads
    go thread2()
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(100*time.Millisecond)
    Println(i)
}