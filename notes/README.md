# Go Parallel Stream Library

## In the wild

* Java stream
* Scala view


## Directions

* Push/pull based:
* Fuse operations



## Usage

```go
  pstream.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
    .MaxWorker(10)
    .Filter(isPrime)
    .ForEach(func(n int) {
      fmt.Printf("%d is a prime number\n", n)
    })
```


## Reading list

* [Eliminating Abstraction Overhead of Java Stream Pipelines using Ahead-of-Time Program Optimization](https://cs.au.dk/~amoeller/papers/streamliner/paper.pdf)
* [Stream Fusion: From Lists to Streams to Nothing at All](https://dl.acm.org/doi/pdf/10.1145/1291220.1291199)
* [A Catalog of Stream Processing Optimizations](https://dl.acm.org/doi/pdf/10.1145/2528412)

## Relevant work

* [samber/lo](https://github.com/samber/lo): simply applys go routine for each element. `WaitGroup` is used to synchronize between stages.
* 