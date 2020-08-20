#Package hybridsort provides an example of stable, concurrent sorting algorithm based on timsort and symmerge algorithms.
The procedure works in a few steps:
- find pre-sorted chunks (called 'run') of size [MinRunSize, MaxRunSize] and run insertion sort in parallel
- push results into priority queue to restore sequence of runs
- fetch runs from priority queue and perform symmerge in parallel
- push back merged results into priority queue util the only one run found in the queue

#A number of settings are available to tune procedure:
- min and max run size
- degree of parallelism

NOTE: symmerge implementation politely borrowed from https://golang.org/src/sort/sort.go
