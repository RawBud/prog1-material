package searching

import "fmt"

func ExampleBinFindRek() {
	l1 := []int{1, 6, 8, 10, 25, 125, 277}

	fmt.Println(BinFindRek(l1, 1))
	fmt.Println(BinFindRek(l1, 277))
	fmt.Println(BinFindRek(l1, 25))
	fmt.Println(BinFindRek(l1, 10))

	// Output:
	// 0
	// 7
	// 4
	// 3
}
