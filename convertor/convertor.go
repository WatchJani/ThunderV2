package convertor

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

const (
	TypeInt    = 1
	TypeString = 2
	TypeDouble = 3
)

type SchemaElement struct {
	Name string
	Type int
}

func SerializeFromStruct(obj interface{}) ([]byte, error) {
	var buf bytes.Buffer
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %s", t.Kind())
	}

	for i := 0; i < t.NumField(); i++ {
		value := v.Field(i)

		switch value.Kind() {
		case reflect.Int:
			if err := binary.Write(&buf, binary.LittleEndian, int64(value.Int())); err != nil {
				return nil, err
			}
		case reflect.String:
			str := value.String()
			if err := binary.Write(&buf, binary.LittleEndian, int16(len(str))); err != nil {
				return nil, err
			}
			buf.WriteString(str)
		case reflect.Float64:
			if err := binary.Write(&buf, binary.LittleEndian, value.Float()); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported field type: %s", value.Kind())
		}
	}

	return buf.Bytes(), nil
}

func DeserializeFromSchema(data []byte, schema []SchemaElement) ([]string, error) {
	var result []string
	buf := bytes.NewReader(data)

	for _, element := range schema {
		switch element.Type {
		case TypeInt:
			var value int64
			if err := binary.Read(buf, binary.LittleEndian, &value); err != nil {
				return nil, err
			}
			result = append(result, fmt.Sprintf("%d", value))

		case TypeString:
			var length int16
			if err := binary.Read(buf, binary.LittleEndian, &length); err != nil {
				return nil, err
			}

			strBytes := make([]byte, length)
			if err := binary.Read(buf, binary.LittleEndian, &strBytes); err != nil {
				return nil, err
			}
			result = append(result, string(strBytes))

		case TypeDouble:
			var value float64
			if err := binary.Read(buf, binary.LittleEndian, &value); err != nil {
				return nil, err
			}
			result = append(result, fmt.Sprintf("%f", value))

		default:
			return nil, fmt.Errorf("unsupported field type: %d", element.Type)
		}
	}

	return result, nil
}

func DeserializeFromSchemaBinary(data []byte, schema []SchemaElement) ([][]byte, error) {
	var (
		result [][]byte
		offset uint16
	)

	for _, element := range schema {
		switch element.Type {
		case TypeInt:
			fmt.Println("int")
			if offset+8 > uint16(len(data)) {
				return nil, fmt.Errorf("not enough data for int")
			}
			result = append(result, data[offset:offset+8])
			offset += 8

		case TypeString:
			fmt.Println("string")
			if offset+2 > uint16(len(data)) {
				return nil, fmt.Errorf("not enough data for string length")
			}
			length := binary.LittleEndian.Uint16(data[offset : offset+2])
			fullStringSize := 2 + length

			if offset+fullStringSize > uint16(len(data)) {
				return nil, fmt.Errorf("not enough data for string content")
			}

			result = append(result, data[offset:offset+fullStringSize])
			offset += fullStringSize

		case TypeDouble:
			fmt.Println("double")
			if offset+8 > uint16(len(data)) {
				return nil, fmt.Errorf("not enough data for double")
			}
			result = append(result, data[offset:offset+8])
			offset += 8

		default:
			return nil, fmt.Errorf("unsupported field type: %d", element.Type)
		}
	}

	return result, nil
}

type Record struct {
	ID    int
	Name  string
	Value float64
}
