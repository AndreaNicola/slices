package slices

func Filter[T any](filter func(t T) bool, ts []T) []T {

	filteredTs := make([]T, len(ts))
	lastFilteredT := 0

	for i := 0; i < len(ts); i++ {

		if filter(ts[i]) {
			filteredTs[lastFilteredT] = ts[i]
			lastFilteredT = lastFilteredT + 1
		}

	}

	recapped := make([]T, lastFilteredT)
	copy(recapped, filteredTs)

	return recapped

}
