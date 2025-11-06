package rows

type Row []string

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
