package iteration

func Repeat(character string, numRepeats int) string {
	repeated := ""

	for i := 0; i < numRepeats; i++ {
		repeated += character
	}

	return repeated
}
