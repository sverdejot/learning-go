* context helps us cancel goroutines whenever needed
    * cancellation token in C#
* google guideline is to providede every downstream function the context as first parameter to help cancel further execution when one inform of an error or cancelation
