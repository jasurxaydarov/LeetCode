package main

import (
    "fmt"
    "runtime"
    "sync"
)

func printNumber(num int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println(num)
}

func main() {
    runtime.GOMAXPROCS(4) // 4 ta yadrodan foydalanish uchun
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go printNumber(i, &wg)
    }

    wg.Wait()
}
