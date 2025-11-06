package searching

import "fmt"

func ExampleBinFind1() {
	l1 := []int{1, 6, 8, 10, 25, 125, 277}

	//println(BinFind1(l1, 3))
	fmt.Println(BinFind1(l1, 1))
	fmt.Println(BinFind1(l1, 277))
	fmt.Println(BinFind1(l1, 25))
	fmt.Println(BinFind1(l1, 10))

	// Output:
	// 0
	// 6
	// 4
	// 3
}

func ExampleBinFind2() {
	l1 := []int{1, 6, 8, 10, 25, 125, 277}

	fmt.Println(BinFind2(l1, 1))
	fmt.Println(BinFind2(l1, 277))
	fmt.Println(BinFind2(l1, 25))
	fmt.Println(BinFind2(l1, 10))

	// Output:
	// 0
	// 7
	// 4
	// 3
}
