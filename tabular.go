package gotabularprint

import (
	"fmt"
	"strings"
)

type Table struct {
	Columns map[string]*Column
	order   *[]string
}

type Column struct {
	Name   string
	Length int
}

type ParsedTable struct {
	Header    string
	SubHeader string
	Format    string
}

func CreateNewTable() *Table {
	return &Table{
		Columns: map[string]*Column{},
		order:   &[]string{},
	}
}

func (tbl Table) ParseTable() *ParsedTable {
	var header string
	var subHeader string
	var format string
	var space string = "| "

	var cols = *tbl.order

	for _, c := range cols {
		cf := tbl.Columns[c].format()
		header = header + space + fmt.Sprintf(cf, tbl.Columns[c].Name)
		subHeader = subHeader + "|" + fmt.Sprintf(cf, tbl.Columns[c].underline())
		format = format + space + cf
		space = " | "
	}

	header = header + " |"
	subHeader = subHeader + "|"
	format = format + " |"

	return &ParsedTable{
		Header: header,
		SubHeader: subHeader,
		Format: format,
	}
}

func (tbl *Table) AddColumn(short, full string, length int) {
	tbl.Columns[short] = &Column{Name: full, Length: length}
	*tbl.order = append(*tbl.order, short)
}

func (tbl *Table) PrintHeaderAndGetFormat() string {
	parsedTable := tbl.ParseTable()
	fmt.Println(parsedTable.Header)
	fmt.Println(parsedTable.SubHeader)
	return parsedTable.Format
}


func (col *Column) format() string {
	pad := "-"
	return fmt.Sprintf("%%%s%dv", pad, col.Length) 
}

func (col *Column) underline() string {
	return strings.Repeat("-", col.Length + 2)
}
