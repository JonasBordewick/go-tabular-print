package main

import(
	"fmt"
	tabular "github.com/JonasBordewick/go-tabular-print"
)

func main() {
	table := tabular.CreateNewTable()

	table.AddColumn("t", "Test", 37)
	table.AddColumn("e", "Test 2", 10)
	table.AddColumn("s", "Test 3", 20)

	var data = []struct {
		t, e string
		s    int
	}{
		struct {
			t string
			e string
			s int
		}{
			t: "Hallo Welt",
			e: "Wie gehts",
			s: 1749172954719825,
		},
	}

	format := table.PrintHeaderAndGetFormat()
	for _, row := range data {
		fmt.Printf(format, row.t, row.e, row.s)
	}

}