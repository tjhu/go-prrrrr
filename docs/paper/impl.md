# Implementation

## Stream Abstraction

For simplicity, we only consider the scenario where operators can only take one input stream and produce one output stream.
We also assume that the types of all operators in the same stream are the same.
Besides the intermediate and terminal stream operation methods, streams also expose methods to query the worker function and the parent to enable possible optimizations. 

## Data Passing

Unlike most other stream libraries, instead of using an iterator pattern to pass data around, \prr{} make use of \goc{}s to allow efficient data flow in the streams. 
This could be more efficient than Java stream's implementation because no virtual call is necessary to pass data through \goc{}s.
In contrary, it requires $O(N * K)$ number of virtual calls in Java to pass elements around for a stream with $N$ elements and depth of $K$ @aot.
The potential downside of using \goc{}s is that a \goc{} is basically a multiple-consumer-multiple-producer queue(mpsc queue) backed by a mutex.
If the degree of parallelism is high and the amount of work on each operator is small, the contention of the mutex could be high.
This issue could be resolved by using a more efficient mpsc queue or using the stream batching and operator batching optimizations to reduce the number of data passing required.

## Stream batching

We batch multiple elements in one slice(array) and pass the slice down the channel to reduce the number of passings.
The slices are reused at each operator to reduce the number of heap allocations.
Since the number of input elements could be much bigger than the number of output elements, we pack each slice full to maximize the effectiveness of batching.

## Operator merging

Since, in our implementation, all operators have the same numbers of input and output and the same data types, all neighboring operators can be merged into one operator.
For each neighboring operator, i.e., an operator and its parent, combine their worker function into a new one, and generate a new operator with the merged worker function.
The merging operation is done recursively until there is no more operators to merge.
