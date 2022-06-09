# Architecture 

## Stream 

A stream is a collection of data of same type. 
Each operator can process one or more input streams and return one or more output streams.
Combining various operators and streams will result in a acyclic-directed-graph(DAG), such as @fig:stream.
Operators that do not generate immediate results, such as `Map()`, `Filter()`, and `Shuffle()`, are known as the intermediate operators.
The other operators are known as the terminal operators, such as `Count()`, `ToArray()`, and `ToFile()`.
Intermediate operators are lazily executed, meaning that they will not be executed until it is invoked by a terminal operator.
Lazy evaluation allows the runtime to perform optimizations at the terminal operators.
The ordering of the data in the stream is not guaranteed to allow further optimization.
Since each operator could run on a different thread, it's important to reduce the number of message passing between the operators since cross thread communication could be very costly.

```{.mermaid caption="Stream DAG diagram" #fig:stream}
graph LR
    S1[HTML files] --> S1a{"HtmlToWords"} 
    S1a --> M{"Unique"}
    S2[PDF files] --> S2a{"PdfToWords"}
    S2a --> M
    M --> C("Count")
    M --> F("ToFile")
```

## Operator Merging

If two neighboring operators has the same number of input streams and output streams and the same type of streams, 
it's possible to create a new operator that combines the functionalities of the original two operators to reduce the number element passings.
For example, in @fig:merge, operator `A` and `B` are merged into operator `AB` and the number of message passing is reduced from 3 to 2.

```{.mermaid caption="Operator merging" #fig:merge}
graph LR
    S1[Source] --> A --> B --> T1[Terminal]

    S2[Source] --> AB --> T2[Terminal]
```

## Stream Batching

Instead of passing each individual element through the stream, each operator can pass a batch of elements at a time to its downstream.
As shown in @fig:batch, stream batching can reduce the number of message passing to how many times the batch size.

```{.mermaid caption="Stream batching" #fig:batch}
graph LR
    S1[Source] --> A1[A] --> B1[B] --> T1[Terminal]

    S2[Source] --> A2[A] --> B2[B] --> T2[Terminal]
    A2[A] --> B2[B]
    A2[A] --> B2[B]
    A2[A] --> B2[B]
    A2[A] --> B2[B]
```
