# Evaluation

We ran our experiment on a dual-socket Dell R7525 server. Each socket runs a 32-core AMD 7542(AMD EPYC Rome) running at 2.9 GHz.
The server has 512 GB of 3200 MHz DDR4 ECC memory.

## Stream Batching 

The benefit of stream batching is significant. As shown in @fig:bench-batch, batching with `batch_size=1024` provides as much as 40 times speed up.
The speed improvement is expected since most of the anticipated slow down comes from the overhead of message passing.

## Operator Merging
