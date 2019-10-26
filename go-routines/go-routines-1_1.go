
package main

import ( "fmt"
         "sync" 
)

var wg = sync.WaitGroup{}

func print(){
    for i:=0; i<10; i++  {
        fmt.Println(i)
    }
    wg.Done()
}

func main() {
    wg.Add(1)    
    go print()
    wg.Wait()
    
}
