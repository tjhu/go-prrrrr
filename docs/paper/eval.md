# Evaluation

We ran our experiment on a dual-socket Dell R7525 server. Each socket runs a 32-core AMD 7542(AMD EPYC Rome) running at 2.9 GHz.
The server has 512 GB of 3200 MHz DDR4 ECC memory.
All benchmarks are done in 64 threads with `Map(int.increment)` for the workload of each operator unless specified.
All numbers are reported in milliseconds.

## Stream Batching 

Optimization/Batch Size | 1 | 32 | 64 | 512 | 1024
-----|-----|-----|-----|-----|-----:
Un-optimized|1847|N/A|N/A|N/A|N/A
Stream Batching|N/A|104|72|45|42

Table: Stream batching benchmarks. {#tbl:bench-batch}

The benefit of stream batching is significant. As shown in @tbl:bench-batch, batching with `batch_size=1024` provides as much as 40 times speed up.
The speed improvement is expected since most of the anticipated slow down comes from the overhead of message passing.

## Operator Merging

Optimization/Depth | 1 | 2 | 4 | 8 | 16
-----|-----|-----|-----|-----|-----:
Un-optimized|3503|5327|9168|16164|30321
Operator Merging|3518|3596|3671|3813|4424

Table: Operator merging benchmarks. {#tbl:bench-merge}

We are also seeing drastic improvement of applying operator merging when the number of stages is large.
For example, in @tbl:bench-merge, we achieve an almost 2x speedup for a two-stage pipeline and a 6x speedup.
Again, this confirms our theory that the message passing is the main overhead in our implementation.

## Merging with Batching

Optimization/Workload | Light | Heavy
-----|-----|-----:
Un-optimized|16265|5327|
Batching+Merging|23930|10208|

Table: Small workload benchmarks. {#tbl:bench-all}

For this benchmark, We ran one big and one small workload with both stream batching and operator merging optimizations enabled.
And we set the number of stages of the stream to be 8 and the batch size to be 1024.
@tbl:bench-all uses integer increment, a small workload, like previous benchmarks.
@tbl:bench-all uses a Taylor Series for sine estimation with depth equals to 4 to estimate a relative heavy workload.
Our optimizations see performance improvement in both workload with the small workload seeing much drastic improvement.
This further confirms that message passing is the main bottle neck.
