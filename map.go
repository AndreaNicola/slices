package slices

func Map[T1 any, T2 any](conversion func(t1 T1) T2, t1s []T1) []T2 {

	t1sCardinality := len(t1s)
	t2s := make([]T2, t1sCardinality)
	for i := 0; i < t1sCardinality; i++ {
		t2s[i] = conversion(t1s[i])
	}

	return t2s

}
