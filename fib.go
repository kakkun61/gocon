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

    fmt.Printf("%d\n", fib(n))
}

func fib (o int) int {
    switch o {
    case 0: return 0
    case 1: return 1
    }
    return fib(o-1) + fib(o-2)
}
