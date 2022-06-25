package helpers

//	A simple reducer callback signature
type Reducer[T, R any] func(accumulator R, current T) R

//	Reduce a slice to a singular value
func reduce[T, R any](slice []T, reducer Reducer[T, R], initializer R) R {
	accumulator := initializer
	for _, value := range slice {
		accumulator = reducer(accumulator, value)
	}
	return accumulator
}

//	Reduce a slice to a singular value in the opposite order
func reduceRight[T, R any](s []T, reducer Reducer[T, R], initializer R) R {
	reverse(s)
	return reduce(s, reducer, initializer)
}

//	Reverse a slice
func reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

}
