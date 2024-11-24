package testcover

func Add(x, y int32) int32 {
	if x < 0 {
		return -1
	}

	if y < 0 {
		return -2
	}

	return x + y
}