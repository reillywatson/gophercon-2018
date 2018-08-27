testing workshop
----------------

github.com/matryer/moq — generates mock implementations of interfaces, with FooFunc() members. We have a few places where we use this pattern, this looks handy for 

major takeaways:
- our tests are testing too many things! our integration tests are testing the entire stack, which makes it harder to identify where the problem lies if there’s a failure
- we need to be spending more time refactoring tests for readability and maintainability
- we should be making waaaay more test helper functions, many of our tests are way too long to be readable
- we have a lot of functions that are annoying to test because you have to mock out so many things. This is an indication those functions are doing too much!
- related: we should be doing more of passing dependencies into functions, rather than having them as global variables that get overwritten in tests. Our database access situation is pretty messy, I think we should be making more interfaces for this!
	- we could use a runner struct that contains all your dependencies instead of passing them in as a million arguments
- interface testing: write a test suite that verifies implementations of some interface conform to the contract of that interface. This would be really useful for places where we have many implementations of an interface! (ie the Benefit interfaces, for instance)

interesting pattern for using multiple assignment to condense checks:
```go
func TestAddTen_Best(t *testing.T) {
	if got, exp := math.AddTen(1), 11; got != exp {
		t.Fatalf("unexpected value, got: %d, exp %d", got, exp)
	}
}
```

Useful check that we aren’t really using:
```go
	if testing.Verbose() {
		t.Log("put extra logging here that normally we don't care about")
	} else {
		// silence my normal loggers
		log.SetOutput(ioutil.Discard)
	}
```

We should be running our tests with “-timeout” so CI can give a more useful error (and return in a more reasonable time frame) if a test deadlocks.

We should be running our tests with “-race” in CI! I’ve been thinking about this for a while, it’s time to pull the trigger.


