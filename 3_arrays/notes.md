* we can loop over the elements of an array via its index or using range <array>, which does the same as python's `enumerate` method.
    * the annonymous operator is the same as in python, `_`
* for typing collections of elements we have two types, `arrays` and `slices`, which are the same except for:
    * arrays have fixed lenght, declared when declaring the variable, and which is incldued in the type:
        * var chain [5]string{"link", "link", "link", "link", "link"}
    * slices are the same as arrays but do not declare fixed length
        * var chain []string{"link", "link", "link"}
* go has a built-in tool to get the test coverage `go test -cover`
* go has built-in slices standard library for comparing slices
* `make` allows to create a slices with a starting capacity of the length provided
* slices can be sliced just like python with [<slicing index>:]
