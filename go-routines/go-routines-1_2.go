
package main

import ( "fmt"
         "sync" 
)

var wg = sync.WaitGroup{}

func print(i int){
    fmt.Println(i)
    wg.Done()
}

func main() {
    for i:=0; i<10; i++ {
        wg.Add(1)    
        go print(i)
    }
    wg.Wait()
}
