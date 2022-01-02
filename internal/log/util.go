package log

func nearestMultiple(j, k uint64) uint64 {
	if j >= 0 {
		return j / k * k
	}

	return (j - k + 1) / k * k
}
