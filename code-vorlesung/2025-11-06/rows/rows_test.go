package rows

import "fmt"

type Row []string

func (r Row) ContainsOnly(s string) any {
	panic("unimplemented")
}

// New liefert eine neue Zeile der Länge `length` für
// ein Spielfeld und füllt sie mit `fill`.
func New(length int, fill string) Row {
	result := make(Row, length) // Erzeugt eine Row (Liste von Strings) mit der angegebenen Länge.

	//for i := 0; i < length; i++{
	for i := range result {
		//result = append(result, fill)
		result[i] = fill
	}

	return result
}

func ContainsOnly(r Row, z string) bool {
	for i := 0; i < len(r); i++ {
		if r[i] != z {
			return false
		}
	}
	return true
}

func ExampleRow_ContainsOnly() {
	r1 := New(3, "+")

	fmt.Println(r1)
	fmt.Println(r1.ContainsOnly("+"))

}
