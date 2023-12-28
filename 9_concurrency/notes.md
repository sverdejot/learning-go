* concurrency in go in declared by adding the keyword *go* in front of any function call (can be annonymous call)
    * take care with annoymous function parameters and using references from the outer functions or loops
* most of the times, we will face race conditions, so we will need to use channels to store the results of the goroutines
    * channels are the way of communicating threads in go and storing results of goroutines

