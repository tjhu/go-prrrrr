# Architecture 

## Stream 



```{.mermaid caption="Stream DAG diagram" #fig:stream}
graph LR
    S1[HTML files] --> S1a{"HtmlToWords"} 
    S1a --> M{"Unique"}
    S2[PDF files] --> S2a{"PdfToWords"}
    S2a --> M
    M --> C("Count")
```




Streams are like this @fig:stream 

* Lazy
* Unordered

## Operator Merging



## Stream Batching

