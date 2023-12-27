* maps do not need to be referenced in functions like pointers since they are pointers itself
* it is useful to create custom types for maps
* maps can take any type as value (including structs and interfaces), but only comparable types as keys
* do not create empty maps using var since it will create a pointer to null, use make instead or declared it with empty values
* to create custom error it is only needed to implement Error() interface, and then we can declared them as constants
