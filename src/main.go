package main

import (
	"fmt"
	"root/src/convertor"
)

type Record struct {
	ID    int
	Name  string
	Value float64
}

func main() {
	record := Record{ID: 12345, Name: "John", Value: 99.99}

	schema := []convertor.SchemaElement{
		{Name: "ID", Type: convertor.TypeInt},       //64 bite
		{Name: "Name", Type: convertor.TypeString},  // offset (16 bite) + content
		{Name: "Value", Type: convertor.TypeDouble}, // 64 bite
	}

	serializedData, err := convertor.SerializeFromStruct(record)
	if err != nil {
		fmt.Println("Error serializing:", err)
		return
	}

	fmt.Println(serializedData)

	deserializedValues, err := convertor.DeserializeFromSchema(serializedData, schema)
	if err != nil {
		fmt.Println("Error deserializing:", err)
		return
	}

	fmt.Println("Deserialized Values:")
	for _, value := range deserializedValues {
		fmt.Println(value)
	}

	deserializedValuesBinary, err := convertor.DeserializeFromSchemaBinary(serializedData, schema)
	if err != nil {
		fmt.Println("Error deserializing:", err)
		return
	}

	fmt.Println("Deserialized byte Values:")
	for _, value := range deserializedValuesBinary {
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
