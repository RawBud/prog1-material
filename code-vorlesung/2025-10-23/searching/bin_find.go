package searching

//Liefert die Position von x in der Liste l
//oder liefert -1, falls x nicht in l vorkommt.
func BinFind1(l []int, x int) int {
	links := 0
	for len(l) > 0 {
		mitte := len(l) / 2
		//Vergleiche x mit den Elemneten in der Mitte.
		//Wenn x == l[mitte], dann fertig.
		if x == l[mitte] {
			return mitte + links
		}

		//Wenn x < l[mitte], dann suche im linken Teil weiter
		if x < l[mitte] {
			//lasse nur den linken Teil der Liste übrig.
			//l = l[0:mitte]
			l = l[:mitte]
		} else {
			//lasse nur den rechten Teil der Liste übrig.
			//l = l[mitte+1 : len(l)]
			l = l[mitte+1:]
			links += mitte + 1
		}
	}

	return -1
}

func BinFind2(l []int, x int) int {
	links := 0
	rechts := len(l)

	for rechts >= links {
		mitte := (rechts - links) / 2

		if l[mitte] == x {
			return mitte
		}

		if x < l[mitte] {
			rechts = mitte
		} else {
			links = mitte + 1
		}
	}

	return -1
}
