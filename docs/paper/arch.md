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

```{.mermaid caption="Stream DAG diagram" #fig:stream}
graph LR
    S1[Source] --> A --> B --> T1[Terminal]

    S2[Source] --> AB' --> T2[Terminal]
```



## Stream Batching

```{.mermaid caption="Stream DAG diagram" #fig:stream}
graph LR
    S1[Source] --> A1[A] --> B1[B] --> T1[Terminal]

    S2[Source] --> A2[A] --> B2[B] --> T2[Terminal]
    A2[A] --> B2[B]
    A2[A] --> B2[B]
    A2[A] --> B2[B]
    A2[A] --> B2[B]
```