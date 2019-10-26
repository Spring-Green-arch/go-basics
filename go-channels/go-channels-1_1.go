
package main

import ( "fmt" 
    "sync" 
)

var wg = sync.WaitGroup{}
var ch = make(chan int)

func sendI(i int) {
    ch <- i
    wg.Done()
}

func receiveI() {
    i := <- ch
    fmt.Println(i)
    wg.Done()
}

func main() {
    i := 10
    wg.Add(2)
    go sendI(i)   
    go receiveI()
    wg.Wait()
}
