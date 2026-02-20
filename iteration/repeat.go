package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteSstring(character)
	}
	return repeated.String()
}
