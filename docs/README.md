# Go Parallel Stream Library

## Overview

Combine the expressiveness of Java stream and the beauty(?) of Go routine to create something cool(efficent yet simple?).

## In the wild

* Java stream
* Scala view


## Directions

* Push/pull based:
  * Because the difference in parallelism between operators, we have to do push based.
* Fuse operations


## Showcase

* Short circuit: if limit is set and reached, we can early terminate.

## Why mine's better

### Java stream virtual call overhead

Java's stream implementation roughly translate to this code:

```java
while (it.hasNext()) consumer.accept(it.next());
```

This is 3 virtual calls per element per operator.
Collections in Java's standard library can have just 1 instead. 
With a pipeline with N elements and depth K, we can have up to O(N * K) virtual calls. _Cite AOT_

Our implementation in Go is faster because no virtual call when pushing the element to the next operator.
However, we use the go channel which is implemented with a mutex. We could potential switch to a lock free
mpmc queue if necessary.




## Usage

```go
  pstream.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
    .MaxWorker(10)
    .Filter(isPrime)
    .ForEach(func(n int) {
      fmt.Printf("%d is a prime number\n", n)
    })
```

## Workloads

* power series of e^x.
* newton's algorithm for square root.
* PI square over six series calculation.


## Reading list

* [Eliminating Abstraction Overhead of Java Stream Pipelines using Ahead-of-Time Program Optimization](https://dl.acm.org/doi/pdf/10.1145/3428236) @ PACMPL'20 by Aarhus Univerisity
* [Stream Fusion: From Lists to Streams to Nothing at All](https://dl.acm.org/doi/pdf/10.1145/1291220.1291199) @ ICFP'07 by Oxford
* [A Catalog of Stream Processing Optimizations](https://dl.acm.org/doi/pdf/10.1145/2528412) @ ACM Computer Survey 2014 by IBM Watson
* [Tutorial: Stream Processing Optimizations](https://dl.acm.org/doi/pdf/10.1145/2488222.2488268) @ DEBS'13 by IBM Watson
* [Expressive and Efficient Streaming Libraries](http://biboudis.github.io/papers/dissertation.pdf) @ Athens Ph.D. thesis.

## Relevant work

* [BWbwchen/MapReduce](https://github.com/BWbwchen/MapReduce): Actual map reduce. Support distributed workers, fault tolerant, and states are presisted.
* [logic-building/functional-go](https://github.com/logic-building/functional-go): simiar similar to go-funk.
* [mariomac/gostream](https://github.com/mariomac/gostream): well written, with template, Java stream API.
* [thoas/go-funk](https://github.com/thoas/go-funk): functional programming lib.
* [reugn/go-streams](https://github.com/reugn/go-streams): data pipeline library. Kinda of map reduce.
* [robpike/filter](https://github.com/robpike/filter): apply and reduce.
* [samber/lo](https://github.com/samber/lo): go-funk but support parallel. The parallel impl simply applys go routine for each element. `WaitGroup` is used to synchronize between stages.

### Amazing relevant work

* [jucardi/go-streams](https://github.com/jucardi/go-streams): Lazy Java stream like. Able to controll parallelism.
