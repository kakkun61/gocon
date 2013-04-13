Fibonacci
=========

fib.go
------

単純なフィボナッチ数関数。メモ化も何もなし。

A simple Fibonacci number function. no memoize.

fibch.go
--------

fib.go のフィボナッチ関数をゴルチン化してみた。

Goroutine-izing the fibonacci number function of fib.go

結果 Result
----

    $ time go run fibch.go 9
    34
    
    real	0m8.710s
    user	0m0.492s
    sys	0m8.085s
				
    $ time go run fib.go 9
    34

    real	0m0.257s
    user	0m0.216s
    sys	0m0.032s


