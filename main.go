package main

import (
	"fmt"
	"root/convertor"
)

type Record struct {
	ID    int
	Name  string
	Value float64
}

func main() {
	record := Record{ID: 12345, Name: "John", Value: 99.99}

	schema := []convertor.SchemaElement{
		{Name: "ID", Type: convertor.TypeInt},
		{Name: "Name", Type: convertor.TypeString},
		{Name: "Value", Type: convertor.TypeDouble},
	}

	serializedData, err := convertor.SerializeFromStruct(record)
	if err != nil {
		fmt.Println("Error serializing:", err)
		return
	}

	deserializedValues, err := convertor.DeserializeFromSchema(serializedData, schema)
	if err != nil {
		fmt.Println("Error deserializing:", err)
		return
	}

	fmt.Println("Deserialized Values:")
	for _, value := range deserializedValues {
		fmt.Println(value)
	}

	// schema := []*parquet.SchemaElement{
	// 	{Name: "id", Type: 1},
	// 	{Name: "name", Type: 2},
	// 	{Name: "value", Type: 3},
	// }

	// writer, err := parquet.NewParquetWriter("example.parquet", schema)
	// if err != nil {
	// 	panic(err)
	// }

	// if err := writer.WriteRowGroup(records); err != nil {
	// 	log.Println(err)
	// }

	// if err := writer.Close(); err != nil {
	// 	log.Println(err)
	// }
}
