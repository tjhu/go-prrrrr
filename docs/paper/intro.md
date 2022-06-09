# Introduction

Stream processing is a popular way to write data pipelines in modern programming languages. 
It allows programmers to express complex data pipelines in concise and readable manner while 
leaving rooms for optimizations for the compiler and the runtime.
Stream libraries typically allow users to create a stream of data from any data source.
A stream can flow into one or more downstream operators.
Programmers can chain an arbitrary number of streams and operators together to form complex data pipelines. 

Some of the most widely used stream processing libraries are Java's stream, C#'s LINQ, Scala's view, Rust's Iterator, and Python's functional.
Most of these libraries will lazily execute the intermediate stages to allow the runtime to optimize the DAG to allow faster execution.
These libraries relies on them compiler to perform ahead-of-time(AOT) optimizations or language runtime to perform just-in-time(JIT) optimizations.
The primary approch to optimize the stream operations has been JIT since it can make smarter optimizations than AOT thanks to its access to runtime heuristics @scala-jit.
However, recent works argue that AOT can achieve better performance than JIT due to JIT's unperditible behavior and its necessarity to make fast 
action during runtime to achive better performance @aot.

Having access to the syntax tree is important for optimizing streams. However, it is possible to optimize the streams without the compiler support.
If the operators do not produce side effects, one can view these stream processing libraries as optimized and user-friendly functional programming libraries.
During the era of functional programming, researchers did extensive research on stream processing optimizitions @fusion-nothing @fusion-complete @parallel-haskell @fusion-catalog. 
\prr{} experiments with a few stream processing optimizations in \go{} and measures the performance improvement of them.

