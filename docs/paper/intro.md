# Introduction

Stream processing is a popular way to write data pipelines in modern programming languages. 
It allows programmers to express complex data pipelines in concise and readable manner while 
leaving rooms for optimizations for the compiler and the runtime.
Stream libraries typically allow users to create a stream of data from any data source.
A stream can flow into one or more downstream operators.
Each operator process one or more input streams and return one or more output streams.
Operators that do not generate immediate results, such as `Map()`, `Filter()`, and `Shuffle()`, are known as the intermediate operators.
The other operators are known as the terminal operators, such as `Count()`, `ToArray()`, and `ToFile()`. 

Some of the most widely used stream processing libraries are Java's stream, C#'s LINQ, Scala's view, Rust's Iterator, and Python's functional.
Most of these libraries will lazily execute the intermediate stages to allow the runtime to optimize the DAG to allow faster execution.
These libraries relies on them compiler to perform ahead-of-time(AOT) optimizations or language runtime to perform just-in-time(JIT) optimizations.
The primary approch to optimize the stream operations has been JIT since it can make smarter optimizations than AOT thanks to its access to runtime heuristics @scala-jit.
However, recent works argue that AOT can achieve better performance than JIT due to JIT's unperditible behavior and its necessarity to make fast 
action during runtime to achive better performance @aot.

One can view stream processing libraries as optimized and user-friendly functional programming libraries.

\prr{}
It also make use of generics that was newly introduced in \go{} 1.18 for 


