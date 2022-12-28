package coalesce

// Coalesce[T] - Return first non-nil object
func Coalesce[T any](args ...*T) *T {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
