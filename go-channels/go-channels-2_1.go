
package main

import ( "fmt" 
    "sync" 
)

var wg = sync.WaitGroup{}

func main() {
    i := 10
    var ch = make(chan int)
    
    wg.Add(2)
    go func(ch chan<- int) {
        ch <- i
        // <-ch //error
        wg.Done()
    }(ch)
    go func(ch <-chan int) {
        i := <-ch 
        fmt.Println(i)
        // ch <- i //error
        wg.Done()
    }(ch)
    wg.Wait()
    
}
