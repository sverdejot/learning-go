* in go, we bind methods to types (also called `structs`) by defining it via:
    `func (receiverName ReceiverType) MethodName(methodArgs) {}`
    thus, we bind MethodName to ReceiverType, being able to reference fields of ReceiverType inside MethodName by receiverName parameter
* in go, interfaces resolution is implicit rather than explcit, thus, we only need to declare the methods of the interfaces on the selected type without expliciting declaring it
* you can also define anonymous structs in inline types for, for example, slices
