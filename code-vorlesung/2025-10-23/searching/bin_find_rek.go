package searching

func BinFindRek(l []int, x int) int {
	if len(l) == 0 {
		return -1
	}

	mitte := len(l) / 2
	if x == l[mitte] {
		return mitte
	}

	if x < l[mitte] {
		return BinFindRek(l[:mitte], x)
	}

	result := BinFindRek(l[mitte+1:], x)
	if result == -1 {
		return -1
	}
	return result + mitte + 1

}
