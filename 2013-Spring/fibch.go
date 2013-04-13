package main

import (
    "fmt"
    "os"
    "strconv"
)

func main () {
    args := os.Args
    n, e := strconv.Atoi(args[1])
    if e != nil {
        fmt.Printf("%v\n", e)
        os.Exit(1)
    }

    ch := make(chan int)
    go fib(n, ch)
    fmt.Printf("%d\n", (<- ch))
}

func fib (o int, ch chan int) {
    switch o {
    case 0: ch <- 0
    case 1: ch <- 1
    }
    ch1, ch2 := make(chan int), make(chan int)
    go fib(o-1, ch1)
    go fib(o-2, ch2)
    ch <- (<- ch1) + (<- ch2)
}

