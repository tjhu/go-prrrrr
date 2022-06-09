# Evaluation

We ran our experiment on a dual-socket Dell R7525 server. Each socket runs a 32-core AMD 7542(AMD EPYC Rome) running at 2.9 GHz.
The server has 512 GB of 3200 MHz DDR4 ECC memory.
All benchmarks are done in 64 threads with `Map(int.increment)` for the workload of each operator unless specified.

## Stream Batching 

The benefit of stream batching is significant. As shown in @fig:bench-batch, batching with `batch_size=1024` provides as much as 40 times speed up.
The speed improvement is expected since most of the anticipated slow down comes from the overhead of message passing.

## Operator Merging

We are also seeing drastic improvement of applying operator merging when the number of stages is large.
For example, in @fig:bench-merge, we achieve an almost 2x speedup for a two-stage pipeline and a 6x speedup.
Again, this confirms our theory that the message passing is the main overhead in our implementation.

## Merging with Batching

For this benchmark, We ran one big and one small workload with both stream batching and operator merging optimizations enabled.
@fig:bench-small uses integer increment, a small workload, like previous benchmarks.
@fig:bench-big uses a Taylor Series for sine estimation with depth equals to 4 to estimate a relative heavy workload.
Our optimizations see performance improvement in both workload with the small workload seeing much drastic improvement.
This further confirms that message passing is the main bottle neck.
