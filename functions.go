package slices

func Map[T1 any, T2 any](conversion func(t1 T1) T2, t1s []T1) []T2 {

	t1sCardinality := len(t1s)
	t2s := make([]T2, t1sCardinality)
	for i := 0; i < t1sCardinality; i++ {
		t2s[i] = conversion(t1s[i])
	}

	return t2s

}

func GoMap[T1 any, T2 any](conversion func(t1 T1) T2, t1s []T1) []T2 {

	t1sCardinality := len(t1s)
	t2s := make([]T2, t1sCardinality)
	for i := 0; i < t1sCardinality; i++ {
		go func(j int) {
			t2s[j] = conversion(t1s[j])
		}(i)
	}

	return t2s

}

func Reverse[T any](ts []T) []T {

	rev := make([]T, len(ts))
	for i := 0; i < len(ts); i++ {
		rev[len(ts)-1-i] = ts[i]
	}

	return rev

}

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

func foldr[T1, T2 any](f func(t1 T1, t2 T2) T2, t T2, ts []T1) T2 {

	if len(ts) == 0 {
		return t
	}

	return f(ts[0], foldr(f, t, ts[1:]))

}

func foldl[T1, T2 any](f func(t1 T1, t2 T2) T2, t T2, ts []T1) T2 {

	if len(ts) == 0 {
		return t
	}

	return foldl(f, f(ts[0], t), ts[1:])

}
