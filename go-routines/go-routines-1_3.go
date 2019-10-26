
package main

import ( "fmt"
         "sync" 
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func print(i int){
    fmt.Println(i)
    m.RUnlock()
    wg.Done()
}

func incrementI(i *int) {
    (*i)++
    m.Unlock()
}

func main() {
    for i:=0; i<10;  {
        wg.Add(1)    
        m.RLock()
        go print(i)
        m.Lock()
        incrementI(&i)
    }
    wg.Wait()
}
