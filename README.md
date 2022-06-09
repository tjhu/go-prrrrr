# go-parallelstream

[![.github/workflows/test.yml](https://github.com/tjhu/go-parallelstream/actions/workflows/test.yml/badge.svg)](https://github.com/tjhu/go-parallelstream/actions/workflows/test.yml)

## How to run

To test

```bash
go test ./stream
```

To benchmark

```bash
go test ./benchmark -bench=. -cpu 1 
```

## How to build 

Install pandoc and [mermaid filter](https://github.com/raghur/mermaid-filter). Then run

```
make
```

## Known issues

* This repo is full of hacks because of the tight deadline. Although there're some cool stuff here, don't take the architecture of this codebase too seriously.
* Forking a stream(a list instead of a DAG) is not supported.
* The source stream will be modified.
