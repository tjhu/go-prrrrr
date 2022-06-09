# Conclusions

While it is optimal to have the compilers, whether static, AOT, or JIT, to optimize the stream operations, it's possible for library implementers to improve the performance with stream optimizations.
\prr{} is a proof-of-concept of implementing DAG optimizations on the library side.
It implements stream batching and operator merging optimizations and evaluate the performance of the optimized version versus the performance of the un-optimized version.
It confirms our theory about the main cost of parallel stream processing is the inter-process communication(IPC) cost.

The source code of this paper and \prr{} is obtainable on [https://github.com/tjhu/go-prrrrr](https://github.com/tjhu/go-prrrrr).

## Acknowledgements

I thank my friends, especially Zhaofeng and Hongyu, for providing me moral and technical support during the grind.