package fizzbuzz

import "fmt"

//fizz gibt die Zahlen von 1 bis 30 aus und ersetzt dabei
//jede durch 3 teilbare Zahl durch "fizz" und jede durch
//5 teilbare Zahl durch "buzz"
//Bei Zahlen, die sowohl durch 3 als auch durch 5 tielbar sind,
//wird "fizzbuzz" ausgegeben

func FizzCompact() {
	for i := 1; i <= 30; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}
		if i%3 == 0 {
			fmt.Print("fizz")
		}
		if i%5 == 0 {
			fmt.Print("buzz")
		}
		fmt.Println()
	}
}
