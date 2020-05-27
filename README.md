# pluralsight-go-functions

Deep dive into Go functions

* `defer` functions are executed in a `LIFO` fashion; `defer`s are executed when a function exits.
* `defer` functions declared before a `panic` will still be executed.