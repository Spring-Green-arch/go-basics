
package main

import ( "fmt" 
    "sync" 
)

var wg = sync.WaitGroup{}

func main() {
    ch := make(chan int, 2)
    
    wg.Add(2)
    go func(ch <-chan int) {
        for i := range ch {
            fmt.Println(i)
        }
        wg.Done()
    }(ch)
    go func(ch chan<- int) {
        ch <- 10
        ch <- 20
        close(ch)
        wg.Done()
    }(ch)
    wg.Wait()
}
