package main

import (
	"log"
	parquet "root/Parquet"
)

type Record struct {
	ID    int64
	Name  string
	Value float64
}

func main() {
	schema := []*parquet.SchemaElement{
		{Name: "id", Type: 1},
		{Name: "name", Type: 2},
		{Name: "value", Type: 3},
	}

	writer, err := parquet.NewParquetWriter("example.parquet", schema)
	if err != nil {
		panic(err)
	}

	records := []Record{
		{ID: 1, Name: "Alice", Value: 12.5},
		{ID: 2, Name: "Bob", Value: 23.7},
		{ID: 3, Name: "Charlie", Value: 34.9},
	}

	if err := writer.WriteRowGroup(records); err != nil {
		log.Println(err)
	}

	if err := writer.Close(); err != nil {
		log.Println(err)
	}
}
