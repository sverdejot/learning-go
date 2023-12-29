* sync libraries allows to handle synchronization between goroutines, with semaphores, mutual exclusion, wait...
* to properly handling waiting for goroutines to finish, we can use Mutex or channels
    * use channels when passing ownership of data
    * use mutex for managing stat
* we can embed a type in our struct, but that will expose public method of the embedded type to our struct, so it is not convenient in most cases, prefer to add a property of the embedded type and handle its state internally
