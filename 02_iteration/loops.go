package loops

func Repeated(char string, cycles int) (result string) {
	for i := 0; i < cycles; i++ {
		result += char
	}
	return
}
