##  并发模式 (concurent pattern)
It is very important to keep in mind that concurrency is about structure and parallelism is about execution.
We must think about making our programs concurrent in a better way, by breaking them down into smaller pieces of work,
and Go's scheduler will try to make them parallel if it's possible and allowed.

