writing accessible go:
	- group logically-related blocks of functionality, keep variables, interfaces, types, etc close to where they're used -- requires less context to keep in your head (text-to-speech is a lot slower than reading so the less context the better)
	- keep names short (ie prefer `var a, b Vector` over `var vectorA, vectorB Vector`) -- fast to listen to, easier to navigate, less effort to type
	- make names meaningful (ie prefer `var total, scaled Vector` over `var tVec, sVec Vector` -- reduces cognitive load
	- use pronounceable names -- way easier on screen readers. Note: concatenating words can be tricky (ie Github being pronounced "geethub")
	- use new lines intentionally, like you would paragraph breaks. Requires thoughtfulness about not overusing or underusing.
	- be consistent in style. If you change style, update it everywhere! -- just good life advice
		- semi-related: We should write down some of the more-frequent things that come up in code reviews as a set of best practices. I've been trying to enforce as much of it through linters as I can, but that can't handle everything!

going serverless:
	- google cloud functions have Go support in alpha, we should ask for access!
	- need to add tracing manually for each function
	- cold start is milliseconds up to a couple seconds, warm start is "fast"
	- configuration is a bit of a headache, we'd probably want to use environment variables
	- key consideration: we'd need to limit the default size of the MongoDB connection pool, we don't want it to preallocate a bunch of sockets for each separate function!
	- probably better to add functions for net-new stuff, rather than trying to rewrite our whole app in this style
	- would need to either redo our API as REST (not terribly hard), or write a weird shim that calls out to functions and forwards responses (easier)
	- unclear to me how URL discovery works for the functions

Go in Debian:
	- their code search is neat, uses https://swtch.com/~rsc/regexp/regexp4.html
	- `pk4 <pkgname>` to get the source, `pk4-replace <pkgname>` to modify the source and rebuild. Handy!

Specifications and MongoDB:
	- guy writing the Go MongoDB driver talked about how specs are complicated! Specs are incomplete or wrong, reference implementations are incorrect, etc.
	- Took him a year to implement the BSON spec, which seems nuts to me, that spec does not seem very complicated.
	- basically: expect everything to be hard.

Adventures in cgo performance:
	- calling C from Go is a bit slow, ~170ns per call
	- calling Go from C is hella slow, 1-5ms per call
	- C can't hold Go pointers, in case Go wants to have its garbage collector relocate memory
		- workaround: use a map[uintptr]interface{}, pretend those are pointers?
