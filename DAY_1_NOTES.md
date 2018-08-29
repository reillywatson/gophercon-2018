go scheduler saga: talked about how the Go scheduler is implemented. Interesting stuff!
	details:
	- map goroutines onto $NUM_CPUS threads
	- maintain a separate run queue for each thread
	- when you make a blocking system call (for instance any I/O), read/write on a channel, or sleep, this calls into the scheduler

macaroons: like cookies but with more layers

machine learning on go code: seems pretty immature still, he was able to do things like predict the next token in source code 60% of the time

how to structure your Go apps:
proposed using domain-driven design and a hexagonal architecture

some interesting proposals for go 2 around error handling and generics!

rethinking asynchronous patterns:
since it's trivial in Go to make an async wrapper around a synchronous API (or vice-versa), it makes sense to make your APIs synchronous!

