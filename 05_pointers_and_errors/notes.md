* go treats functions argumnets as values, not references, so you need pointers to mutate state over an object (a method of a struct needs a pointer to mutate its own state)
* nil is the `null` in go, pointers can be null as well as any struct, err and so on, so need to check for nils for avoiding runtime errs
* you can create new types from existing ones via `type NewType oldType`
