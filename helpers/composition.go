package helpers

//	Function that takes in a string and returns a modified string
type StringModifierFn func(s string) string

//	Pipes the given functions
func Pipe(fns ...StringModifierFn) StringModifierFn {
	return func(s string) string {
		return reduce(fns, func(accumulator string, current StringModifierFn) string {
			return current(accumulator)
		}, s)
	}
}

//	Composes the given functions
func Compose(fns ...StringModifierFn) StringModifierFn {
	return func(s string) string {
		return reduceRight(fns, func(accumulator string, current StringModifierFn) string {
			return current(accumulator)
		}, s)
	}
}
